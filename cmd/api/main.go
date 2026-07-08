package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"proyecto/internal/config"
	"proyecto/internal/handlers"
	"proyecto/internal/httpserver"
	"proyecto/internal/middleware"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// main queda DELGADO: carga la configuracion, delega en run y traduce el error
// a un exit code. Toda la logica de arranque vive en run, que devuelve error en
// lugar de llamar a log.Fatal en cada paso (mas testeable y mas limpio).
func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

// run construye las dependencias, levanta el servidor y bloquea hasta recibir
// una senal de apagado (Ctrl+C / SIGTERM); en ese momento hace un cierre
// ordenado (graceful shutdown): deja de aceptar conexiones, termina las que
// estan en curso y cierra la base de datos.
func run(cfg config.Config) error {
	// 1. Recursos de almacenamiento (Factory): abre DB (segun el motor elegido
	//    en la config: sqlite local o postgres en Docker), migra, siembra y elige backend.

	log.Printf("Inicializando almacenamiento: motor=%s, backend=%s", cfg.DBDriver, cfg.Backend)
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend de productos/categorias: %s", cfg.DBDriver, recursos.BackendUsado)

	// 2. Capa de servicio. AuthService recibe secreto y duracion por Options,
	//    tomados de la configuracion (antes eran globales hardcodeadas).
	// 3. Inyección de dependencias hacia Servicios y Handlers
	citaService := service.NuevoCitaService(recursos.Almacen)
	puntoService := service.NuevoPuntoEncuentroService(recursos.Almacen)
	soporteService := service.NuevoSoporteService(recursos.Almacen)

	authSvc := service.NuevoAuthService(
		recursos.Usuarios,
		service.WithSecreto(cfg.JWTSecreto),
		service.WithDuracionToken(cfg.JWTDuracion),
	)

	// 3. Server con sus dependencias agrupadas en un struct (escala sin crecer
	//    la firma del constructor).
	servidor := handlers.NewServer(handlers.Deps{
		Citas:           citaService,
		PuntosEncuentro: puntoService,
		Soportes:        soporteService,
		Auth:            authSvc,
	})

	// 4. Router Chi + Middlewares Globales
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Definición de Rutas API v1
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/registrar", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Rutas protegidas bajo Bearer Token
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			// Rutas Citas
			r.Get("/citas", servidor.ListarCitas)
			r.Post("/citas", servidor.CrearCita)
			r.Get("/citas/{id}", servidor.ObtenerCita)
			r.Put("/citas/{id}", servidor.ActualizarCita)
			r.Delete("/citas/{id}", servidor.BorrarCita)

			// Rutas Puntos de Encuentro
			r.Get("/puntos-encuentro", servidor.ListarPuntos)
			r.Post("/puntos-encuentro", servidor.CrearPunto)
			r.Get("/puntos-encuentro/{id}", servidor.ObtenerPunto)
			r.Put("/puntos-encuentro/{id}", servidor.ActualizarPunto)
			r.Delete("/puntos-encuentro/{id}", servidor.BorrarPunto)

			// Rutas Soportes
			r.Get("/soportes", servidor.ListarSoportes)
			r.Post("/soportes", servidor.CrearSoporte)
			r.Get("/soportes/{id}", servidor.ObtenerSoporte)
			r.Put("/soportes/{id}", servidor.ActualizarSoporte)
			r.Delete("/soportes/{id}", servidor.BorrarSoporte)
		})
	})

	// 6. Servidor HTTP configurado por Options (puerto + timeouts desde config).
	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	// 7. Contexto que se cancela al recibir Ctrl+C o SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 8. Arrancar el servidor en una goroutine para no bloquear la espera de la senal.
	errServidor := make(chan error, 1)
	go func() {
		log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	// 9. Esperar: o el servidor falla, o llega la senal de apagado.
	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Senal de apagado recibida, cerrando ordenadamente...")
	}

	// 10. Graceful shutdown: hasta 10s para terminar las requests en curso.
	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}

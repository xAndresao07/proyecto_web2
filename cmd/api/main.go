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

func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	log.Printf("Inicializando almacenamiento: motor=%s, backend=%s", cfg.DBDriver, cfg.Backend)
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend de solicitantes/dispositivos: %s", cfg.DBDriver, recursos.BackendUsado)

	solicitanteService := service.NuevoSolicitanteService(recursos.Almacen)
	dispositivoService := service.NuevoDispositivoService(recursos.Almacen)
	ticketService := service.NuevoTicketAyudaService(recursos.Almacen)
	tecnicoService := service.NuevoTecnicoService(recursos.Almacen)
	citaService := service.NuevoCitaService(recursos.Almacen)
	puntoService := service.NuevoPuntoEncuentroService(recursos.Almacen)
	soporteService := service.NuevoSoporteService(recursos.Almacen)

	authSvc := service.NuevoAuthService(
		recursos.Usuarios,
		service.WithSecreto(cfg.JWTSecreto),
		service.WithDuracionToken(cfg.JWTDuracion),
	)

	servidor := handlers.NewServer(handlers.Deps{
		Solicitantes: solicitanteService,
		Dispositivos: dispositivoService,
		Tickets:      ticketService,
		Auth:         authSvc,
		Tecnicos:     tecnicoService,
		Citas:        citaService,
		PuntosEncuentro:       puntoService,
		Soportes:     soporteService,
	})

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/registrar", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			// Rutas Solicitantes
			r.Get("/solicitantes", servidor.ListarSolicitantes)
			r.Post("/solicitantes", servidor.CrearSolicitante)
			r.Get("/solicitantes/{id}", servidor.ObtenerSolicitante)
			r.Put("/solicitantes/{id}", servidor.ActualizarSolicitante)
			r.Delete("/solicitantes/{id}", servidor.BorrarSolicitantes)

			// Rutas Dispositivos
			r.Get("/dispositivos", servidor.ListarDispositivos)
			r.Post("/dispositivos", servidor.CrearDispositivo)
			r.Get("/dispositivos/{id}", servidor.ObtenerDispositivo)
			r.Put("/dispositivos/{id}", servidor.ActualizarDispositivo)
			r.Delete("/dispositivos/{id}", servidor.BorrarDispositivo)

			
			// Rutas Tecnicos
			r.Get("/tecnicos", servidor.GetAllTecnicos)
			r.Post("/tecnicos", servidor.CreateTecnico)
			r.Get("/tecnicos/{id}", servidor.GetTecnicoPorID)
			r.Put("/tecnicos/{id}", servidor.UpdateTecnico)
			r.Delete("/tecnicos/{id}", servidor.DeleteTecnico)

			
			// Rutas Citas
			r.Get("/citas", servidor.ListarCitas)
			r.Post("/citas", servidor.CrearCita)
			r.Get("/citas/{id}", servidor.ObtenerCita)
			r.Put("/citas/{id}", servidor.ActualizarCita)
			r.Delete("/citas/{id}", servidor.BorrarCita)

			// Rutas Puntos Encuentro
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

			// Rutas Tickets de Ayuda
			r.Get("/tickets", servidor.ListarTickets)
			r.Post("/tickets", servidor.CrearTicket)
			r.Get("/tickets/{id}", servidor.ObtenerTicket)
			r.Put("/tickets/{id}", servidor.ActualizarTicket)
			r.Delete("/tickets/{id}", servidor.BorrarTicket)
		})
	})

	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errServidor := make(chan error, 1)
	go func() {
		log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Senal de apagado recibida, cerrando ordenadamente...")
	}

	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}

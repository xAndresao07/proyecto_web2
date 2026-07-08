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
	log.Printf("Motor de base de datos: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)

	// Inyección de dependencias hacia Servicios
	tecnicoService := service.NuevoTecnicoService(recursos.Almacen)

	authSvc := service.NuevoAuthService(
		recursos.Usuarios,
		service.WithSecreto(cfg.JWTSecreto),
		service.WithDuracionToken(cfg.JWTDuracion),
	)

	// Server con dependencias agrupadas
	servidor := handlers.NewServer(handlers.Deps{
		Tecnicos: tecnicoService,
		Auth:     authSvc,
	})

	// Router Chi + Middlewares Globales
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// Definición de Rutas API v1
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/registrar", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Rutas protegidas bajo Bearer Token
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			// Rutas Tecnicos
			r.Get("/tecnicos", servidor.ListarTecnicos)
			r.Post("/tecnicos", servidor.CrearTecnico)
			r.Get("/tecnicos/{id}", servidor.ObtenerTecnico)
			r.Put("/tecnicos/{id}", servidor.ActualizarTecnico)
			r.Delete("/tecnicos/{id}", servidor.BorrarTecnico)
		})
	})

	// Servidor HTTP configurado por Options
	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	// Contexto que se cancela al recibir Ctrl+C
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

	// Graceful shutdown
	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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
	// 1. Recursos de almacenamiento (Factory): abre DB, migra y elige backend.
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()

	// 2. Capa de servicio (Inyección)
	tecnicoService := service.NuevoTecnicoService(recursos.Almacen)
	authSvc := service.NuevoAuthService(recursos.Usuarios)

	// 3. Server con dependencias agrupadas
	servidor := handlers.NewServer(handlers.Deps{
		Tecnicos: tecnicoService,
		Auth:     authSvc,
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

			// Rutas Tecnicos
			r.Get("/tecnicos", servidor.GetAllTecnicos)
			r.Post("/tecnicos", servidor.CreateTecnico)
			r.Get("/tecnicos/{id}", servidor.GetTecnicoPorID)
			r.Put("/tecnicos/{id}", servidor.UpdateTecnico)
			r.Delete("/tecnicos/{id}", servidor.DeleteTecnico)
		})
	})

	// 6. Servidor HTTP configurado por Options
	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	// 7. Contexto que se cancela al recibir Ctrl+C
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

	// 8. Graceful shutdown
	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}

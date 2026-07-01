package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"proyecto/internal/handlers"
	"proyecto/internal/middleware"
	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

func main() {
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB, migra y siembra.
	//    Esto corre siempre, sin importar qué backend sirva después.
	gdb, err := gorm.Open(sqlite.Open("proyecto.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(&models.Cita{}, &models.PuntoEncuentro{}, &models.Soporte{}); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}
	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend que SIRVE las peticiones según la variable STORAGE.
	//    >>> Esta es la ÚNICA decisión que cambia entre GORM y sqlc. <<<
	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		// Ya migramos y sembramos con GORM; cerramos esa conexión para que
		// sqlc sea el único dueño del archivo cafeteria.db en tiempo de servicio.
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", "proyecto.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacen = almacenGorm
		log.Println("Backend de almacenamiento: GORM")
	}

	// 3. Inyección de dependencias hacia Servicios y Handlers
	usuarioRepo := storage.NewUsuarioRepository(gdb)
	authService := service.NuevoAuthService(usuarioRepo)
	citaService := service.NuevoCitaService(almacen)
	puntoService := service.NuevoPuntoEncuentroService(almacen)
	soporteService := service.NuevoSoporteService(almacen)

	servidor := handlers.NewServer(citaService, puntoService, soporteService, authService)

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
			r.Use(middleware.Auth(authService))

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

	log.Println("Servidor de Operaciones escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

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
	mimiddleware "proyecto/internal/middleware"
	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

func main() {
	// 1. GORM administra el esquema y la tabla de usuarios
	gdb, err := gorm.Open(sqlite.Open("tecnicos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(&models.Tecnico{}, &models.ServicioOfrecido{}, &models.HorarioTecnico{}, &models.Usuario{}); err != nil {
		log.Fatal("fallo AutoMigrate: ", err)
	}

	almacenGorm := storage.NuevoAlmacenSQLite(gdb)

	// 2. Elegir el backend de técnicos según la variable de entorno
	var almacen storage.TecnicoRepository
	switch os.Getenv("STORAGE") {
	case "sqlc":
		sdb, err := sql.Open("sqlite", "tecnicos.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de técnicos: sqlc (database/sql)")
	case "memoria":
		mem := storage.NewMemoria() // Asegúrate de haber adaptado tu memory.go a los IDs int
		almacen = mem
		log.Println("Backend de técnicos: Memoria")
	default:
		almacen = almacenGorm
		log.Println("Backend de técnicos: GORM")
	}

	usuarioRepo := storage.NuevoUsuarioGORM(gdb)

	// 3. Capa de servicio
	tecnicoSvc := service.NuevoTecnicoService(almacen)
	authSvc := service.NuevoAuthService(usuarioRepo)

	// 4. Inyección en los handlers
	servidor := handlers.NewServer(tecnicoSvc, authSvc)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(mimiddleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(mimiddleware.Auth(authSvc))

			r.Get("/tecnicos", servidor.GetAllTecnicos)
			r.Post("/tecnicos", servidor.CreateTecnico)
			r.Get("/tecnicos/{id}", servidor.GetTecnicoPorID)
			r.Put("/tecnicos/{id}", servidor.UpdateTecnico)
			r.Delete("/tecnicos/{id}", servidor.DeleteTecnico)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

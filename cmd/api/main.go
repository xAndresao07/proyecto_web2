package main

import (
	"log"
	"net/http"

	"proyecto/internal/handlers"
	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	// 1. GORM abre la Base de Datos SQLite local y autoconstruye las tablas
	gdb, err := gorm.Open(sqlite.Open("tecnicos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(&models.Tecnico{}, &models.ServicioOfrecido{}, &models.HorarioTecnico{}); err != nil {
		log.Fatal("fallo AutoMigrate: ", err)
	}

	// 2. Inyección de dependencias (Repository -> Service -> Handler)
	repo := storage.NuevoAlmacenSQLite(gdb)
	tecnicoSvc := service.NuevoTecnicoService(repo)
	servidor := handlers.NewServer(tecnicoSvc)

	// 3. Inicializamos el enrutador principal de Chi
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 4. Montamos los endpoints
	r.Route("/api/v1/tecnicos", func(r chi.Router) {
		r.Get("/", servidor.GetAllTecnicos)
		r.Post("/", servidor.CreateTecnico)
		r.Get("/{id}", servidor.GetTecnicoPorID)
		r.Put("/{id}", servidor.UpdateTecnico)
		r.Delete("/{id}", servidor.DeleteTecnico)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

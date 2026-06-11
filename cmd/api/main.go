package main

import (
	"log"
	"net/http"

	"proyecto/internal/handlers"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	almacen := storage.NewMemoria()
	almacen.Seed()
	servidor := handlers.NewServer(almacen)
	//Inicializamos el enrutador principal de Chi
	r := chi.NewRouter()

	// Middlewares básicos recomendados
	// Logger: Imprime en tu terminal cada vez que alguien hace una petición
	// Recoverer: Evita que el servidor se apague por completo si ocurre un error inesperado

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Subrouter: Agrupamos todas las rutas del módulo bajo prefijos

	r.Route("/api/v1/tecnicos", func(r chi.Router) {
		r.Get("/", servidor.GetAllTecnicos)
		r.Post("/", servidor.CreateTecnico)
		r.Get("/{id}", servidor.GetTecnicoPorID)
		r.Put("/{id}", servidor.UpdateTecnico)
		r.Delete("/{id}", servidor.DeleteTecnico)
	})
	// Iniciamos el servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

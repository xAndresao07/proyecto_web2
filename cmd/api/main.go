package main

import (
	"fmt"
	"net/http"

	"proyecto/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//Inicializamos el enrutador principal de Chi
	r := chi.NewRouter()

	// Middlewares básicos recomendados
	// Logger: Imprime en tu terminal cada vez que alguien hace una petición
	// Recoverer: Evita que el servidor se apague por completo si ocurre un error inesperado

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Subrouter: Agrupamos todas las rutas del módulo bajo prefijos

	r.Route("/api/v1/tecnicos", func(r chi.Router) {
		r.Get("/", handlers.GetAllTecnicos)
		r.Post("/", handlers.CreateTecnico)
		r.Get("/{id}", handlers.GetTecnicoPorID)
		r.Put("/{id}", handlers.UpdateTecnico)
		r.Delete("/{id}", handlers.DeleteTecnico)
	})
	// Iniciamos el servidor
	fmt.Println("Servidor de Tecnicos desplegado y escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", r)
}

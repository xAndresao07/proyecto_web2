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
	// Logger: Imprime en tu terminal cada vez que alguien hace una petición (útil para ver si funciona)
	// Recoverer: Evita que el servidor se apague por completo si ocurre un error inesperado

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Subrouter: Agrupamos todas las rutas del módulo bajo prefijos , ejemplo: /intervenciones

	r.Route("/api/v1/intervenciones", func(r chi.Router) {
		r.Get("/", handlers.GetAllIntervenciones)
		r.Post("/", handlers.CreateIntervencion)
		r.Get("/{id}", handlers.GetIntervencionPorID)
		r.Put("/{id}", handlers.UpdateIntervencion)
		r.Delete("/{id}", handlers.DeleteIntervencion)
	})
	// Iniciamos el servidor
	fmt.Println("Servidor de Intervenciones desplegado y escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", r)
}

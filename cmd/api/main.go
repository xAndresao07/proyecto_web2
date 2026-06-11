package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware" // Conservamos el paquete de middleware

	"proyecto/internal/handlers"
	"proyecto/internal/storage"
)

func main() {
	// 1. "Encendemos" el almacenamiento y cargamos las intervenciones semilla
	almacen := storage.NewMemoria()
	almacen.Seed()

	// 2. Inyectamos la base de datos al controlador de tu módulo
	servidor := handlers.NewServer(almacen)

	// 3. Inicializamos el router principal de Chi
	r := chi.NewRouter()

	// 4. Aplicamos los middlewares requeridos
	// Logger: para ver cada petición en la terminal
	// Recoverer: para que el servidor no "muera" si hay un error crítico
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Agrupamos y registramos las rutas de tu módulo de intervenciones
	r.Route("/api/v1/intervenciones", func(r chi.Router) {
		r.Get("/", servidor.ListarIntervenciones)
		r.Post("/", servidor.CrearIntervencion)
		r.Get("/{id}", servidor.ObtenerIntervencion)
		r.Delete("/{id}", servidor.BorrarIntervencion)
		r.Put("/{id}", servidor.ActualizarIntervencion)
	})

	// 6. Levantamos el servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

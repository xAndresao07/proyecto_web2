package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

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

	// 5. Agrupamos y registramos las rutas de tus endpoints bajo el prefijo /api/v1
	r.Route("/api/v1", func(r chi.Router) {

		// CRUD COMPLETO: Citas
		r.Route("/citas", func(r chi.Router) {
			r.Get("/", servidor.ListarCitas)
			r.Post("/", servidor.CrearCita)
			r.Get("/{id}", servidor.ObtenerCita)
			r.Put("/{id}", servidor.ActualizarCita)
			r.Delete("/{id}", servidor.EliminarCita)
		})

		// CRUD COMPLETO: Puntos de Encuentro
		r.Route("/puntos-encuentro", func(r chi.Router) {
			r.Get("/", servidor.ListarPuntos)
			r.Post("/", servidor.CrearPunto)
			r.Get("/{id}", servidor.ObtenerPunto)
			r.Put("/{id}", servidor.ActualizarPunto)
			r.Delete("/{id}", servidor.BorrarPunto)
		})

		// CRUD COMPLETO: Soportes
		r.Route("/soportes", func(r chi.Router) {
			r.Get("/", servidor.ListarSoportes)
			r.Post("/", servidor.CrearSoporte)
			r.Get("/{id}", servidor.ObtenerSoporte)
			r.Put("/{id}", servidor.ActualizarSoporte)
			r.Delete("/{id}", servidor.BorrarSoporte)
		})
	})

	// 6. Levantamos el servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

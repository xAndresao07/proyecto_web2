// Command solicitantes-api arranca el servidor HTTP del módulo Solicitantes y Hardware.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // driver GORM (pure-Go)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"solicitantesYHardware/internal/handlers"
	"solicitantesYHardware/internal/middleware"
	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/service"
	"solicitantesYHardware/internal/storage"
)

func main() {
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB, migra y siembra.
	//    Esto corre siempre, sin importar qué backend sirva después.
	gdb, err := gorm.Open(sqlite.Open("solicitantes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&models.Usuario{},
		&models.Solicitante{},
		&models.Dispositivo{},
		&models.TicketAyuda{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}
	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend que SIRVE las peticiones según la variable STORAGE.
	//    >>> Esta es la ÚNICA decisión que cambia entre Memoria, GORM y sqlc. <<<
	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "memoria":
		mem := storage.NuevaMemoria()
		mem.SeedSolicitantes()
		mem.SeedDispositivos()
		mem.SeedTicketAyudas()
		almacen = mem
		log.Println("Backend de almacenamiento: Memoria (datos volátiles)")
	case "sqlc":
		// Ya migramos y sembramos con GORM; cerramos esa conexión para que
		// sqlc sea el único dueño del archivo solicitantes.db en tiempo de servicio.
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", "solicitantes.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacen = almacenGorm
		log.Println("Backend de almacenamiento: GORM")
	}

	// 3. Services con inyección de dependencias. No saben qué backend recibieron.
	usuarioRepo := storage.NewUsuarioRepositorio(gdb)
	authService := service.NuevoAuthService(usuarioRepo)
	solicitanteService := service.NewSolicitanteService(almacen)
	dispositivoService := service.NewDispositivoService(almacen)
	ticketAyudaService := service.NewTicketAyudaService(almacen)

	servidor := handlers.NuevoServer(solicitanteService, dispositivoService, ticketAyudaService, authService)

	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {

		// Públicas — sin token.
		r.Post("/registro", servidor.Registrar)
		r.Post("/login", servidor.Login)

		// Protegidas — requieren: Authorization: Bearer <token>.
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/solicitantes", servidor.ListarSolicitantes)
			r.Post("/solicitantes", servidor.CrearSolicitante)
			r.Get("/solicitantes/{id}", servidor.ObtenerSolicitante)
			r.Put("/solicitantes/{id}", servidor.ActualizarSolicitante)
			r.Delete("/solicitantes/{id}", servidor.BorrarSolicitante)

			r.Get("/dispositivos", servidor.ListarDispositivos)
			r.Post("/dispositivos", servidor.CrearDispositivo)
			r.Get("/dispositivos/{id}", servidor.ObtenerDispositivo)
			r.Put("/dispositivos/{id}", servidor.ActualizarDispositivo)
			r.Delete("/dispositivos/{id}", servidor.BorrarDispositivo)

			r.Get("/tickets", servidor.ListarTickets)
			r.Post("/tickets", servidor.CrearTicket)
			r.Get("/tickets/{id}", servidor.ObtenerTicket)
			r.Put("/tickets/{id}", servidor.ActualizarTicket)
			r.Delete("/tickets/{id}", servidor.BorrarTicket)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

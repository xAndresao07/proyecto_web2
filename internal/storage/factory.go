package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // dialector GORM para SQLite (pure-Go)
	"gorm.io/driver/postgres"         // dialector GORM para PostgreSQL
	"gorm.io/gorm"

	"proyecto/internal/models"
)

type Recursos struct {
	Almacen      Almacen
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

func Inicializar(driver, dsn, rutaDB, backend string) (*Recursos, error) {
	log.Printf("Inicializando almacenamiento: motor=%s, backend=%s", driver, backend)

	gdb, err := abrirGorm(driver, dsn, rutaDB)
	if err != nil {
		return nil, err
	}
	if err := gdb.AutoMigrate(&models.Solicitante{}, &models.Dispositivo{}, &models.TicketAyuda{}, &models.Tecnico{}, &models.ServicioOfrecido{}, &models.HorarioTecnico{}, &models.Cita{}, &models.PuntoEncuentro{}, &models.Soporte{}); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}
	almacenGorm := NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	var almacen Almacen
	var sdb *sql.DB
	backendUsado := "gorm"
	if backend == "sqlc" && driver != "postgres" {
		log.Printf("Motor de base de datos: %s | Backend de solicitantes/dispositivos: sqlc (SQLite)", driver)
		sdb, err = sql.Open("sqlite", rutaDB)
		if err != nil {
			return nil, fmt.Errorf("abrir sql.DB para sqlc: %w", err)
		}
		almacen = NuevoAlmacenSQLC(sdb)
		backendUsado = "sqlc"
	} else {
		almacen = almacenGorm
	}

	usuarios := NewUsuarioRepository(gdb)

	cerrar := func() error {
		if sdb != nil {
			if err := sdb.Close(); err != nil {
				return err
			}
		}
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Almacen:      almacen,
		Usuarios:     usuarios,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

func abrirGorm(driver, dsn, rutaDB string) (*gorm.DB, error) {
	switch driver {
	case "postgres":
		var gdb *gorm.DB
		var err error
		for intento := 1; intento <= 10; intento++ {
			gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return gdb, nil
			}
			log.Printf("PostgreSQL no esta listo (intento %d/10): %v", intento, err)
			time.Sleep(2 * time.Second)
		}
		return nil, fmt.Errorf("conectar a PostgreSQL tras reintentos: %w", err)
	default: // "sqlite"
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}

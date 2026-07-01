package storage_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

func TestRepositorioGORM_CrearYBuscar(t *testing.T) {
	// 1. Abrimos una base de datos real, pero volátil (en memoria RAM)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos en memoria: %v", err)
	}

	// 2. Ejecutamos AutoMigrate para crear las tablas
	err = db.AutoMigrate(&models.Tecnico{}, &models.ServicioOfrecido{}, &models.HorarioTecnico{})
	if err != nil {
		t.Fatalf("falló la migración: %v", err)
	}

	repo := storage.NuevoAlmacenSQLite(db)

	// 3. Creamos un técnico con servicios anidados
	nuevo := models.Tecnico{
		Nombre:     "Carlos Test",
		Reputacion: 5.0,
		Servicios: []models.ServicioOfrecido{
			{NombreServicio: "Cambio de Pantalla"},
		},
	}

	// El repositorio lo guarda en SQLite (en memoria)
	creado := repo.CrearTecnico(nuevo)
	if creado.ID == 0 {
		t.Fatalf("se esperaba un ID generado por la BD, pero fue 0")
	}

	// 4. Buscamos el registro recién creado
	encontrado, existe := repo.BuscarTecnicoPorID(creado.ID)

	if !existe {
		t.Fatalf("se esperaba encontrar el técnico con ID %d, pero no existió", creado.ID)
	}
	if encontrado.Nombre != "Carlos Test" {
		t.Errorf("se esperaba el nombre 'Carlos Test', se obtuvo '%s'", encontrado.Nombre)
	}

	// 5. Validamos que Preload haya traído la relación (los servicios)
	if len(encontrado.Servicios) != 1 {
		t.Errorf("se esperaba cargar 1 servicio, se obtuvieron %d", len(encontrado.Servicios))
	}
}

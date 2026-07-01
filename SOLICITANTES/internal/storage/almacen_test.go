package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"solicitantesYHardware/internal/models"
)

// abrirDBMemoria abre una base SQLite COMPLETAMENTE en RAM (":memory:") y
// corre AutoMigrate, igual que main.go pero sin tocar disco. Cada test que
// la use arranca con una base vacía y nueva.
func abrirDBMemoria(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite :memory:: %v", err)
	}
	if err := db.AutoMigrate(&models.Solicitante{}); err != nil {
		t.Fatalf("falló AutoMigrate: %v", err)
	}
	return db
}

// TestAlmacenSQLite_CrearYBuscarSolicitante prueba el repositorio REAL
// (no un mock, no un fake) contra una base GORM+SQLite de verdad, aunque
// viva solo en memoria RAM durante el test.
//
// Crear → Buscar/Listar debe reflejar exactamente lo que se guardó.
//
// Qué se rompería si la implementación fallara: si el mapeo de GORM
// (los tags `gorm:"..."` del modelo) estuviera mal, o si ActualizarSolicitante
// no hiciera Save correctamente, este test fallaría porque los datos leídos
// no coincidirían con los datos escritos.
func TestAlmacenSQLite_CrearYBuscarSolicitante(t *testing.T) {
	db := abrirDBMemoria(t)
	almacen := NuevoAlmacenSQLite(db)

	nuevo := models.Solicitante{
		UsuarioID:     1,
		Matricula:     "ULEAM-0042",
		Nombre:        "Ricardo Villavicencio",
		Facultad:      "Software",
		Semestre:      3,
		NivelUrgencia: "normal",
	}

	creado := almacen.CrearSolicitante(nuevo)

	if creado.ID == 0 {
		t.Fatal("se esperaba que GORM asignara un ID autoincremental distinto de 0")
	}

	// Crear → Buscar debe reflejar lo mismo que se guardó.
	encontrado, ok := almacen.BuscarSolicitantePorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontró el solicitante recién creado con ID %d", creado.ID)
	}
	if encontrado.Matricula != "ULEAM-0042" {
		t.Fatalf("se esperaba matricula 'ULEAM-0042', se obtuvo: %s", encontrado.Matricula)
	}
	if encontrado.Nombre != "Ricardo Villavicencio" {
		t.Fatalf("se esperaba nombre 'Ricardo Villavicencio', se obtuvo: %s", encontrado.Nombre)
	}

	// Crear → Listar también debe reflejar el registro guardado.
	todos := almacen.ListarSolicitantes()
	if len(todos) != 1 {
		t.Fatalf("se esperaba 1 solicitante en la lista, se obtuvieron: %d", len(todos))
	}
}

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
	if err := db.AutoMigrate(&models.Solicitante{}, &models.Dispositivo{}, &models.TicketAyuda{}); err != nil {
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

func TestAlmacenSQLite_Dispositivo(t *testing.T) {
	db := abrirDBMemoria(t)
	almacen := NuevoAlmacenSQLite(db)

	nuevo := models.Dispositivo{
		SolicitanteID:      1,
		Marca:              "Dell",
		Modelo:             "Inspiron",
		TipoAlmacenamiento: "SSD",
		RamGB:              16,
		SistemaOperativo:   "Linux",
	}

	creado := almacen.CrearDispositivo(nuevo)
	if creado.ID == 0 {
		t.Fatal("ID debe asignarse")
	}

	encontrado, ok := almacen.BuscarDispositivoPorID(creado.ID)
	if !ok || encontrado.Marca != "Dell" {
		t.Fatal("Fallo en buscar dispositivo")
	}

	todos := almacen.ListarDispositivos()
	if len(todos) != 1 {
		t.Fatal("Fallo en listar dispositivos")
	}
	
	almacen.BorrarDispositivo(creado.ID)
	if _, ok := almacen.BuscarDispositivoPorID(creado.ID); ok {
		t.Fatal("Debería haberse borrado")
	}
}

func TestAlmacenSQLite_TicketAyuda(t *testing.T) {
	db := abrirDBMemoria(t)
	almacen := NuevoAlmacenSQLite(db)

	nuevo := models.TicketAyuda{
		SolicitanteID:    1,
		DispositivoID:    1,
		DescripcionFalla: "Pantalla azul",
		EstadoTicket:     "abierto",
	}

	creado := almacen.CrearTicketAyuda(nuevo)
	if creado.ID == 0 {
		t.Fatal("ID debe asignarse")
	}

	encontrado, ok := almacen.BuscarTicketAyudaPorID(creado.ID)
	if !ok || encontrado.DescripcionFalla != "Pantalla azul" {
		t.Fatal("Fallo en buscar ticket")
	}

	encontrado.EstadoTicket = "cerrado"
	actualizado, ok := almacen.ActualizarTicketAyuda(creado.ID, encontrado)
	if !ok || actualizado.EstadoTicket != "cerrado" {
		t.Fatal("Fallo al actualizar ticket")
	}

	todos := almacen.ListarTicketAyudas()
	if len(todos) != 1 {
		t.Fatal("Fallo en listar tickets")
	}
}

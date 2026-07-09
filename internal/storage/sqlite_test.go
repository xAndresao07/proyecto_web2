// Tests del repositorio REAL (GORM) contra una SQLite EN MEMORIA y desechable.
//
// A diferencia de los tests de service/handler —que usan dobles para evitar la
// base—, AQUI si tocamos una base de datos: es el unico lugar donde probamos la
// persistencia de verdad. Usamos ":memory:" para que sea rapida y sin archivos,
// y libreria testing estandar (sin testify).
//
// Nota: requiere el driver glebarez/sqlite (el mismo de produccion). Corre con
// `go test ./internal/storage`.
package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"proyecto/internal/models"
)

// nuevaDBPrueba abre una SQLite en memoria, migra el esquema y la devuelve.
// SetMaxOpenConns(1) garantiza que migracion y consultas usen la MISMA conexion
// (con ":memory:" cada conexion tendria su propia base, vacia).
func nuevaDBPrueba(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de prueba: %v", err)
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		t.Fatalf("no se pudo obtener *sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(1)
	if err := gdb.AutoMigrate(&models.Solicitante{}, &models.Dispositivo{}, &models.TicketAyuda{}, &models.Tecnico{}, &models.ServicioOfrecido{}, &models.HorarioTecnico{}, &models.Cita{}, &models.PuntoEncuentro{}, &models.Soporte{}, &models.Usuario{}); err != nil {
		t.Fatalf("fallo AutoMigrate: %v", err)
	}
	return gdb
}

func TestSQLite_SolicitanteCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// Crear: GORM debe asignar el ID autogenerado.
	creado := alm.CrearSolicitante(models.Solicitante{Nombre: "Yandry Cedeño", Facultad: "FACCI", Semestre: 6})
	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado por la base, obtuve 0")
	}

	// Buscar el recien creado.
	encontrado, ok := alm.BuscarSolicitantePorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontro el solicitante id=%d", creado.ID)
	}
	if encontrado.Nombre != "Yandry Cedeño" {
		t.Errorf("nombre = %q; esperaba %q", encontrado.Nombre, "Yandry Cedeño")
	}

	// Actualizar.
	if _, ok := alm.ActualizarSolicitante(creado.ID, models.Solicitante{Nombre: "Yandry Cedeño", Facultad: "FACCI", Semestre: 6}); !ok {
		t.Fatalf("no se pudo actualizar el solicitante id=%d", creado.ID)
	}

	// Borrar y confirmar que ya no esta.
	if !alm.BorrarSolicitante(creado.ID) {
		t.Errorf("esperaba poder borrar el solicitante id=%d", creado.ID)
	}
	if _, ok := alm.BuscarSolicitantePorID(creado.ID); ok {
		t.Errorf("el solicitante id=%d deberia haber sido borrado", creado.ID)
	}
}

func TestSQLite_BuscarInexistente(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	// El error de GORM (registro no encontrado) se traduce a comma-ok = false.
	if _, ok := alm.BuscarSolicitantePorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

// TestSQLite_UsuarioEmailUnico prueba una garantia que SOLO la base puede dar:
// el indice unico de email impide dos usuarios con el mismo correo.
func TestSQLite_UsuarioEmailUnico(t *testing.T) {
	repo := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	if _, err := repo.CrearUsuario(models.Usuario{Email: "ana@uleam.edu.ec", PasswordHash: "hash1"}); err != nil {
		t.Fatalf("el primer usuario deberia crearse sin error: %v", err)
	}
	if _, err := repo.CrearUsuario(models.Usuario{Email: "ana@uleam.edu.ec", PasswordHash: "hash2"}); err == nil {
		t.Errorf("esperaba error por email duplicado (indice unico), no lo hubo")
	}
}


func TestSQLite_DispositivoCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	creado := alm.CrearDispositivo(models.Dispositivo{Marca: "HP"})
	if creado.ID == 0 { t.Fatalf("ID no asignado") }
	encontrado, ok := alm.BuscarDispositivoPorID(creado.ID)
	if !ok || encontrado.Marca != "HP" { t.Fatalf("No se encontro") }
	_, ok = alm.ActualizarDispositivo(creado.ID, models.Dispositivo{Marca: "DELL"})
	if !ok { t.Fatalf("No se actualizo") }
	if !alm.BorrarDispositivo(creado.ID) { t.Fatalf("No se borro") }
	if len(alm.ListarDispositivos()) != 0 { t.Fatalf("Deberia estar vacio") }
}

func TestSQLite_TicketAyudaCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	creado := alm.CrearTicket(models.TicketAyuda{DescripcionFalla: "Falla"})
	if creado.ID == 0 { t.Fatalf("ID no asignado") }
	encontrado, ok := alm.BuscarTicketPorID(creado.ID)
	if !ok || encontrado.DescripcionFalla != "Falla" { t.Fatalf("No se encontro") }
	_, ok = alm.ActualizarTicket(creado.ID, models.TicketAyuda{DescripcionFalla: "Ok"})
	if !ok { t.Fatalf("No se actualizo") }
	if !alm.BorrarTicket(creado.ID) { t.Fatalf("No se borro") }
	if len(alm.ListarTickets()) != 0 { t.Fatalf("Deberia estar vacio") }
}

func TestSQLite_TecnicoCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	creado := alm.CrearTecnico(models.Tecnico{Nombre: "Tech"})
	if creado.ID == 0 { t.Fatalf("ID no asignado") }
	encontrado, ok := alm.BuscarTecnicoPorID(creado.ID)
	if !ok || encontrado.Nombre != "Tech" { t.Fatalf("No se encontro") }
	_, ok = alm.ActualizarTecnico(creado.ID, models.Tecnico{Nombre: "Tech2"})
	if !ok { t.Fatalf("No se actualizo") }
	if !alm.BorrarTecnico(creado.ID) { t.Fatalf("No se borro") }
	if len(alm.ListarTecnicos()) != 0 { t.Fatalf("Deberia estar vacio") }
}

func TestSQLite_CitaCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	creado := alm.CrearCita(models.Cita{Estado: "Pendiente"})
	if creado.ID == 0 { t.Fatalf("ID no asignado") }
	encontrado, ok := alm.BuscarCitaPorID(creado.ID)
	if !ok || encontrado.Estado != "Pendiente" { t.Fatalf("No se encontro") }
	_, ok = alm.ActualizarCita(creado.ID, models.Cita{Estado: "Fin"})
	if !ok { t.Fatalf("No se actualizo") }
	if !alm.BorrarCita(creado.ID) { t.Fatalf("No se borro") }
	if len(alm.ListarCitas()) != 0 { t.Fatalf("Deberia estar vacio") }
}

func TestSQLite_PuntoEncuentroCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	creado := alm.CrearPuntoEncuentro(models.PuntoEncuentro{NombreLugar: "Lab"})
	if creado.ID == 0 { t.Fatalf("ID no asignado") }
	encontrado, ok := alm.BuscarPuntoEncuentroPorID(creado.ID)
	if !ok || encontrado.NombreLugar != "Lab" { t.Fatalf("No se encontro") }
	_, ok = alm.ActualizarPuntoEncuentro(creado.ID, models.PuntoEncuentro{NombreLugar: "Lab2"})
	if !ok { t.Fatalf("No se actualizo") }
	if !alm.BorrarPuntoEncuentro(creado.ID) { t.Fatalf("No se borro") }
	if len(alm.ListarPuntosEncuentro()) != 0 { t.Fatalf("Deberia estar vacio") }
}

func TestSQLite_SoporteCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	creado := alm.CrearSoporte(models.Soporte{Solucion: "Sol"})
	if creado.ID == 0 { t.Fatalf("ID no asignado") }
	encontrado, ok := alm.BuscarSoportePorID(creado.ID)
	if !ok || encontrado.Solucion != "Sol" { t.Fatalf("No se encontro") }
	_, ok = alm.ActualizarSoporte(creado.ID, models.Soporte{Solucion: "Sol2"})
	if !ok { t.Fatalf("No se actualizo") }
	if !alm.BorrarSoporte(creado.ID) { t.Fatalf("No se borro") }
	if len(alm.ListarSoportes()) != 0 { t.Fatalf("Deberia estar vacio") }
}

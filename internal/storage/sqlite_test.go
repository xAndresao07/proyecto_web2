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

	// Migramos TODAS las tablas de tu dominio
	if err := gdb.AutoMigrate(&models.Cita{}, &models.PuntoEncuentro{}, &models.Soporte{}, &models.Usuario{}); err != nil {
		t.Fatalf("fallo AutoMigrate: %v", err)
	}
	return gdb
}

func TestSQLite_CitaCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// Crear: GORM debe asignar el ID autogenerado.
	creada := alm.CrearCita(models.Cita{
		SolicitanteID:  "est_102",
		TecnicoID:      "tec_05",
		Estado:         "pendiente",
		HoraAcordada:   "09:00",
		PuntoEncuentro: "Lab CISCO",
	})

	if creada.ID == 0 {
		t.Fatalf("esperaba un ID asignado por la base, obtuve 0")
	}

	// Buscar el recien creado.
	encontrada, ok := alm.BuscarCitaPorID(creada.ID)
	if !ok {
		t.Fatalf("no se encontro la cita id=%d", creada.ID)
	}
	if encontrada.SolicitanteID != "est_102" {
		t.Errorf("solicitante = %q; esperaba %q", encontrada.SolicitanteID, "est_102")
	}

	// Actualizar.
	datosActualizados := models.Cita{
		SolicitanteID:  "est_102",
		TecnicoID:      "tec_05",
		Estado:         "completada",
		HoraAcordada:   "10:00",
		PuntoEncuentro: "Lab CISCO",
	}
	if _, ok := alm.ActualizarCita(creada.ID, datosActualizados); !ok {
		t.Fatalf("no se pudo actualizar la cita id=%d", creada.ID)
	}

	// Borrar y confirmar que ya no esta.
	if !alm.BorrarCita(creada.ID) {
		t.Errorf("esperaba poder borrar la cita id=%d", creada.ID)
	}
	if _, ok := alm.BuscarCitaPorID(creada.ID); ok {
		t.Errorf("la cita id=%d deberia haber sido borrada", creada.ID)
	}
}

func TestSQLite_BuscarInexistente(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	// El error de GORM (registro no encontrado) se traduce a comma-ok = false.
	if _, ok := alm.BuscarCitaPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

// TestSQLite_UsuarioEmailUnico prueba una garantia que SOLO la base puede dar:
// el indice unico de email impide dos usuarios con el mismo correo.
func TestSQLite_UsuarioEmailUnico(t *testing.T) {
	repo := NewUsuarioRepository(nuevaDBPrueba(t))

	if _, err := repo.CrearUsuario(models.Usuario{Email: "ana@uleam.edu.ec", PasswordHash: "hash1", Rol: "admin"}); err != nil {
		t.Fatalf("el primer usuario deberia crearse sin error: %v", err)
	}

	// GORM intentará insertar este registro, pero SQLite lanzará error UNIQUE CONSTRAINT
	if _, err := repo.CrearUsuario(models.Usuario{Email: "ana@uleam.edu.ec", PasswordHash: "hash2", Rol: "tecnico"}); err == nil {
		t.Errorf("esperaba error por email duplicado (indice unico), no lo hubo")
	}
}

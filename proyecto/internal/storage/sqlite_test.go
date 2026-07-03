package storage_test

import (
	"testing"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/glebarez/sqlite" // El driver que usas
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAlmacenSQLite_CrearYBuscarCita(t *testing.T) {
	// 1. Abrimos una conexión a SQLite, pero usando ":memory:"
	// Esto crea una DB que solo existe mientras dura el test.
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err, "No debería fallar al abrir DB en memoria")

	// 2. Migramos la tabla para que exista
	err = db.AutoMigrate(&models.Cita{})
	assert.NoError(t, err, "No debería fallar al migrar")

	// 3. Inicializamos tu Repositorio SQLite
	repo := storage.NuevoAlmacenSQLite(db)

	// 4. Ejecutamos CrearCita (La acción a probar)
	nuevaCita := models.Cita{
		SolicitanteID:  "est_999",
		TecnicoID:      "tec_888",
		Estado:         "pendiente",
		HoraAcordada:   "10:00",
		PuntoEncuentro: "Laboratorio Redes",
	}
	citaCreada := repo.CrearCita(nuevaCita)

	// 5. Ejecutamos BuscarCitaPorID (Verificamos que el reflejo persista)
	citaEncontrada, existe := repo.BuscarCitaPorID(citaCreada.ID)

	// 6. Aserciones
	assert.True(t, existe, "La cita debería existir en la base de datos")
	assert.Equal(t, "est_999", citaEncontrada.SolicitanteID, "El ID del solicitante debe coincidir")
	assert.Equal(t, "Laboratorio Redes", citaEncontrada.PuntoEncuentro, "El lugar debe coincidir")
}

package service_test

import (
	"proyecto/internal/models"
	"proyecto/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 1. Definimos un "Mock" (Simulador) de nuestra interfaz CitaRepository
type MockCitaRepo struct {
	mock.Mock
}

// Simulamos los métodos de la interfaz
func (m *MockCitaRepo) ListarCitas() []models.Cita                 { return nil }
func (m *MockCitaRepo) BuscarCitaPorID(id int) (models.Cita, bool) { return models.Cita{}, false }
func (m *MockCitaRepo) ActualizarCita(id int, d models.Cita) (models.Cita, bool) {
	return models.Cita{}, false
}
func (m *MockCitaRepo) BorrarCita(id int) bool { return false }

// Este es el método que nos interesa espiar
func (m *MockCitaRepo) CrearCita(c models.Cita) models.Cita {
	args := m.Called(c)
	return args.Get(0).(models.Cita)
}

// 2. Escribimos el Test
func TestCrearCita_RechazaDatosVacios(t *testing.T) {
	// Preparamos nuestro repositorio falso y el servicio real
	mockRepo := new(MockCitaRepo)
	citaService := service.NuevoCitaService(mockRepo) // Inyectamos el mock

	// Creamos una cita sin SolicitanteID (Dato inválido según validacionCita)
	citaInvalida := models.Cita{
		SolicitanteID:  "",
		TecnicoID:      "tec_99",
		HoraAcordada:   "14:00",
		PuntoEncuentro: "Biblioteca",
	}

	// Ejecutamos la función
	_, err := citaService.Crear(citaInvalida)

	// Aserciones (Verificaciones)
	// Comprobamos que sí retornó el error esperado
	assert.ErrorIs(t, err, service.ErrNombreVacio)

	// ¡Lo más importante! Comprobamos que el repositorio NUNCA fue llamado.
	// Esto demuestra que la validación detuvo el proceso antes de tocar la DB.
	mockRepo.AssertNotCalled(t, "CrearCita")
}

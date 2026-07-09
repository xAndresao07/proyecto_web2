package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// puntoEncuentroRepoMock: doble de storage.PuntoEncuentroRepository (5 metodos).
type puntoEncuentroRepoMock struct {
	mock.Mock
}

func (m *puntoEncuentroRepoMock) ListarPuntosEncuentro() []models.PuntoEncuentro {
	return m.Called().Get(0).([]models.PuntoEncuentro)
}
func (m *puntoEncuentroRepoMock) BuscarPuntoEncuentroPorID(id int) (models.PuntoEncuentro, bool) {
	a := m.Called(id)
	return a.Get(0).(models.PuntoEncuentro), a.Bool(1)
}
func (m *puntoEncuentroRepoMock) CrearPuntoEncuentro(p models.PuntoEncuentro) models.PuntoEncuentro {
	return m.Called(p).Get(0).(models.PuntoEncuentro)
}
func (m *puntoEncuentroRepoMock) ActualizarPuntoEncuentro(id int, datos models.PuntoEncuentro) (models.PuntoEncuentro, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.PuntoEncuentro), a.Bool(1)
}
func (m *puntoEncuentroRepoMock) BorrarPuntoEncuentro(id int) bool {
	return m.Called(id).Bool(0)
}

// Red de seguridad en tiempo de compilación para garantizar que el Mock cumpla la interfaz
var _ storage.PuntoEncuentroRepository = (*puntoEncuentroRepoMock)(nil)

// El punto de encuentro tiene su propia regla: NombreLugar y FacultadPerteneciente son obligatorios.
func TestPuntoEncuentroService_Crear(t *testing.T) {
	t.Run("nombre lugar o facultad vacio rechazado", func(t *testing.T) {
		repo := new(puntoEncuentroRepoMock)

		// Probamos enviando campos en blanco, lo que debe disparar el ErrNombreVacio de tu validacion
		entradaInvalida := models.PuntoEncuentro{NombreLugar: "  ", FacultadPerteneciente: ""}
		_, err := service.NuevoPuntoEncuentroService(repo).Crear(entradaInvalida)

		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "CrearPuntoEncuentro")
	})

	t.Run("valida se persiste", func(t *testing.T) {
		repo := new(puntoEncuentroRepoMock)
		entrada := models.PuntoEncuentro{
			NombreLugar:           "Lab CISCO",
			FacultadPerteneciente: "FACCI",
			DisponibleParaSoporte: true,
		}

		guardada := entrada
		guardada.ID = 7
		repo.On("CrearPuntoEncuentro", entrada).Return(guardada)

		p, err := service.NuevoPuntoEncuentroService(repo).Crear(entrada)

		require.NoError(t, err)
		assert.Equal(t, 7, p.ID)
	})
}

func TestPuntoEncuentroService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(puntoEncuentroRepoMock)
	repo.On("BuscarPuntoEncuentroPorID", 999).Return(models.PuntoEncuentro{}, false)

	_, err := service.NuevoPuntoEncuentroService(repo).Obtener(999)
	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestPuntoEncuentroService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(puntoEncuentroRepoMock)
	repo.On("BorrarPuntoEncuentro", 999).Return(false)

	require.ErrorIs(t, service.NuevoPuntoEncuentroService(repo).Borrar(999), service.ErrNoEncontrado)
}

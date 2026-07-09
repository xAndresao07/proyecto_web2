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

// solicitanteRepoMock: doble de storage.SolicitanteRepository (5 metodos).
type solicitanteRepoMock struct {
	mock.Mock
}

func (m *solicitanteRepoMock) ListarSolicitantes() []models.Solicitante {
	return m.Called().Get(0).([]models.Solicitante)
}
func (m *solicitanteRepoMock) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Solicitante), a.Bool(1)
}
func (m *solicitanteRepoMock) CrearSolicitante(s models.Solicitante) models.Solicitante {
	return m.Called(s).Get(0).(models.Solicitante)
}
func (m *solicitanteRepoMock) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Solicitante), a.Bool(1)
}
func (m *solicitanteRepoMock) BorrarSolicitante(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.SolicitanteRepository = (*solicitanteRepoMock)(nil)

// El solicitante tiene su propia regla: nombre obligatorio.
func TestSolicitanteService_Crear(t *testing.T) {
	t.Run("nombre vacio rechazado", func(t *testing.T) {
		repo := new(solicitanteRepoMock)
		_, err := service.NuevoSolicitanteService(repo).Crear(models.Solicitante{Nombre: "  "})
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "CrearSolicitante")
	})
	t.Run("valida se persiste", func(t *testing.T) {
		repo := new(solicitanteRepoMock)

		// 1. Entrada que llega al servicio
		entrada := models.Solicitante{Nombre: "Yandry Cedeño", Facultad: "FACCI", Semestre: 6}

		// 2. Lo que el servicio realmente enviará al repo (con "media" incluido)
		esperadaEnRepo := entrada
		esperadaEnRepo.NivelUrgencia = "media"

		// 3. Resultado que el repo devuelve
		guardada := esperadaEnRepo
		guardada.ID = 7

		// AJUSTE: Aquí usamos la variable que tiene "media"
		repo.On("CrearSolicitante", esperadaEnRepo).Return(guardada)

		c, err := service.NuevoSolicitanteService(repo).Crear(entrada)
		require.NoError(t, err)
		assert.Equal(t, 7, c.ID)
		assert.Equal(t, "media", c.NivelUrgencia) // Verificamos que se haya seteado
	})
}

func TestSolicitanteService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(solicitanteRepoMock)
	repo.On("BuscarSolicitantePorID", 999).Return(models.Solicitante{}, false)
	_, err := service.NuevoSolicitanteService(repo).Obtener(999)
	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestSolicitanteService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(solicitanteRepoMock)
	repo.On("BorrarSolicitante", 999).Return(false)
	require.ErrorIs(t, service.NuevoSolicitanteService(repo).Borrar(999), service.ErrNoEncontrado)
}

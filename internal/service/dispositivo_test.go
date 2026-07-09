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

type dispositivoRepoMock struct {
	mock.Mock
}

func (m *dispositivoRepoMock) ListarDispositivos() []models.Dispositivo {
	return m.Called().Get(0).([]models.Dispositivo)
}
func (m *dispositivoRepoMock) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Dispositivo), a.Bool(1)
}
func (m *dispositivoRepoMock) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	return m.Called(d).Get(0).(models.Dispositivo)
}
func (m *dispositivoRepoMock) ActualizarDispositivo(id int, d models.Dispositivo) (models.Dispositivo, bool) {
	a := m.Called(id, d)
	return a.Get(0).(models.Dispositivo), a.Bool(1)
}
func (m *dispositivoRepoMock) BorrarDispositivo(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.DispositivoRepository = (*dispositivoRepoMock)(nil)

func TestDispositivoService_Crear(t *testing.T) {
	t.Run("datos invalidos", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		_, err := service.NuevoDispositivoService(repo).Crear(models.Dispositivo{})
		require.ErrorIs(t, err, service.ErrNombreVacio) // Validación falla
	})

	t.Run("crear exitoso", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		d := models.Dispositivo{ID: 1, SolicitanteID: 1, Marca: "Asus", Modelo: "ROG"}
		repo.On("CrearDispositivo", d).Return(d)

		creado, err := service.NuevoDispositivoService(repo).Crear(d)
		require.NoError(t, err)
		assert.Equal(t, 1, int(creado.ID))
	})
}

func TestDispositivoService_ListarYObtener(t *testing.T) {
	t.Run("listar", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		repo.On("ListarDispositivos").Return([]models.Dispositivo{{ID: 1, Marca: "HP"}})
		
		lista := service.NuevoDispositivoService(repo).Listar()
		assert.Len(t, lista, 1)
	})

	t.Run("obtener exitoso", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		repo.On("BuscarDispositivoPorID", 1).Return(models.Dispositivo{ID: 1, Marca: "HP"}, true)
		
		d, err := service.NuevoDispositivoService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "HP", d.Marca)
	})

	t.Run("obtener no encontrado", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		repo.On("BuscarDispositivoPorID", 99).Return(models.Dispositivo{}, false)
		
		_, err := service.NuevoDispositivoService(repo).Obtener(99)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

func TestDispositivoService_ActualizarYBorrar(t *testing.T) {
	t.Run("actualizar datos invalidos", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		_, err := service.NuevoDispositivoService(repo).Actualizar(1, models.Dispositivo{})
		require.ErrorIs(t, err, service.ErrNombreVacio)
	})

	t.Run("actualizar no encontrado", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		d := models.Dispositivo{ID: 1, SolicitanteID: 1, Marca: "Asus", Modelo: "ROG"}
		repo.On("ActualizarDispositivo", 1, d).Return(models.Dispositivo{}, false)

		_, err := service.NuevoDispositivoService(repo).Actualizar(1, d)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
    
    t.Run("actualizar exitoso", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		d := models.Dispositivo{ID: 1, SolicitanteID: 1, Marca: "Asus", Modelo: "ROG"}
		repo.On("ActualizarDispositivo", 1, d).Return(d, true)

		res, err := service.NuevoDispositivoService(repo).Actualizar(1, d)
		require.NoError(t, err)
        assert.Equal(t, "Asus", res.Marca)
	})

	t.Run("borrar exitoso", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		repo.On("BorrarDispositivo", 1).Return(true)

		err := service.NuevoDispositivoService(repo).Borrar(1)
		require.NoError(t, err)
	})

	t.Run("borrar no encontrado", func(t *testing.T) {
		repo := new(dispositivoRepoMock)
		repo.On("BorrarDispositivo", 99).Return(false)

		err := service.NuevoDispositivoService(repo).Borrar(99)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

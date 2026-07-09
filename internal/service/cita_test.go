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

// citaRepoMock es un doble de storage.CitaRepository (la interfaz
// estrecha de 5 metodos). Cada metodo solo registra la llamada y devuelve lo que
// el test configuro con On(...). No persiste nada.
type citaRepoMock struct {
	mock.Mock
}

func (m *citaRepoMock) ListarCitas() []models.Cita {
	return m.Called().Get(0).([]models.Cita)
}
func (m *citaRepoMock) BuscarCitaPorID(id int) (models.Cita, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Cita), a.Bool(1)
}
func (m *citaRepoMock) CrearCita(c models.Cita) models.Cita {
	return m.Called(c).Get(0).(models.Cita)
}
func (m *citaRepoMock) ActualizarCita(id int, datos models.Cita) (models.Cita, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Cita), a.Bool(1)
}
func (m *citaRepoMock) BorrarCita(id int) bool {
	return m.Called(id).Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato de CitaRepository.
var _ storage.CitaRepository = (*citaRepoMock)(nil)

// --- Crear: la regla de negocio (validacionCita), aislada de la base ---

func TestCitaService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Cita
		errEsperado   error // nil = exito
		debePersistir bool
	}{
		{
			nombre:        "campos obligatorios vacios rechazado",
			entrada:       models.Cita{SolicitanteID: "  ", TecnicoID: "tec_01"}, // Faltan hora y punto
			errEsperado:   service.ErrNombreVacio,
			debePersistir: false,
		},
		{
			nombre:        "cita valida se persiste",
			entrada:       models.Cita{SolicitanteID: "est_102", TecnicoID: "tec_05", HoraAcordada: "10:00", PuntoEncuentro: "Lab 1"},
			errEsperado:   nil,
			debePersistir: true,
		},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(citaRepoMock)

			// Si es válido, el servicio debería asignarle estado "pendiente" antes de llamar al repo
			entradaAlRepo := c.entrada
			if c.debePersistir && entradaAlRepo.Estado == "" {
				entradaAlRepo.Estado = "pendiente"
			}

			if c.debePersistir {
				guardado := entradaAlRepo
				guardado.ID = 42
				repo.On("CrearCita", entradaAlRepo).Return(guardado)
			}
			svc := service.NuevoCitaService(repo)

			creada, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearCita") // la validacion corto antes
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creada.ID)
				assert.Equal(t, "pendiente", creada.Estado) // Verificamos que se asignó el estado por defecto
				repo.AssertCalled(t, "CrearCita", entradaAlRepo)
			}
		})
	}
}

// --- Obtener: comma-ok del repo traducido a error de dominio ---

func TestCitaService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(citaRepoMock)
		repo.On("BuscarCitaPorID", 1).Return(models.Cita{ID: 1, SolicitanteID: "est_123"}, true)
		c, err := service.NuevoCitaService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "est_123", c.SolicitanteID)
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(citaRepoMock)
		repo.On("BuscarCitaPorID", 999).Return(models.Cita{}, false)
		_, err := service.NuevoCitaService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})
}

// --- Actualizar: valida ANTES de tocar el repo, y mapea el no encontrado ---

func TestCitaService_Actualizar(t *testing.T) {
	datosValidos := models.Cita{
		SolicitanteID:  "est_102",
		TecnicoID:      "tec_05",
		HoraAcordada:   "10:00",
		PuntoEncuentro: "Lab 1",
		Estado:         "completada",
	}

	t.Run("valido", func(t *testing.T) {
		repo := new(citaRepoMock)
		actualizado := datosValidos
		actualizado.ID = 1
		repo.On("ActualizarCita", 1, datosValidos).Return(actualizado, true)

		c, err := service.NuevoCitaService(repo).Actualizar(1, datosValidos)
		require.NoError(t, err)
		assert.Equal(t, 1, c.ID)
	})

	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(citaRepoMock)
		repo.On("ActualizarCita", 999, datosValidos).Return(models.Cita{}, false)

		_, err := service.NuevoCitaService(repo).Actualizar(999, datosValidos)
		require.ErrorIs(t, err, service.ErrNoEncontrado)
	})

	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(citaRepoMock)
		datosInvalidos := models.Cita{SolicitanteID: ""} // Campo vacio

		_, err := service.NuevoCitaService(repo).Actualizar(1, datosInvalidos)
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "ActualizarCita")
	})
}

// --- Borrar ---

func TestCitaService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(citaRepoMock)
		repo.On("BorrarCita", 1).Return(true)
		require.NoError(t, service.NuevoCitaService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrNoEncontrado", func(t *testing.T) {
		repo := new(citaRepoMock)
		repo.On("BorrarCita", 999).Return(false)
		require.ErrorIs(t, service.NuevoCitaService(repo).Borrar(999), service.ErrNoEncontrado)
	})
}

// --- Listar: el service solo delega ---

func TestCitaService_Listar(t *testing.T) {
	repo := new(citaRepoMock)
	repo.On("ListarCitas").Return([]models.Cita{{ID: 1}, {ID: 2}})
	lista := service.NuevoCitaService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}

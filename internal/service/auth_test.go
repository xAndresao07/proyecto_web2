package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

type usuarioRepoMock struct {
	mock.Mock
}

func (m *usuarioRepoMock) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	a := m.Called(u)
	return a.Get(0).(models.Usuario), a.Error(1)
}

func (m *usuarioRepoMock) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	a := m.Called(email)
	return a.Get(0).(models.Usuario), a.Bool(1)
}

var _ storage.UserRepository = (*usuarioRepoMock)(nil)

func TestAuthService_Registrar(t *testing.T) {
	t.Run("campos vacios rechazados", func(t *testing.T) {
		repo := new(usuarioRepoMock)
		svc := service.NuevoAuthService(repo)
		_, err := svc.Registrar("", "123", "admin")
		require.ErrorIs(t, err, service.ErrNombreVacio)
		repo.AssertNotCalled(t, "BuscarUsuarioPorEmail")
	})

	t.Run("email en uso", func(t *testing.T) {
		repo := new(usuarioRepoMock)
		repo.On("BuscarUsuarioPorEmail", "usado@uleam.edu.ec").Return(models.Usuario{ID: 1}, true)
		svc := service.NuevoAuthService(repo)
		_, err := svc.Registrar("usado@uleam.edu.ec", "123", "admin")
		require.ErrorIs(t, err, service.ErrEmailEnUso)
	})

	t.Run("registro exitoso", func(t *testing.T) {
		repo := new(usuarioRepoMock)
		repo.On("BuscarUsuarioPorEmail", "nuevo@uleam.edu.ec").Return(models.Usuario{}, false)

		// El hash de password no lo podemos predecir exacto al mockear CrearUsuario,
		// así que usamos mock.Anything
		repo.On("CrearUsuario", mock.AnythingOfType("models.Usuario")).Return(models.Usuario{ID: 2, Email: "nuevo@uleam.edu.ec"}, nil)

		svc := service.NuevoAuthService(repo)
		u, err := svc.Registrar("nuevo@uleam.edu.ec", "123", "admin")
		require.NoError(t, err)
		assert.Equal(t, 2, u.ID)
		repo.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	t.Run("email no existe", func(t *testing.T) {
		repo := new(usuarioRepoMock)
		repo.On("BuscarUsuarioPorEmail", "no@uleam.edu.ec").Return(models.Usuario{}, false)
		svc := service.NuevoAuthService(repo)
		_, err := svc.Login("no@uleam.edu.ec", "123")
		require.ErrorIs(t, err, service.ErrEmailEnUso) // Según auth.go devuelve ErrEmailEnUso
	})

	t.Run("contrasena incorrecta", func(t *testing.T) {
		repo := new(usuarioRepoMock)
		hash, _ := bcrypt.GenerateFromPassword([]byte("correcta"), bcrypt.DefaultCost)
		repo.On("BuscarUsuarioPorEmail", "si@uleam.edu.ec").Return(models.Usuario{ID: 1, PasswordHash: string(hash)}, true)

		svc := service.NuevoAuthService(repo)
		_, err := svc.Login("si@uleam.edu.ec", "incorrecta")
		require.ErrorIs(t, err, service.ErrCredencialesInvalidas)
	})

	t.Run("login exitoso", func(t *testing.T) {
		repo := new(usuarioRepoMock)
		hash, _ := bcrypt.GenerateFromPassword([]byte("correcta"), bcrypt.DefaultCost)
		repo.On("BuscarUsuarioPorEmail", "si@uleam.edu.ec").Return(models.Usuario{ID: 1, PasswordHash: string(hash)}, true)

		svc := service.NuevoAuthService(repo, service.WithDuracionToken(time.Minute))
		token, err := svc.Login("si@uleam.edu.ec", "correcta")
		require.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

func TestAuthService_ValidarToken(t *testing.T) {
	repo := new(usuarioRepoMock)
	svc := service.NuevoAuthService(repo, service.WithSecreto([]byte("secreto-test")), service.WithDuracionToken(time.Minute))

	// Para poder generar un token válido, usamos el propio servicio. No hay un método expuesto,
	// pero podemos usar el hash correcto
	hash, _ := bcrypt.GenerateFromPassword([]byte("correcta"), bcrypt.DefaultCost)
	repo.On("BuscarUsuarioPorEmail", "valid@uleam.edu.ec").Return(models.Usuario{ID: 55, PasswordHash: string(hash)}, true)

	token, err := svc.Login("valid@uleam.edu.ec", "correcta")
	require.NoError(t, err)

	t.Run("token valido", func(t *testing.T) {
		id, err := svc.ValidarToken(token)
		require.NoError(t, err)
		assert.Equal(t, 55, id)
	})

	t.Run("token invalido", func(t *testing.T) {
		_, err := svc.ValidarToken("token_cualquiera_falso")
		require.ErrorIs(t, err, service.ErrCredencialesInvalidas)
	})
}

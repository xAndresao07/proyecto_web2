package service

import (
	"testing"
	"solicitantesYHardware/internal/models"
)

type mockUserRepo struct {
	usuarios map[string]models.Usuario
	nextID   int
}

func (m *mockUserRepo) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := m.usuarios[email]
	return u, ok
}

func (m *mockUserRepo) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = m.nextID
	m.nextID++
	m.usuarios[u.Email] = u
	return u, nil
}

func TestAuthService(t *testing.T) {
	repo := &mockUserRepo{
		usuarios: make(map[string]models.Usuario),
		nextID:   1,
	}
	svc := NuevoAuthService(repo)

	// Test Registrar
	t.Run("Registrar", func(t *testing.T) {
		u, err := svc.Registrar("test@example.com", "123456")
		if err != nil {
			t.Fatalf("No se esperaba error al registrar: %v", err)
		}
		if u.Email != "test@example.com" {
			t.Fatalf("Email incorrecto, esperado %s, obtenido %s", "test@example.com", u.Email)
		}
	})

	t.Run("Registrar duplicado", func(t *testing.T) {
		_, err := svc.Registrar("test@example.com", "123456")
		if err != ErrEmailEnUso {
			t.Fatalf("Esperado ErrEmailEnUso, obtenido %v", err)
		}
	})

	t.Run("Registrar vacio", func(t *testing.T) {
		_, err := svc.Registrar("", "123")
		if err != ErrNombreVacio {
			t.Fatalf("Esperado ErrNombreVacio, obtenido %v", err)
		}
	})

	// Test Login
	t.Run("Login valido", func(t *testing.T) {
		token, err := svc.Login("test@example.com", "123456")
		if err != nil {
			t.Fatalf("Error en login: %v", err)
		}
		if token == "" {
			t.Fatal("Token vacio")
		}

		// Validar token
		uid, err := svc.ValidarToken(token)
		if err != nil {
			t.Fatalf("Error validando token: %v", err)
		}
		if uid != 1 {
			t.Fatalf("UID incorrecto: esperado 1, obtenido %d", uid)
		}
	})

	t.Run("Login invalido password", func(t *testing.T) {
		_, err := svc.Login("test@example.com", "123")
		if err != ErrCredencialesInvalidas {
			t.Fatalf("Esperado ErrCredencialesInvalidas, obtenido %v", err)
		}
	})

	t.Run("Login invalido email", func(t *testing.T) {
		_, err := svc.Login("noexiste@example.com", "123")
		if err != ErrCredencialesInvalidas {
			t.Fatalf("Esperado ErrCredencialesInvalidas, obtenido %v", err)
		}
	})
	
	t.Run("Token invalido", func(t *testing.T) {
		_, err := svc.ValidarToken("tokenfalso")
		if err == nil {
			t.Fatal("Esperaba error con token falso")
		}
	})
}

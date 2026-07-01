package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"proyecto/internal/handlers"
	"proyecto/internal/middleware"
	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// 1. Creamos un "Fake" ultraligero exclusivo para el repositorio de Usuarios
type FakeUsuarioRepo struct{}

func (f *FakeUsuarioRepo) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	return u, nil
}
func (f *FakeUsuarioRepo) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	return models.Usuario{}, false
}

func TestListarCitas_Retorna401SinToken(t *testing.T) {
	// 2. Usamos tu implementación en Memoria como un "Fake" para Citas
	memoriaFake := storage.NuevaMemoria()

	// 3. Inicializamos los servicios
	citaService := service.NuevoCitaService(memoriaFake)

	//Le pasamos nuestro nuevo FakeUsuarioRepo al AuthService
	authService := service.NuevoAuthService(&FakeUsuarioRepo{})

	// 4. Inicializamos el Server
	server := handlers.NewServer(citaService, nil, nil, authService)

	// 5. Configuramos un enrutador (Router) igual que en main.go
	r := chi.NewRouter()

	// Le inyectamos el Middleware de Autenticación que protegerá la ruta
	r.Use(middleware.Auth(authService))
	r.Get("/api/v1/citas", server.ListarCitas)

	// 6. Simulamos una petición HTTP del cliente (SIN encabezado Authorization)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/citas", nil)

	// httptest.NewRecorder() actúa como el "navegador"
	rec := httptest.NewRecorder()

	// 7. Ejecutamos la petición contra nuestro router
	r.ServeHTTP(rec, req)

	// 8. Aserciones: Comprobamos que el código de estado sea 401
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestListarCitas_ConFakeEnMemoria(t *testing.T) {
	// 1. Usamos tu implementación en Memoria
	memoriaFake := storage.NuevaMemoria()

	// Inyectamos datos semilla para que listar devuelva algo
	memoriaFake.SeedCitas()

	// 2. Inicializamos los servicios
	citaService := service.NuevoCitaService(memoriaFake)

	// 3. Inicializamos el Server (sin Auth porque no lo montaremos en el router de prueba)
	server := handlers.NewServer(citaService, nil, nil, nil)

	// 4. Configuramos el router
	// (En este test NO usamos r.Use(Auth) para poder probar el handler directamente sin lidiar con tokens)
	r := chi.NewRouter()
	r.Get("/api/v1/citas", server.ListarCitas)

	// 5. Simulamos la Petición HTTP
	req := httptest.NewRequest(http.MethodGet, "/api/v1/citas", nil)
	rec := httptest.NewRecorder()

	// 6. Ejecutamos
	r.ServeHTTP(rec, req)

	// 7. Aserciones
	assert.Equal(t, http.StatusOK, rec.Code, "Debería responder 200 OK")

	// Comprobamos que el JSON de respuesta contenga el solicitante que inyectó el Seed
	assert.Contains(t, rec.Body.String(), "est_102", "El cuerpo de la respuesta debe contener la cita semilla")
}

package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"proyecto/internal/handlers"
	"proyecto/internal/middleware"
	"proyecto/internal/models"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// =====================================================================
// Dobles de prueba
// =====================================================================

// usuarioFake implementa storage.UsuarioRepository en memoria.
// Esto es necesario porque storage.Memoria no maneja usuarios.
type usuarioFake struct {
	porEmail map[string]models.Usuario
	nextID   int
}

func nuevoUsuarioFake() *usuarioFake {
	return &usuarioFake{porEmail: map[string]models.Usuario{}, nextID: 1}
}

func (f *usuarioFake) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioFake) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

var _ storage.UsuarioRepository = (*usuarioFake)(nil)

// =====================================================================
// Router de prueba
// =====================================================================

// construirEntorno devuelve el handler listo y el entorno sembrado (con la cita 1).
func construirEntorno() (http.Handler, *storage.Memoria, *usuarioFake) {
	// 1. Usamos tu implementación Memoria ya existente en tu carpeta storage
	almacen := storage.NuevaMemoria()

	// Sembramos una cita inicial para que listar y obtener devuelvan datos
	almacen.SeedCitas()

	// 2. Usamos el fake de usuarios que creamos arriba
	usuarios := nuevoUsuarioFake()

	// 3. Inicializamos tus servicios
	citaSvc := service.NuevoCitaService(almacen)
	puntoSvc := service.NuevoPuntoEncuentroService(almacen)
	soporteSvc := service.NuevoSoporteService(almacen)
	authSvc := service.NuevoAuthService(usuarios)

	// 4. Inicializamos los handlers (NewServer)
	srv := handlers.NewServer(handlers.Deps{
		Citas:           citaSvc,
		PuntosEncuentro: puntoSvc,
		Soportes:        soporteSvc,
		Auth:            authSvc,
	})

	// 5. Configuramos el Router idéntico al main.go
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		// Rutas públicas
		r.Post("/auth/registrar", srv.Registrar)
		r.Post("/auth/login", srv.Login)

		// Rutas protegidas
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc)) // <- Tu middleware real

			r.Get("/citas", srv.ListarCitas)
			r.Post("/citas", srv.CrearCita)
			r.Get("/citas/{id}", srv.ObtenerCita)
			r.Put("/citas/{id}", srv.ActualizarCita)
			r.Delete("/citas/{id}", srv.BorrarCita)
		})
	})
	return r, almacen, usuarios
}

// jsonReq arma una petición con cuerpo JSON y, si se pasa token, el header Bearer.
func jsonReq(metodo, ruta, cuerpo, token string) *http.Request {
	var body *strings.Reader
	if cuerpo == "" {
		body = strings.NewReader("")
	} else {
		body = strings.NewReader(cuerpo)
	}
	req := httptest.NewRequest(metodo, ruta, body)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req
}

// tokenValido registra y loguea un usuario contra el router y devuelve su JWT.
func tokenValido(t *testing.T, h http.Handler) string {
	t.Helper()

	// Tu API necesita el "rol" al registrar, por eso lo añadimos aquí:
	cred := `{"email":"estudiante@uleam.edu.ec","password":"123","rol":"solicitante"}`

	// Paso 1: Registrar al usuario falso
	h.ServeHTTP(httptest.NewRecorder(), jsonReq(http.MethodPost, "/api/v1/auth/registrar", cred, ""))

	// Paso 2: Loguearse para obtener el token
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, jsonReq(http.MethodPost, "/api/v1/auth/login", cred, ""))
	require.Equal(t, http.StatusOK, rec.Code, "El login falló al generar el token")

	// Paso 3: Extraer el token de la respuesta JSON
	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)

	return resp.Token
}

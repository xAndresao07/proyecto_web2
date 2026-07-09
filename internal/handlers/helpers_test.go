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

func construirEntorno() (http.Handler, *storage.Memoria, *usuarioFake) {
	almacen := storage.NuevaMemoria()
	almacen.SeedSolicitantes()
	almacen.SeedDispositivos()
	almacen.SeedTickets()
	almacen.SeedCitas()

	usuarios := nuevoUsuarioFake()
	authSvc := service.NuevoAuthService(usuarios)

	srv := handlers.NewServer(handlers.Deps{
		Solicitantes:    service.NuevoSolicitanteService(almacen),
		Dispositivos:    service.NuevoDispositivoService(almacen),
		Citas:           service.NuevoCitaService(almacen),
		PuntosEncuentro: service.NuevoPuntoEncuentroService(almacen),
		Soportes:        service.NuevoSoporteService(almacen),
		Auth:            authSvc,
	})

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", srv.Registrar)
		r.Post("/auth/login", srv.Login)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Get("/solicitantes", srv.ListarSolicitantes)
			r.Post("/solicitantes", srv.CrearSolicitante)
			r.Get("/solicitantes/{id}", srv.ObtenerSolicitante)
			r.Put("/solicitantes/{id}", srv.ActualizarSolicitante)
			r.Delete("/solicitantes/{id}", srv.BorrarSolicitantes)

			r.Get("/dispositivos", srv.ListarDispositivos)
			r.Post("/dispositivos", srv.CrearDispositivo)
			r.Get("/dispositivos/{id}", srv.ObtenerDispositivo)
			r.Put("/dispositivos/{id}", srv.ActualizarDispositivo)
			r.Delete("/dispositivos/{id}", srv.BorrarDispositivo)

			r.Get("/citas", srv.ListarCitas)
			r.Post("/citas", srv.CrearCita)
			r.Get("/citas/{id}", srv.ObtenerCita)
			r.Put("/citas/{id}", srv.ActualizarCita)
			r.Delete("/citas/{id}", srv.BorrarCita)

			r.Get("/puntos-encuentro", srv.ListarPuntos)
			r.Post("/puntos-encuentro", srv.CrearPunto)
			r.Get("/puntos-encuentro/{id}", srv.ObtenerPunto)
			r.Put("/puntos-encuentro/{id}", srv.ActualizarPunto)
			r.Delete("/puntos-encuentro/{id}", srv.BorrarPunto)

			r.Get("/soportes", srv.ListarSoportes)
			r.Post("/soportes", srv.CrearSoporte)
			r.Get("/soportes/{id}", srv.ObtenerSoporte)
			r.Put("/soportes/{id}", srv.ActualizarSoporte)
			r.Delete("/soportes/{id}", srv.BorrarSoporte)
		})
	})
	return r, almacen, usuarios
}

func jsonReq(metodo, ruta, cuerpo, token string) *http.Request {
	var body *strings.Reader
	if cuerpo == "" {
		body = strings.NewReader("")
	} else {
		body = strings.NewReader(cuerpo)
	}
	req := httptest.NewRequest(metodo, ruta, body)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req
}

func tokenValido(t *testing.T, h http.Handler) string {
	t.Helper()
	cred := `{"email":"docente@uleam.edu.ec","password":"secreta123","rol":"admin"}`

	recReg := httptest.NewRecorder()
	h.ServeHTTP(recReg, jsonReq(http.MethodPost, "/api/v1/auth/register", cred, ""))

	if recReg.Code != http.StatusCreated && recReg.Code != http.StatusConflict {
		require.Equal(t, http.StatusCreated, recReg.Code)
	}

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, jsonReq(http.MethodPost, "/api/v1/auth/login", cred, ""))
	require.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)
	return resp.Token
}

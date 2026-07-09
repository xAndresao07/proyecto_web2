package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"solicitantesYHardware/internal/handlers"
	"solicitantesYHardware/internal/middleware"
	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/service"
	"solicitantesYHardware/internal/storage"
)

// =====================================================================
// Almacen Fake
// =====================================================================

type almacenFake struct {
	solicitantes map[int]models.Solicitante
	dispositivos map[int]models.Dispositivo
	nextID       int
}

func nuevoAlmacenFake() *almacenFake {
	return &almacenFake{
		solicitantes: make(map[int]models.Solicitante),
		dispositivos: make(map[int]models.Dispositivo),
		nextID:       1,
	}
}

func (a *almacenFake) ListarSolicitantes() []models.Solicitante {
	out := make([]models.Solicitante, 0, len(a.solicitantes))
	for _, s := range a.solicitantes {
		out = append(out, s)
	}
	return out
}

func (a *almacenFake) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	s, ok := a.solicitantes[id]
	return s, ok
}

func (a *almacenFake) CrearSolicitante(s models.Solicitante) models.Solicitante {
	s.ID = a.nextID
	a.nextID++
	a.solicitantes[s.ID] = s
	return s
}

func (a *almacenFake) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	if _, ok := a.solicitantes[id]; !ok {
		return models.Solicitante{}, false
	}
	datos.ID = id
	a.solicitantes[id] = datos
	return datos, true
}

func (a *almacenFake) BorrarSolicitante(id int) bool {
	if _, ok := a.solicitantes[id]; !ok {
		return false
	}
	delete(a.solicitantes, id)
	return true
}

// Asegúrate de implementar los métodos de otras interfaces si las tienes (ej. Dispositivo, Ticket)
var _ storage.SolicitanteRepositorio = (*almacenFake)(nil)
var _ storage.DispositivoRepositorio = (*almacenFake)(nil)
var _ storage.TicketAyudaRepositorio = (*almacenFake)(nil)

func (a *almacenFake) ListarDispositivos() []models.Dispositivo {
	out := make([]models.Dispositivo, 0, len(a.dispositivos))
	for _, d := range a.dispositivos {
		out = append(out, d)
	}
	return out
}

func (a *almacenFake) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	d, ok := a.dispositivos[id]
	return d, ok
}

func (a *almacenFake) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	d.ID = a.nextID
	a.nextID++
	a.dispositivos[int(d.ID)] = d
	return d
}

func (a *almacenFake) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	if _, ok := a.dispositivos[id]; !ok {
		return models.Dispositivo{}, false
	}
	datos.ID = id
	a.dispositivos[id] = datos
	return datos, true
}

func (a *almacenFake) BorrarDispositivo(id int) bool {
	if _, ok := a.dispositivos[id]; !ok {
		return false
	}
	delete(a.dispositivos, id)
	return true
}

func (a *almacenFake) ListarTicketAyudas() []models.TicketAyuda { return nil }
func (a *almacenFake) BuscarTicketAyudaPorID(id int) (models.TicketAyuda, bool) { return models.TicketAyuda{}, false }
func (a *almacenFake) CrearTicketAyuda(t models.TicketAyuda) models.TicketAyuda { return t }
func (a *almacenFake) ActualizarTicketAyuda(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) { return datos, true }
func (a *almacenFake) BorrarTicketAyuda(id int) bool { return true }



// =====================================================================
// Usuario Fake
// =====================================================================

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

var _ storage.UserRepositorio = (*usuarioFake)(nil)



// =====================================================================
// Entorno de Prueba
// =====================================================================

func construirEntorno() (http.Handler, *almacenFake, *usuarioFake) {
	almacen := nuevoAlmacenFake()
	almacen.CrearSolicitante(models.Solicitante{Nombre: "Jandry Cedeño", Facultad: "FACCI", Semestre: 6})

	usuarios := nuevoUsuarioFake()
	authSvc := service.NuevoAuthService(usuarios)

	srv := handlers.NuevoServer(
		service.NewSolicitanteService(almacen),
		service.NewDispositivoService(almacen),
		service.NewTicketAyudaService(almacen),
		authSvc,
	)

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
			r.Delete("/solicitantes/{id}", srv.BorrarSolicitante)

			r.Get("/dispositivos", srv.ListarDispositivos)
			r.Post("/dispositivos", srv.CrearDispositivo)
			r.Get("/dispositivos/{id}", srv.ObtenerDispositivo)
			r.Put("/dispositivos/{id}", srv.ActualizarDispositivo)
			r.Delete("/dispositivos/{id}", srv.BorrarDispositivo)
		})
	})
	return r, almacen, usuarios
}

// =====================================================================
// Helpers (La magia para testear HTTP)
// =====================================================================

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
	// Añadimos el "rol" porque es probable que tu handler lo valide como obligatorio
	cred := `{"email":"docente@uleam.edu.ec","password":"secreta123","rol":"admin"}`

	// 1. Registrar usuario
	recReg := httptest.NewRecorder()
	h.ServeHTTP(recReg, jsonReq(http.MethodPost, "/api/v1/auth/register", cred, ""))

	// Si da 409 (conflicto) es porque ya existe, lo cual está bien para los tests
	if recReg.Code != http.StatusCreated && recReg.Code != http.StatusConflict {
		// Si da 400, imprimimos el cuerpo para saber qué validación falló
		t.Logf("Error en registro: %s", recReg.Body.String())
		require.Equal(t, http.StatusCreated, recReg.Code)
	}

	// 2. Loguear usuario
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

func ejecutar(h http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"solicitantesYHardware/internal/middleware"
	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/service"
)

// fakeSolicitanteRepo es un DOBLE que SÍ guarda datos (a diferencia del mock
// del service_test, que solo detecta llamadas). Vive en un slice en memoria,
// no es la base real (SQLite), pero se comporta como un repositorio de verdad
// para poder probar el handler de punta a punta sin levantar una base.
type fakeSolicitanteRepo struct {
	datos  []models.Solicitante
	nextID int
}

func nuevoFakeSolicitanteRepo() *fakeSolicitanteRepo {
	return &fakeSolicitanteRepo{nextID: 1}
}

func (f *fakeSolicitanteRepo) ListarSolicitantes() []models.Solicitante {
	return f.datos
}

func (f *fakeSolicitanteRepo) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	for _, s := range f.datos {
		if s.ID == id {
			return s, true
		}
	}
	return models.Solicitante{}, false
}

func (f *fakeSolicitanteRepo) CrearSolicitante(s models.Solicitante) models.Solicitante {
	s.ID = f.nextID
	f.nextID++
	f.datos = append(f.datos, s)
	return s
}

func (f *fakeSolicitanteRepo) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	for i, s := range f.datos {
		if s.ID == id {
			datos.ID = id
			f.datos[i] = datos
			return datos, true
		}
	}
	return models.Solicitante{}, false
}

func (f *fakeSolicitanteRepo) BorrarSolicitante(id int) bool {
	for i, s := range f.datos {
		if s.ID == id {
			f.datos = append(f.datos[:i], f.datos[i+1:]...)
			return true
		}
	}
	return false
}

// TestCrearSolicitante_HTTP_CasoFeliz prueba que POST /solicitantes con un
// JSON válido responde 201 Created y devuelve el solicitante con ID asignado.
//
// Qué se rompería si la implementación fallara: si el handler dejara de
// decodificar el body, o el service dejara de asignar el ID, este test
// detecta que la respuesta HTTP ya no trae los datos correctos.
func TestCrearSolicitante_HTTP_CasoFeliz(t *testing.T) {
	repo := nuevoFakeSolicitanteRepo()
	svc := service.NewSolicitanteService(repo)
	srv := &Server{Solicitantes: svc}

	body := `{
		"usuario_id": 1,
		"matricula": "ULEAM-0001",
		"nombre": "Carlos Moreira",
		"facultad": "TI",
		"semestre": 4,
		"nivel_urgencia": "alto"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/solicitantes", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	srv.CrearSolicitante(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("se esperaba status 201, se obtuvo: %d", res.StatusCode)
	}

	var creado models.Solicitante
	if err := json.NewDecoder(res.Body).Decode(&creado); err != nil {
		t.Fatalf("no se pudo decodificar la respuesta: %v", err)
	}
	if creado.ID == 0 {
		t.Fatal("el solicitante creado debería tener un ID asignado distinto de 0")
	}
	if creado.Nombre != "Carlos Moreira" {
		t.Fatalf("se esperaba nombre 'Carlos Moreira', se obtuvo: %s", creado.Nombre)
	}
}

// TestRutaProtegida_SinToken_Devuelve401 prueba el middleware Auth real
// (no un fake) montado sobre una ruta protegida, tal como está armada en
// cmd/api/main.go. Una petición SIN header Authorization debe responder
// 401 Unauthorized y NUNCA debe ejecutar el handler de abajo.
//
// Qué se rompería si la implementación fallara: si alguien quita el
// middleware Auth de la ruta, o lo deja pasar por error, este test detecta
// que datos protegidos quedarían expuestos sin autenticación.
func TestRutaProtegida_SinToken_Devuelve401(t *testing.T) {
	repo := nuevoFakeSolicitanteRepo()
	svc := service.NewSolicitanteService(repo)
	srv := &Server{Solicitantes: svc}

	// AuthService con repo nil: es válido porque el 401 ocurre ANTES de que
	// el middleware necesite consultar ningún usuario (no hay header Bearer).
	authSvc := service.NuevoAuthService(nil)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSvc))
		r.Get("/api/v1/solicitantes", srv.ListarSolicitantes)
	})

	// Petición SIN header Authorization.
	req := httptest.NewRequest(http.MethodGet, "/api/v1/solicitantes", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("se esperaba status 401, se obtuvo: %d", res.StatusCode)
	}
}

package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"proyecto/internal/models"
)

// ejecutar corre una peticion contra el handler y devuelve el recorder.
func ejecutar(h http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestListarCitas_OK(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/citas", "", token))

	require.Equal(t, http.StatusOK, rec.Code)
	var lista []models.Cita
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	// Asegúrate de que tu fake en memoria siembre al menos 1 cita para que esto pase
	assert.Len(t, lista, 1)
}

func TestObtenerCita(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/citas/1", "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/citas/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
	t.Run("id no numerico -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/citas/abc", "", token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestCrearCita(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 201", func(t *testing.T) {
		body := `{"solicitante_id":"est_102","tecnico_id":"tec_05","estado":"pendiente","hora_acordada":"09:00","punto_encuentro":"Lab CISCO"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/citas", body, token))
		assert.Equal(t, http.StatusTeapot, rec.Code)
		var creada models.Cita
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))
		assert.NotZero(t, creada.ID)
	})
	t.Run("campos vacios -> 400", func(t *testing.T) {
		// Probamos enviando campos en blanco, lo que debe disparar tu ErrNombreVacio
		body := `{"solicitante_id":"   ","tecnico_id":"tec_05","hora_acordada":"","punto_encuentro":""}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/citas", body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("JSON malformado -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/citas", `{roto`, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestActualizarCita(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 200", func(t *testing.T) {
		body := `{"solicitante_id":"est_102","tecnico_id":"tec_05","estado":"completada","hora_acordada":"10:30","punto_encuentro":"Lab CISCO"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/citas/1", body, token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		body := `{"solicitante_id":"est_102","tecnico_id":"tec_05","estado":"pendiente","hora_acordada":"09:00","punto_encuentro":"Lab CISCO"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/citas/9999", body, token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestBorrarCita(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 204", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/citas/1", "", token))
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/citas/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

// El corazón de la seguridad: el middleware corta ANTES del handler.
func TestRutaProtegida_SinToken(t *testing.T) {
	h, _, _ := construirEntorno()
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/citas", "", "")) // sin Bearer
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestRutaProtegida_TokenInvalido(t *testing.T) {
	h, _, _ := construirEntorno()
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/citas", "", "token.falso.123"))
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

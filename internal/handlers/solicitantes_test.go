package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"proyecto/internal/models"
)

func TestListarSolicitantes_OK(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/solicitantes", "", token))

	require.Equal(t, http.StatusOK, rec.Code)
	var lista []models.Solicitante
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.Len(t, lista, 1) // el solicitante sembrado en construirEntorno
}

func TestObtenerSolicitante(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/solicitantes/1", "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/solicitantes/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestCrearSolicitante(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 201", func(t *testing.T) {
		// Ajustado a los campos reales de tu modelo Solicitante
		body := `{"nombre":"Jandry Cedeño","facultad":"FACCI","semestre":6,"nivel_urgencia":"alta"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/solicitantes", body, token))
		// Pon esto ANTES del require.Equal que falla
		if rec.Code != http.StatusCreated {
			t.Logf("Cuerpo de error del servidor: %s", rec.Body.String())
		}
		require.Equal(t, http.StatusCreated, rec.Code)
		require.Equal(t, http.StatusCreated, rec.Code)

		var creado models.Solicitante
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
		assert.NotZero(t, creado.ID)
	})
	t.Run("nombre vacio -> 400", func(t *testing.T) {
		body := `{"nombre":"  ","facultad":"FACCI"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/solicitantes", body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestActualizarSolicitante(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 200", func(t *testing.T) {
		body := `{"nombre":"Jandry Actualizado","facultad":"FACCI","semestre":6,"nivel_urgencia":"alta"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/solicitantes/1", body, token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestBorrarSolicitante(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 204", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/solicitantes/1", "", token))
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

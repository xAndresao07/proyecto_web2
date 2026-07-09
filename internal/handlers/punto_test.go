package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListarPuntos_OK(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/puntos-encuentro", "", token))
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestObtenerPunto(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/puntos-encuentro/1", "", token))
	assert.Equal(t, http.StatusNotFound, rec.Code) // o 200 si hay seed
}

func TestCrearPunto(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	body := `{"nombre_lugar":"Lab 1","facultad_perteneciente":"FACCI","disponible_para_soporte":true}`
	rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/puntos-encuentro", body, token))
	require.Equal(t, http.StatusCreated, rec.Code)
}

func TestActualizarPunto(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	body := `{"nombre_lugar":"Lab 1","facultad_perteneciente":"FACCI","disponible_para_soporte":true}`
	rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/puntos-encuentro/9999", body, token))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarPunto(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/puntos-encuentro/9999", "", token))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

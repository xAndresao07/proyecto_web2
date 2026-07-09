package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListarSoportes_OK(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/soportes", "", token))
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestObtenerSoporte(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/soportes/1", "", token))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCrearSoporte(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	body := `{"cita_id":1,"dispositivo_id":1,"solucion":"Listo","piezas_cambiadas":"Ninguna"}`
	rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/soportes", body, token))
	require.Equal(t, http.StatusCreated, rec.Code)
}

func TestActualizarSoporte(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	body := `{"cita_id":1,"dispositivo_id":1,"solucion":"Listo","piezas_cambiadas":"Ninguna"}`
	rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/soportes/9999", body, token))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestBorrarSoporte(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)
	rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/soportes/9999", "", token))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

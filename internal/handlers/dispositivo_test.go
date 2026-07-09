package handlers_test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"proyecto/internal/models"
)

func TestCrearDispositivo(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 201", func(t *testing.T) {
		body := `{"id":100, "solicitante_id":1, "marca":"HP", "modelo":"Pavilion"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/dispositivos", body, token))
		assert.Equal(t, http.StatusCreated, rec.Code)
		
		var resp models.Dispositivo
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
		assert.Equal(t, "HP", resp.Marca)
		assert.NotZero(t, resp.ID)
	})

	t.Run("JSON malformado -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/dispositivos", `{roto`, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
    
    t.Run("validacion falla -> 400", func(t *testing.T) {
		// Falla porque le falta el solicitante_id, o la marca
		body := `{"modelo":"Pavilion"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/dispositivos", body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestObtenerDispositivos(t *testing.T) {
	h, almacen, _ := construirEntorno()
	token := tokenValido(t, h)
    
	// Crear uno manual
	creado := almacen.CrearDispositivo(models.Dispositivo{SolicitanteID: 1, Marca: "Dell", Modelo: "XPS"})
    idStr := strconv.Itoa(int(creado.ID))

	t.Run("listar todos -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/dispositivos", "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var lista []models.Dispositivo
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
		assert.NotEmpty(t, lista)
	})

	t.Run("obtener uno valido -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/dispositivos/"+idStr, "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
    
    t.Run("no encontrado -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/dispositivos/999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
    
    t.Run("id invalido -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/dispositivos/abc", "", token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestActualizarDispositivo(t *testing.T) {
	h, almacen, _ := construirEntorno()
	token := tokenValido(t, h)
	creado := almacen.CrearDispositivo(models.Dispositivo{SolicitanteID: 1, Marca: "Asus", Modelo: "ROG"})
	idStr := strconv.Itoa(int(creado.ID))

	t.Run("valido -> 200", func(t *testing.T) {
		body := `{"id":`+idStr+`, "solicitante_id":1, "marca":"Asus", "modelo":"TUF"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/dispositivos/"+idStr, body, token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("no encontrado -> 404", func(t *testing.T) {
		body := `{"id":999, "solicitante_id":1, "marca":"Asus", "modelo":"TUF"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/dispositivos/999", body, token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestBorrarDispositivo(t *testing.T) {
	h, almacen, _ := construirEntorno()
	token := tokenValido(t, h)
	creado := almacen.CrearDispositivo(models.Dispositivo{SolicitanteID: 1, Marca: "Lenovo", Modelo: "Thinkpad"})
	idStr := strconv.Itoa(int(creado.ID))

	t.Run("valido -> 204", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/dispositivos/"+idStr, "", token))
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("no encontrado -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/dispositivos/999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

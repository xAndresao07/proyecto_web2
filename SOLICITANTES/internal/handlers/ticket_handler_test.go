package handlers_test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"solicitantesYHardware/internal/models"
)

func TestCrearTicket(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 201", func(t *testing.T) {
		body := `{"id":100, "solicitante_id":1, "dispositivo_id":1, "descripcion_falla":"falla", "software_requerido":"ninguno"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/tickets", body, token))
		assert.Equal(t, http.StatusCreated, rec.Code)
		
		var resp models.TicketAyuda
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
		assert.Equal(t, "falla", resp.DescripcionFalla)
		assert.Equal(t, "abierto", resp.EstadoTicket) // El service lo fuerza a abierto
		assert.NotZero(t, resp.ID)
	})

	t.Run("JSON malformado -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/tickets", `{roto`, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
    
    t.Run("validacion falla -> 400", func(t *testing.T) {
		// Falla porque le falta el dispositivo_id
		body := `{"solicitante_id":1, "descripcion_falla":"falla"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/tickets", body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestObtenerTickets(t *testing.T) {
	h, almacen, _ := construirEntorno()
	token := tokenValido(t, h)
    
	// Crear uno manual
	creado := almacen.CrearTicketAyuda(models.TicketAyuda{SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "Falla PC"})
    idStr := strconv.Itoa(int(creado.ID))

	t.Run("listar todos -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/tickets", "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var lista []models.TicketAyuda
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
		assert.NotEmpty(t, lista)
	})

	t.Run("obtener uno valido -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/tickets/"+idStr, "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
    
    t.Run("no encontrado -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/tickets/999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
    
    t.Run("id invalido -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/tickets/abc", "", token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestActualizarTicket(t *testing.T) {
	h, almacen, _ := construirEntorno()
	token := tokenValido(t, h)
	creado := almacen.CrearTicketAyuda(models.TicketAyuda{SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "Test"})
	idStr := strconv.Itoa(int(creado.ID))

	t.Run("valido -> 200", func(t *testing.T) {
		body := `{"id":`+idStr+`, "solicitante_id":1, "dispositivo_id":1, "descripcion_falla":"Falla editada", "estado_ticket":"cerrado"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/tickets/"+idStr, body, token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("estado invalido -> 400", func(t *testing.T) {
		body := `{"id":`+idStr+`, "solicitante_id":1, "dispositivo_id":1, "descripcion_falla":"Falla editada", "estado_ticket":"invalido"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/tickets/"+idStr, body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("no encontrado -> 404", func(t *testing.T) {
		body := `{"id":999, "solicitante_id":1, "dispositivo_id":1, "descripcion_falla":"Falla editada", "estado_ticket":"abierto"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/tickets/999", body, token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestBorrarTicket(t *testing.T) {
	h, almacen, _ := construirEntorno()
	token := tokenValido(t, h)
	creado := almacen.CrearTicketAyuda(models.TicketAyuda{SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "Borrar"})
	idStr := strconv.Itoa(int(creado.ID))

	t.Run("valido -> 204", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/tickets/"+idStr, "", token))
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("no encontrado -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/tickets/999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

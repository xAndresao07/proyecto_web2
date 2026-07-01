package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"solicitantesYHardware/internal/models"
)

// ListarTickets atiende GET /api/v1/tickets.
func (s *Server) ListarTickets(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.TicketAyudas.Listar())
}

// ObtenerTicket atiende GET /api/v1/tickets/{id}.
func (s *Server) ObtenerTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	t, err := s.TicketAyudas.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, t)
}

// CrearTicket atiende POST /api/v1/tickets.
func (s *Server) CrearTicket(w http.ResponseWriter, r *http.Request) {
	var nuevo models.TicketAyuda
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.TicketAyudas.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarTicket atiende PUT /api/v1/tickets/{id}.
func (s *Server) ActualizarTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.TicketAyuda
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, err := s.TicketAyudas.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarTicket atiende DELETE /api/v1/tickets/{id}.
func (s *Server) BorrarTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.TicketAyudas.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

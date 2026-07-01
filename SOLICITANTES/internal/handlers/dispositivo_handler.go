package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"solicitantesYHardware/internal/models"
)

// ListarDispositivos atiende GET /api/v1/dispositivos.
func (s *Server) ListarDispositivos(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Dispositivos.Listar())
}

// ObtenerDispositivo atiende GET /api/v1/dispositivos/{id}.
func (s *Server) ObtenerDispositivo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	d, err := s.Dispositivos.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, d)
}

// CrearDispositivo atiende POST /api/v1/dispositivos.
func (s *Server) CrearDispositivo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Dispositivo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Dispositivos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarDispositivo atiende PUT /api/v1/dispositivos/{id}.
func (s *Server) ActualizarDispositivo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Dispositivo
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, err := s.Dispositivos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarDispositivo atiende DELETE /api/v1/dispositivos/{id}.
func (s *Server) BorrarDispositivo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Dispositivos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

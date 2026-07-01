// Package handlers contiene los handlers HTTP del módulo Solicitantes y Hardware.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"solicitantesYHardware/internal/models"
)

// ListarSolicitantes atiende GET /api/v1/solicitantes.
func (s *Server) ListarSolicitantes(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Solicitantes.Listar())
}

// ObtenerSolicitante atiende GET /api/v1/solicitantes/{id}.
func (s *Server) ObtenerSolicitante(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	sol, err := s.Solicitantes.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, sol)
}

// CrearSolicitante atiende POST /api/v1/solicitantes.
func (s *Server) CrearSolicitante(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Solicitante
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Solicitantes.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarSolicitante atiende PUT /api/v1/solicitantes/{id}.
func (s *Server) ActualizarSolicitante(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Solicitante
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, err := s.Solicitantes.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarSolicitante atiende DELETE /api/v1/solicitantes/{id}.
func (s *Server) BorrarSolicitante(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Solicitantes.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

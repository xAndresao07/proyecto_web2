package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"

	"github.com/go-chi/chi/v5"
)

// GET /api/v1/tecnicos
func (s *Server) GetAllTecnicos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Tecnicos.Listar())
}

// POST /api/v1/tecnicos
func (s *Server) CreateTecnico(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Tecnicos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

// GET /api/v1/tecnicos/{id}
func (s *Server) GetTecnicoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID debe ser numérico")
		return
	}

	tecnico, err := s.Tecnicos.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, tecnico)
}

// PUT /api/v1/tecnicos/{id}
func (s *Server) UpdateTecnico(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID debe ser numérico")
		return
	}

	var datos models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, err := s.Tecnicos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// DELETE /api/v1/tecnicos/{id}
func (s *Server) DeleteTecnico(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID debe ser numérico")
		return
	}

	if err := s.Tecnicos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

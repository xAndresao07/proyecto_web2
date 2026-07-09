package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto/internal/models"
)

func (s *Server) ListarSolicitantes(w http.ResponseWriter, _ *http.Request) {
	solicitantes := s.Solicitantes.Listar()
	RespondJSON(w, http.StatusOK, solicitantes)
}

func (s *Server) ObtenerSolicitante(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	solicitante, err := s.Solicitantes.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, solicitante)
}

func (s *Server) CrearSolicitante(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Solicitante
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	creado, err := s.Solicitantes.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarSolicitante(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	var datos models.Solicitante
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	actualizado, err := s.Solicitantes.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarSolicitantes(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	if err := s.Solicitantes.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

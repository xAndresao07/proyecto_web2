package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto/internal/models"
)

func (s *Server) ListarCitas(w http.ResponseWriter, _ *http.Request) {
	citas := s.Citas.Listar()
	RespondJSON(w, http.StatusOK, citas)
}

func (s *Server) ObtenerCita(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	cita, err := s.Citas.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cita)
}

func (s *Server) CrearCita(w http.ResponseWriter, r *http.Request) {
	var nueva models.Cita
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	creada, err := s.Citas.Crear(nueva)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarCita(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	var datos models.Cita
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	actualizada, err := s.Citas.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarCita(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	if err := s.Citas.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Server) ListarPuntos(w http.ResponseWriter, _ *http.Request) {
	puntos := s.PuntosEncuentro.Listar()
	RespondJSON(w, http.StatusOK, puntos)
}

func (s *Server) ObtenerPunto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	punto, err := s.PuntosEncuentro.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, punto)
}

func (s *Server) CrearPunto(w http.ResponseWriter, r *http.Request) {
	var nuevo models.PuntoEncuentro
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	creado, err := s.PuntosEncuentro.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarPunto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	var datos models.PuntoEncuentro
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	actualizado, err := s.PuntosEncuentro.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarPunto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	if err := s.PuntosEncuentro.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto/internal/models"
)

func (s *Server) ListarDispositivos(w http.ResponseWriter, _ *http.Request) {
	dispositivos := s.Dispositivos.Listar()
	RespondJSON(w, http.StatusOK, dispositivos)
}

func (s *Server) ObtenerDispositivo(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	dispositivo, err := s.Dispositivos.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, dispositivo)
}

func (s *Server) CrearDispositivo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Dispositivo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	creado, err := s.Dispositivos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarDispositivo(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	var datos models.Dispositivo
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}

	actualizado, err := s.Dispositivos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarDispositivo(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}

	if err := s.Dispositivos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

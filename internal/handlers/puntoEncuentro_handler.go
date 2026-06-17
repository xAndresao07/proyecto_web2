package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"proyecto/internal/models"

	"github.com/go-chi/chi/v5"
)

// =========================================================
// HANDLERS: PUNTOS DE ENCUENTRO
// =========================================================

func (s *Server) ListarPuntos(w http.ResponseWriter, _ *http.Request) {
	puntos := s.storage.ListarPuntos()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(puntos)
}

func (s *Server) CrearPunto(w http.ResponseWriter, r *http.Request) {
	var nuevo models.PuntoEncuentro
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(nuevo.NombreLugar) == "" || strings.TrimSpace(nuevo.FacultadPerteneciente) == "" {
		http.Error(w, "nombre_lugar y facultad_perteneciente son obligatorios", http.StatusBadRequest)
		return
	}

	nuevo.DisponibleParaSoporte = true

	nuevo = s.storage.CrearPunto(nuevo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

func (s *Server) ObtenerPunto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un entero válido", http.StatusBadRequest)
		return
	}

	punto, encontrado := s.storage.BuscarPuntoPorID(id)
	if !encontrado {
		http.Error(w, "Punto de encuentro no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(punto)
}

func (s *Server) ActualizarPunto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un entero válido", http.StatusBadRequest)
		return
	}

	var datos models.PuntoEncuentro
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.storage.ActualizarPunto(id, datos)
	if !encontrado {
		http.Error(w, "Punto no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizado)
}

func (s *Server) BorrarPunto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un entero válido", http.StatusBadRequest)
		return
	}

	if !s.storage.EliminarPunto(id) {
		http.Error(w, "Punto no encontrado para eliminar", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

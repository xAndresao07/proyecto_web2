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
// HANDLERS: SOPORTES
// =========================================================

func (s *Server) ListarSoportes(w http.ResponseWriter, _ *http.Request) {
	soportes := s.storage.ListarSoportes()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(soportes)
}

func (s *Server) CrearSoporte(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Soporte
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if nuevo.CitaID <= 0 || nuevo.DispositivoID <= 0 || strings.TrimSpace(nuevo.Solucion) == "" {
		http.Error(w, "cita_id, dispositivo_id y solucion son campos obligatorios", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(nuevo.PiezasCambiadas) == "" {
		nuevo.PiezasCambiadas = "Ninguna"
	}

	nuevo = s.storage.CrearSoporte(nuevo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

func (s *Server) ObtenerSoporte(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un entero válido", http.StatusBadRequest)
		return
	}

	soporte, encontrado := s.storage.BuscarSoportePorID(id)
	if !encontrado {
		http.Error(w, "Soporte no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(soporte)
}

func (s *Server) ActualizarSoporte(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un entero válido", http.StatusBadRequest)
		return
	}

	var datos models.Soporte
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.storage.ActualizarSoporte(id, datos)
	if !encontrado {
		http.Error(w, "Soporte no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizado)
}

func (s *Server) BorrarSoporte(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un entero válido", http.StatusBadRequest)
		return
	}

	if !s.storage.EliminarSoporte(id) {
		http.Error(w, "Soporte no encontrado para eliminar", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

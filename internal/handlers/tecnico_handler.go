package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"

	"github.com/go-chi/chi/v5"
)

// TecnicoServer agrupa los endpoints de tu módulo y guarda su propia dependencia.
type Server struct {
	storage *storage.Memoria
}

// NewTecnicoServer es el constructor que INYECTA la dependencia de almacenamiento.
func NewServer(s *storage.Memoria) *Server {
	return &Server{storage: s}
}

// 1. GET /api/v1/tecnicos
func (s *Server) GetAllTecnicos(w http.ResponseWriter, r *http.Request) {
	tecnicos := s.storage.ListarTecnicos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tecnicos)
}

// 2. POST /api/v1/tecnicos
func (s *Server) CreateTecnico(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(nuevo.Nombre) == "" || len(nuevo.Habilidades) == 0 {
		http.Error(w, "El nombre y al menos una habilidad son obligatorios", http.StatusBadRequest)
		return
	}

	nuevo.Reputacion = 5.0
	nuevo = s.storage.CrearTecnico(nuevo) // Usamos la capa inyectada

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

// 3. GET /api/v1/tecnicos/{id}
func (s *Server) GetTecnicoPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	tecnico, encontrado := s.storage.BuscarTecnicoPorID(idParam)
	if !encontrado {
		http.Error(w, "Técnico no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tecnico)
}

// 4. PUT /api/v1/tecnicos/{id}
func (s *Server) UpdateTecnico(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var datos models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.storage.ActualizarTecnico(idParam, datos)
	if !encontrado {
		http.Error(w, "Técnico no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizado)
}

// 5. DELETE /api/v1/tecnicos/{id}
func (s *Server) DeleteTecnico(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	seBorro := s.storage.EliminarTecnico(idParam)
	if !seBorro {
		http.Error(w, "Técnico no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

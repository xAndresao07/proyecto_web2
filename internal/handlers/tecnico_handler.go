package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/service"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Tecnicos *service.TecnicoService
}

func NewServer(tecnicos *service.TecnicoService) *Server {
	return &Server{Tecnicos: tecnicos}
}

// 1. GET /api/v1/tecnicos
func (s *Server) GetAllTecnicos(w http.ResponseWriter, r *http.Request) {
	tecnicos := s.Tecnicos.Listar()

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

	creado, err := s.Tecnicos.Crear(nuevo)
	if err != nil {
		// Evaluamos si el error viene de nuestras reglas de negocio
		if err == service.ErrNombreVacio || err == service.ErrSinServicios {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// 3. GET /api/v1/tecnicos/{id}
func (s *Server) GetTecnicoPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID debe ser numérico", http.StatusBadRequest)
		return
	}

	tecnico, err := s.Tecnicos.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tecnico)
}

// 4. PUT /api/v1/tecnicos/{id}
func (s *Server) UpdateTecnico(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID debe ser numérico", http.StatusBadRequest)
		return
	}

	var datos models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, err := s.Tecnicos.Actualizar(id, datos)
	if err != nil {
		if err == service.ErrNoEncontrado {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizado)
}

// 5. DELETE /api/v1/tecnicos/{id}
func (s *Server) DeleteTecnico(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID debe ser numérico", http.StatusBadRequest)
		return
	}

	if err := s.Tecnicos.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

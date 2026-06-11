package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

// Server unifica los endpoints de intervenciones y recibe su almacenamiento (Memoria)
type Server struct {
	storage *storage.Memoria
}

// NewServer es el constructor que inyecta la base de datos en el controlador
func NewServer(s *storage.Memoria) *Server {
	return &Server{storage: s}
}

// 1. Listar todas las intervenciones (GET)
func (s *Server) ListarIntervenciones(w http.ResponseWriter, _ *http.Request) {
	intervenciones := s.storage.ListarIntervenciones()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(intervenciones); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 2. Crear una intervención (POST)
func (s *Server) CrearIntervencion(w http.ResponseWriter, r *http.Request) {
	var nueva models.Intervencion
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validación de campos obligatorios básicos para el Hito 2
	if strings.TrimSpace(nueva.SolicitanteID) == "" || strings.TrimSpace(nueva.TecnicoID) == "" || strings.TrimSpace(nueva.HoraAcordada) == "" || strings.TrimSpace(nueva.PuntoEncuentro) == "" {
		http.Error(w, "SolicitanteID, TecnicoID, HoraAcordada y PuntoEncuentro son obligatorios", http.StatusBadRequest)
		return
	}

	nueva.Estado = "pendiente"

	nueva = s.storage.CrearIntervencion(nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(nueva); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 3. Obtener intervención por ID (GET)
func (s *Server) ObtenerIntervencion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convertimos el ID de texto de la URL a entero (int)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest) // 400 Bad Request
		return
	}

	intervencion, encontrado := s.storage.BuscarIntervencionPorID(id)
	if !encontrado {
		http.Error(w, "intervención no encontrada", http.StatusNotFound) // 404 Not Found
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(intervencion); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 4. Actualizar intervención (PUT)
func (s *Server) ActualizarIntervencion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Intervencion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(datos.Estado) == "" && strings.TrimSpace(datos.HoraAcordada) == "" && strings.TrimSpace(datos.PuntoEncuentro) == "" {
		http.Error(w, "debe enviar al menos un campo para actualizar", http.StatusBadRequest)
		return
	}

	actualizada, encontrado := s.storage.ActualizarIntervencion(id, datos)
	if !encontrado {
		http.Error(w, "intervención no encontrada para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 5. Eliminar intervención (DELETE)
func (s *Server) BorrarIntervencion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	seBorro := s.storage.EliminarIntervencion(id)
	if !seBorro {
		http.Error(w, "intervención no encontrada para eliminar", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

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
func (s *Server) ListarCitas(w http.ResponseWriter, _ *http.Request) {
	citas := s.storage.ListarCitas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(citas); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 2. Crear una cita (POST)
func (s *Server) CrearCita(w http.ResponseWriter, r *http.Request) {
	var nueva models.Cita
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

	nueva = s.storage.CrearCita(nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(nueva); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 3. Obtener cita por ID (GET)
func (s *Server) ObtenerCita(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convertimos el ID de texto de la URL a entero (int)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest) // 400 Bad Request
		return
	}

	cita, encontrado := s.storage.BuscarCitaPorID(id)
	if !encontrado {
		http.Error(w, "cita no encontrada", http.StatusNotFound) // 404 Not Found
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cita); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 4. Actualizar cita (PUT)
func (s *Server) ActualizarCita(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Cita
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(datos.Estado) == "" && strings.TrimSpace(datos.HoraAcordada) == "" && strings.TrimSpace(datos.PuntoEncuentro) == "" {
		http.Error(w, "debe enviar al menos un campo para actualizar", http.StatusBadRequest)
		return
	}

	actualizada, encontrado := s.storage.ActualizarCita(id, datos)
	if !encontrado {
		http.Error(w, "cita no encontrada para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// 5. Eliminar cita (DELETE)
func (s *Server) EliminarCita(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	seBorro := s.storage.EliminarCita(id)
	if !seBorro {
		http.Error(w, "cita no encontrada para eliminar", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

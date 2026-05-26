package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/cmd/internal/models"
	"proyecto/cmd/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Listamos todas las intervenciones (GET /intervenciones)

func GetAllIntervenciones(w http.ResponseWriter, r *http.Request) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.Intervenciones)
}

// Creamos una intervención nueva(POST /intervenciones)

func CreateIntervencion(w http.ResponseWriter, r *http.Request) {
	var nueva models.Intervencion
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "Error decodificando la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	if nueva.SolicitanteID == "" || nueva.TecnicoID == "" || nueva.HoraAcordada == "" || nueva.PuntoEncuentro == "" {
		http.Error(w, "SolicitanteID, TecnicoID, HoraAcordada y PuntoEncuentro son obligatorios", http.StatusBadRequest) // error 400 Bad Request
	}

	// Forzamos el estado inicial, por si el cliente intenta mandar otro estado
	nueva.Estado = "pendiente"

	storage.Mu.Lock()
	nueva.ID = strconv.Itoa(storage.SiguienteID) // Asignamos un ID único
	storage.SiguienteID++
	storage.Intervenciones = append(storage.Intervenciones, nueva)
	storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(nueva)
}

//Obtener una intervención por ID (GET /intervenciones/{id})

func GetIntervencionPorID(w http.ResponseWriter, r *http.Request) {
	// Usamos Chi para extraer el ID directamente de la URL
	idParam := chi.URLParam(r, "id")

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for _, intervencion := range storage.Intervenciones {
		if intervencion.ID == idParam {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // 200 OK
			json.NewEncoder(w).Encode(intervencion)
			return
		}
	}
	// Si termina el bucle y no la encuentra, enviamos 404
	http.Error(w, "Intervención no encontrada", http.StatusNotFound)
}

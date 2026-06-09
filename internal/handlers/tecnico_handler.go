package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strconv"
)

func GetAllTecnicos(w http.ResponseWriter, r *http.Request) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.Tecnicos)
}

func CreateTecnico(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Error decodificando: "+err.Error(), http.StatusBadRequest)
		return
	}

	if nuevo.Nombre == "" || len(nuevo.Habilidades) == 0 {
		http.Error(w, "El nombre y al menos una habilidad son obligatorios", http.StatusBadRequest)
		return
	}

	storage.Mu.Lock()
	nuevo.ID = strconv.Itoa(storage.SiguienteID)
	storage.SiguienteID++
	storage.Tecnicos = append(storage.Tecnicos, nuevo)
	storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

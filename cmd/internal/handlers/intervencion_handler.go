package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/cmd/internal/storage"
)

// Listamos todas las intervenciones

func GetAllIntervenciones(w http.ResponseWriter, r *http.Request) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.Intervenciones)
}

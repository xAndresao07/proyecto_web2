package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/storage"
)

func GetAllTecnicos(w http.ResponseWriter, r *http.Request) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.Tecnicos)
}

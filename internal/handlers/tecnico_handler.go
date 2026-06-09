package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	nuevo.Reputacion = 5.0
	storage.Mu.Lock()
	nuevo.ID = strconv.Itoa(storage.SiguienteID)
	storage.SiguienteID++
	storage.Tecnicos = append(storage.Tecnicos, nuevo)
	storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

func GetTecnicoPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for _, tecnico := range storage.Tecnicos {
		if tecnico.ID == idParam {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tecnico)
			return
		}
	}
	http.Error(w, "Técnico no encontrado", http.StatusNotFound)
}

func UpdateTecnico(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var datosActualizados models.Tecnico
	if err := json.NewDecoder(r.Body).Decode(&datosActualizados); err != nil {
		http.Error(w, "Error decodificando", http.StatusBadRequest)
		return
	}

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for i, tecnico := range storage.Tecnicos {
		if tecnico.ID == idParam {
			datosActualizados.ID = idParam

			// Actualización parcial
			if datosActualizados.Nombre == "" {
				datosActualizados.Nombre = tecnico.Nombre
			}
			if len(datosActualizados.Habilidades) == 0 {
				datosActualizados.Habilidades = tecnico.Habilidades
			}
			if datosActualizados.HorarioLibre == "" {
				datosActualizados.HorarioLibre = tecnico.HorarioLibre
			}
			if datosActualizados.Reputacion == 0 {
				datosActualizados.Reputacion = tecnico.Reputacion
			}

			storage.Tecnicos[i] = datosActualizados
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(datosActualizados)
			return
		}
	}
	http.Error(w, "Técnico no encontrado", http.StatusNotFound)
}

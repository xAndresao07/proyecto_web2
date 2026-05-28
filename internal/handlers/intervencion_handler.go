package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/models"
	"proyecto/internal/storage"
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

//Actualizar intervencion (PUT /intervenciones/{id})

func UpdateIntervencion(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var datosActualizados models.Intervencion
	if err := json.NewDecoder(r.Body).Decode(&datosActualizados); err != nil {
		http.Error(w, "Error decodificando la solicitud: ", http.StatusBadRequest)
		return
	}

	if datosActualizados.Estado == "" && datosActualizados.HoraAcordada == "" {
		http.Error(w, "Debe enviar al menos un campo para actualizar (Estado u HoraAcordada)", http.StatusBadRequest)
		return
	}

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for i, intervencion := range storage.Intervenciones {
		if intervencion.ID == idParam {
			// Actualizamos solo los datos permitidos, sin tocar los IDs de los usuarios
			datosActualizados.ID = idParam
			datosActualizados.SolicitanteID = intervencion.SolicitanteID
			datosActualizados.TecnicoID = intervencion.TecnicoID

			if datosActualizados.PuntoEncuentro == "" {
				datosActualizados.PuntoEncuentro = intervencion.PuntoEncuentro
			}
			if datosActualizados.Estado == "" {
				datosActualizados.Estado = intervencion.Estado
			}
			if datosActualizados.HoraAcordada == "" {
				datosActualizados.HoraAcordada = intervencion.HoraAcordada
			}

			// Reemplazamos en el slice usando el índice 'i'
			storage.Intervenciones[i] = datosActualizados

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(datosActualizados)
			return
		}
	}
	http.Error(w, "Intervencion no encontrada para poderla actualizar", http.StatusNotFound) // 404 Not Found
}

// Eliminar una intervención (DELETE /intervenciones/{id})

func DeleteIntervencion(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for i, intervencion := range storage.Intervenciones {
		if intervencion.ID == idParam {
			// Rebanamos el slice: tomamos lo de antes del elemento 'i' y lo unimos con lo de después
			storage.Intervenciones = append(storage.Intervenciones[:i], storage.Intervenciones[i+1:]...)

			w.WriteHeader(http.StatusNoContent) // 204 No Content (Borrado exitoso)
			return
		}
	}
	http.Error(w, "Intervencion no encontrada para poderla eliminar", http.StatusNotFound) // 404 not found
}

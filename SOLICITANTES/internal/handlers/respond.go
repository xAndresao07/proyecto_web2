package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"solicitantesYHardware/internal/service"
)

// RespondJSON escribe data como JSON con el status HTTP indicado.
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// RespondError escribe un error en formato JSON consistente: {"error": "..."}.
func RespondError(w http.ResponseWriter, status int, mensaje string) {
	RespondJSON(w, status, map[string]string{"error": mensaje})
}

func statusDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusNotFound
	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized
	case errors.Is(err, service.ErrEmailEnUso):
		return http.StatusConflict

	case errors.Is(err, service.ErrNombreVacio),
		errors.Is(err, service.ErrMatriculaVacia),
		errors.Is(err, service.ErrFacultadVacia),
		errors.Is(err, service.ErrSemestreInvalido),
		errors.Is(err, service.ErrNivelUrgenciaInvalido),
		errors.Is(err, service.ErrMarcaVacia),
		errors.Is(err, service.ErrModeloVacio),
		errors.Is(err, service.ErrTipoAlmacenamientoInvalido),
		errors.Is(err, service.ErrRamInvalida),
		errors.Is(err, service.ErrSistemaOperativoVacio),
		errors.Is(err, service.ErrSolicitanteIDInvalido),
		errors.Is(err, service.ErrDescripcionFallaVacia),
		errors.Is(err, service.ErrDispositivoIDInvalido),
		errors.Is(err, service.ErrEstadoTicketInvalido):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}

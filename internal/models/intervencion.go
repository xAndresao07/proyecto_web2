package models

// Intervencion representa la cita técnica entre un estudiante y un compañero que le brindará ayuda.
type Intervencion struct {
	ID             int    `json:"id"`
	SolicitanteID  string `json:"solicitante_id"`
	TecnicoID      string `json:"tecnico_id"`
	Estado         string `json:"estado"`
	HoraAcordada   string `json:"hora_acordada"`
	PuntoEncuentro string `json:"punto_encuentro"`
}

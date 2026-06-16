package models

type HorarioTecnico struct {
	DiaSemana            string `json:"dia_semana"`            // Ej: Lunes
	HoraInicio           string `json:"hora_inicio"`           // Ej: 14:00
	HoraFin              string `json:"hora_fin"`              // Ej: 16:00
	EstadoDisponibilidad string `json:"estado_disponibilidad"` // Ej: disponible, ocupado
}

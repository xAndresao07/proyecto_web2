package models

type HorarioTecnico struct {
	ID                   int    `json:"id" gorm:"primaryKey"`
	TecnicoID            int    `json:"tecnico_id"`
	DiaSemana            string `json:"dia_semana" gorm:"not null"`
	HoraInicio           string `json:"hora_inicio" gorm:"not null"`
	HoraFin              string `json:"hora_fin" gorm:"not null"`
	EstadoDisponibilidad string `json:"estado_disponibilidad"`
}
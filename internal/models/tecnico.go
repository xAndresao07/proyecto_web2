package models

type Tecnico struct {
	ID         string             `json:"id"`
	Nombre     string             `json:"nombre"`
	Reputacion float64            `json:"reputacion"`
	Servicios  []ServicioOfrecido `json:"servicios"` // Un técnico tiene múltiples servicios
	Horarios   []HorarioTecnico   `json:"horarios"`  // Un técnico tiene múltiples bloques de horario
}

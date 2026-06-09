package models

type Tecnico struct {
	ID           string   `json:"id"`
	Nombre       string   `json:"nombre"`
	Habilidades  []string `json:"habilidades"`
	HorarioLibre string   `json:"horario_libre"`
	Reputacion   float64  `json:"reputacion"`
}

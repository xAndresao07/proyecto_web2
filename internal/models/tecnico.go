package models

type Tecnico struct {
	ID         int                `json:"id" gorm:"primaryKey"`
	Nombree    string             `json:"nombre" gorm:"not null"`
	Reputacion float64            `json:"reputacion"`
	Servicios  []ServicioOfrecido `json:"servicios" gorm:"foreignKey:TecnicoID"`
	Horarios   []HorarioTecnico   `json:"horarios" gorm:"foreignKey:TecnicoID"`
}

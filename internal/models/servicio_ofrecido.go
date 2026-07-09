package models

type ServicioOfrecido struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	TecnicoID        int    `json:"tecnico_id"`
	NombreServicio   string `json:"nombre_servicio" gorm:"not null"`
	NivelExperiencia string `json:"nivel_experiencia"`
	TiempoEstimado   string `json:"tiempo_estimado"`
}
package models

// Cita representa la cita técnica entre un estudiante y un compañero que le brindará ayuda.

type Cita struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	SolicitanteID  string `json:"solicitante_id"`
	TecnicoID      string `json:"tecnico_id"`
	Estado         string `json:"estado"`
	HoraAcordada   string `json:"hora_acordada"`
	PuntoEncuentro string `json:"punto_encuentro"`

	// Relación Has-Many: Una Cita tiene Soportes asociados
	Soportes []Soporte `json:"soportes,omitempty" gorm:"foreignKey:CitaID"`
}

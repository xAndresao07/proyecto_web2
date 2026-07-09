package models

// Soporte representa la auditoría y cierre del ciclo una vez terminada la cita.

type Soporte struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	CitaID          int    `json:"cita_id"`
	DispositivoID   int    `json:"dispositivo_id"`
	Solucion        string `json:"solucion"`
	PiezasCambiadas string `json:"piezas_cambiadas"`

	// Relación Belongs-To: Este soporte pertenece a una Cita
	Cita *Cita `json:"cita,omitempty" gorm:"foreignKey:CitaID"`
}

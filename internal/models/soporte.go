package models

// Soporte representa la auditoría y cierre del ciclo una vez terminada la cita.
type Soporte struct {
	ID              int    `json:"id"`
	CitaID          int    `json:"cita_id"`          // Conecta con la Cita finalizada
	DispositivoID   int    `json:"dispositivo_id"`   // Conecta con el dispositivo del módulo de Jandry
	Solucion        string `json:"solucion"`         // Detalle técnico de lo que se reparó
	PiezasCambiadas string `json:"piezas_cambiadas"` // Detalle de las piezas que se cambiaron, si hubo
}

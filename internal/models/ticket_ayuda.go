package models

// TicketAyuda representa el reporte de falla creado por el solicitante sobre un dispositivo.

type TicketAyuda struct {
	ID                int    `json:"id" gorm:"primaryKey"` // Reporte # incremental
	SolicitanteID     int    `json:"solicitante_id"`
	DispositivoID     int    `json:"dispositivo_id"`
	DescripcionFalla  string `json:"descripcion_falla"`
	SoftwareRequerido string `json:"software_requerido"`
	EstadoTicket      string `json:"estado_ticket"`
}

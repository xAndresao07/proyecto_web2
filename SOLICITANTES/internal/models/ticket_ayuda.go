package models

// TicketAyuda es el reporte formal del problema, previo a que exista la
// intervención técnica.
//
// Se asocia a un Solicitante y a un Dispositivo específico de ese
// solicitante (SolicitanteID y DispositivoID).
type TicketAyuda struct {
	ID                int    `json:"id" gorm:"primaryKey"`
	SolicitanteID     int    `json:"solicitante_id" gorm:"not null"`
	DispositivoID     int    `json:"dispositivo_id" gorm:"not null"`
	DescripcionFalla  string `json:"descripcion_falla" gorm:"not null"`
	SoftwareRequerido string `json:"software_requerido"`
	EstadoTicket      string `json:"estado_ticket" gorm:"not null;default:abierto"`

	// Relaciones (Belongs-To)
	Solicitante *Solicitante `json:"solicitante,omitempty" gorm:"foreignKey:SolicitanteID"`
	Dispositivo *Dispositivo `json:"dispositivo,omitempty" gorm:"foreignKey:DispositivoID"`

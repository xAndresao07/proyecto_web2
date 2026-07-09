// Package models define las entidades del dominio de Solicitantes y Hardware.
package models

// Solicitante representa el perfil del estudiante que necesita ayuda técnica.
//
// UsuarioID referencia el ID de un Usuario (1:1) — quien se autentica.
type Solicitante struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	UsuarioID     int    `json:"usuario_id" gorm:"not null;uniqueIndex"`
	Matricula     string `json:"matricula" gorm:"unique;not null"`
	Nombre        string `json:"nombre" gorm:"not null"`
	Facultad      string `json:"facultad" gorm:"not null"`
	Semestre      int    `json:"semestre" gorm:"not null"`
	NivelUrgencia string `json:"nivel_urgencia" gorm:"not null;default:normal"`

	// Relaciones (Has-Many)
	Dispositivos []Dispositivo `json:"dispositivos,omitempty" gorm:"foreignKey:SolicitanteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tickets      []TicketAyuda `json:"tickets,omitempty" gorm:"foreignKey:SolicitanteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

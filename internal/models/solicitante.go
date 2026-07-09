package models

// Solicitante representa al estudiante que reporta una falla de hardware/software.

type Solicitante struct {
	ID            int    `json:"id" gorm:"primaryKey"` // ULEAM-matricula
	Nombre        string `json:"nombre"`
	Facultad      string `json:"facultad"`
	Semestre      int    `json:"semestre"`
	NivelUrgencia string `json:"nivel_urgencia"`
}

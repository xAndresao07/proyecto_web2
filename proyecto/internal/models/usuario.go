package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Rol          string    `json:"rol" gorm:"not null"` // admin, solicitante, tecnico
	CreadoEn     time.Time `json:"creado_en"`
}

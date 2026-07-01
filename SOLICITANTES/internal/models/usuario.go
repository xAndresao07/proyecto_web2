package models

import "time"

// Usuario es la cuenta de acceso al sistema (auth compartido del proyecto).
type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreadoEn     time.Time `json:"creado_en"`
}

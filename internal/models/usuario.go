package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement:false"` // <-- ¡El secreto está aquí!
	Email        string    `json:"email" gorm:"not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"column:password;not null"`
	Rol          string    `json:"rol" gorm:"not null"`
	CreadoEn     time.Time `json:"creado_en" gorm:"-"`
}

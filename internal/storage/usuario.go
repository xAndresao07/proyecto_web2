package storage

import (
	"proyecto/internal/models"
	"time"

	"gorm.io/gorm"
)

type UsuarioGORM struct{ db *gorm.DB }

func NewUsuarioRepository(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

func (r *UsuarioGORM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.CreadoEn = time.Now()
	return u, r.db.Create(&u).Error
}
func (r *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

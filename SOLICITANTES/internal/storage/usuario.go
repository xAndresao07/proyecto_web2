package storage

import (
	"time"

	"gorm.io/gorm"

	"solicitantesYHardware/internal/models"
)

// UsuarioGORM implementa UserRepositorio usando GORM. El auth siempre
// corre sobre GORM, sin importar qué backend (Memoria/SQLite/SQLC) sirva
// las peticiones de Solicitantes, Dispositivos y TicketAyudas.
type UsuarioGORM struct {
	db *gorm.DB
}

func NewUsuarioRepositorio(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

func (r *UsuarioGORM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.CreadoEn = time.Now()
	if err := r.db.Create(&u).Error; err != nil {
		return models.Usuario{}, err
	}
	return u, nil
}

func (r *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

var _ UserRepositorio = (*UsuarioGORM)(nil)

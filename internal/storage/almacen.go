package storage

import "proyecto/internal/models"

type TecnicoRepository interface {
	ListarTecnicos() []models.Tecnico
	BuscarTecnicoPorID(id int) (models.Tecnico, bool)
	CrearTecnico(t models.Tecnico) models.Tecnico
	ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool)
	BorrarTecnico(id int) bool
}

// Interfaz para la autenticación
type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// Almacen consolida todas las interfaces
type Almacen interface {
	TecnicoRepository
}

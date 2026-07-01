package storage

import "proyecto/internal/models"

type CitaRepository interface {
	ListarCitas() []models.Cita
	BuscarCitaPorID(id int) (models.Cita, bool)
	CrearCita(c models.Cita) models.Cita
	ActualizarCita(id int, datos models.Cita) (models.Cita, bool)
	BorrarCita(id int) bool
}

type PuntoEncuentroRepository interface {
	ListarPuntosEncuentro() []models.PuntoEncuentro
	BuscarPuntoEncuentroPorID(id int) (models.PuntoEncuentro, bool)
	CrearPuntoEncuentro(p models.PuntoEncuentro) models.PuntoEncuentro
	ActualizarPuntoEncuentro(id int, datos models.PuntoEncuentro) (models.PuntoEncuentro, bool)
	BorrarPuntoEncuentro(id int) bool
}

type SoporteRepository interface {
	ListarSoportes() []models.Soporte
	BuscarSoportePorID(id int) (models.Soporte, bool)
	CrearSoporte(s models.Soporte) models.Soporte
	ActualizarSoporte(id int, datos models.Soporte) (models.Soporte, bool)
	BorrarSoporte(id int) bool
}

type UsuarioRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

type Almacen interface {
	CitaRepository
	PuntoEncuentroRepository
	SoporteRepository
}

var _ Almacen = (*Memoria)(nil)
var _ Almacen = (*AlmacenSQLite)(nil)
var _ Almacen = (*AlmacenSQLC)(nil)

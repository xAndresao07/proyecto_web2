package storage

import "proyecto/internal/models"

type SolicitanteRepository interface {
	ListarSolicitantes() []models.Solicitante
	BuscarSolicitantePorID(id int) (models.Solicitante, bool)
	CrearSolicitante(s models.Solicitante) models.Solicitante
	ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool)
	BorrarSolicitante(id int) bool
}

type DispositivoRepository interface {
	ListarDispositivos() []models.Dispositivo
	BuscarDispositivoPorID(id int) (models.Dispositivo, bool)
	CrearDispositivo(d models.Dispositivo) models.Dispositivo
	ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool)
	BorrarDispositivo(id int) bool
}

type TicketAyudaRepository interface {
	ListarTickets() []models.TicketAyuda
	BuscarTicketPorID(id int) (models.TicketAyuda, bool)
	CrearTicket(t models.TicketAyuda) models.TicketAyuda
	ActualizarTicket(id int, datos models.TicketAyuda) (models.TicketAyuda, bool)
	BorrarTicket(id int) bool
}

type UsuarioRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

type Almacen interface {
	SolicitanteRepository
	DispositivoRepository
	TicketAyudaRepository
}

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

var _ Almacen = (*Memoria)(nil)
var _ Almacen = (*AlmacenSQLite)(nil)
var _ Almacen = (*AlmacenSQLC)(nil)

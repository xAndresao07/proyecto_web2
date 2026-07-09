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
	TecnicoRepository
	CitaRepository
	PuntoEncuentroRepository
	SoporteRepository
}

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

var _ Almacen = (*Memoria)(nil)
var _ Almacen = (*AlmacenSQLite)(nil)
var _ Almacen = (*AlmacenSQLC)(nil)

type TecnicoRepository interface {
	ListarTecnicos() []models.Tecnico
	BuscarTecnicoPorID(id int) (models.Tecnico, bool)
	CrearTecnico(t models.Tecnico) models.Tecnico
	ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool)
	BorrarTecnico(id int) bool
}

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

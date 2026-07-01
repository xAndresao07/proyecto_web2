package storage

import "solicitantesYHardware/internal/models"

// Almacen define QUÉ sabe hacer un almacén de Solicitantes y Hardware, sin
// decir CÓMO. Memoria, AlmacenSQLite (GORM) y AlmacenSQLC cumplen esta misma
// interfaz — el Server no se entera de cuál de los tres recibió.

type SolicitanteRepositorio interface {
	ListarSolicitantes() []models.Solicitante
	BuscarSolicitantePorID(id int) (models.Solicitante, bool)
	CrearSolicitante(s models.Solicitante) models.Solicitante
	ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool)
	BorrarSolicitante(id int) bool
}

type DispositivoRepositorio interface {
	ListarDispositivos() []models.Dispositivo
	BuscarDispositivoPorID(id int) (models.Dispositivo, bool)
	CrearDispositivo(d models.Dispositivo) models.Dispositivo
	ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool)
	BorrarDispositivo(id int) bool
}

type TicketAyudaRepositorio interface {
	ListarTicketAyudas() []models.TicketAyuda
	BuscarTicketAyudaPorID(id int) (models.TicketAyuda, bool)
	CrearTicketAyuda(t models.TicketAyuda) models.TicketAyuda
	ActualizarTicketAyuda(id int, datos models.TicketAyuda) (models.TicketAyuda, bool)
	BorrarTicketAyuda(id int) bool
}

type UserRepositorio interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// Almacen agrupa las tres entidades del módulo. Una sola implementación
// (Memoria, AlmacenSQLite o AlmacenSQLC) cubre el módulo completo.
type Almacen interface {
	SolicitanteRepositorio
	DispositivoRepositorio
	TicketAyudaRepositorio
}

// Chequeos en tiempo de compilación.
var _ Almacen = (*Memoria)(nil)
var _ Almacen = (*AlmacenSQLite)(nil)
var _ Almacen = (*AlmacenSQLC)(nil)

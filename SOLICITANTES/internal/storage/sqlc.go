package storage

import (
	"context"
	"database/sql"

	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/storage/sqlcdb"
)

// AlmacenSQLC implementa la interfaz Almacen usando código generado por sqlc
// (SQL escrito a mano + tipado generado) sobre database/sql.
//
// Es el TERCER backend del módulo, hermano de Memoria y AlmacenSQLite.
// El Server y los handlers no se enteran de cuál reciben: todos cumplen Almacen.
//
// Diferencias con sqlc que el adaptador tiene que resolver:
//  1. Los métodos generados piden context.Context  -> lo inyectamos acá dentro.
//  2. sqlc devuelve sus propios structs (int64)     -> los MAPEAMOS a models (int).
//  3. sqlc devuelve (T, error)                       -> lo absorbemos a (T, bool).
type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

// NuevoAlmacenSQLC envuelve una conexión *sql.DB ya abierta.
func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// =========================================================
// MAPEO sqlc <-> dominio (la "capa anticorrupción")
// =========================================================

func aSolicitanteDominio(s sqlcdb.Solicitante) models.Solicitante {
	return models.Solicitante{
		ID:            int(s.ID),
		UsuarioID:     int(s.UsuarioID),
		Matricula:     s.Matricula,
		Nombre:        s.Nombre,
		Facultad:      s.Facultad,
		Semestre:      int(s.Semestre),
		NivelUrgencia: s.NivelUrgencia,
	}
}

func aDispositivoDominio(d sqlcdb.Dispositivo) models.Dispositivo {
	return models.Dispositivo{
		ID:                 int(d.ID),
		SolicitanteID:      int(d.SolicitanteID),
		Marca:              d.Marca,
		Modelo:             d.Modelo,
		TipoAlmacenamiento: d.TipoAlmacenamiento,
		RamGB:              int(d.RamGb),
		SistemaOperativo:   d.SistemaOperativo,
	}
}

func aTicketAyudaDominio(t sqlcdb.TicketAyuda) models.TicketAyuda {
	return models.TicketAyuda{
		ID:                int(t.ID),
		SolicitanteID:     int(t.SolicitanteID),
		DispositivoID:     int(t.DispositivoID),
		DescripcionFalla:  t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido,
		EstadoTicket:      t.EstadoTicket,
	}
}

// =========================================================
// SOLICITANTES
// =========================================================

func (a *AlmacenSQLC) ListarSolicitantes() []models.Solicitante {
	filas, err := a.q.ListarSolicitantes(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Solicitante, 0, len(filas))
	for _, f := range filas {
		out = append(out, aSolicitanteDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	f, err := a.q.BuscarSolicitantePorID(context.Background(), int64(id))
	if err != nil {
		// Absorbemos sql.ErrNoRows (y cualquier otro error) y conservamos la firma comma-ok.
		return models.Solicitante{}, false
	}
	return aSolicitanteDominio(f), true
}

func (a *AlmacenSQLC) CrearSolicitante(s models.Solicitante) models.Solicitante {
	f, err := a.q.CrearSolicitante(context.Background(), sqlcdb.CrearSolicitanteParams{
		UsuarioID:     int64(s.UsuarioID),
		Matricula:     s.Matricula,
		Nombre:        s.Nombre,
		Facultad:      s.Facultad,
		Semestre:      int64(s.Semestre),
		NivelUrgencia: s.NivelUrgencia,
	})
	if err != nil {
		// La interfaz no permite reportar el fallo de una creación (igual que Memoria
		// y AlmacenSQLite). Devolvemos el zero value.
		return models.Solicitante{}
	}
	return aSolicitanteDominio(f)
}

func (a *AlmacenSQLC) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	f, err := a.q.ActualizarSolicitante(context.Background(), sqlcdb.ActualizarSolicitanteParams{
		Matricula:     datos.Matricula,
		Nombre:        datos.Nombre,
		Facultad:      datos.Facultad,
		Semestre:      int64(datos.Semestre),
		NivelUrgencia: datos.NivelUrgencia,
		ID:            int64(id),
	})
	if err != nil {
		return models.Solicitante{}, false
	}
	return aSolicitanteDominio(f), true
}

func (a *AlmacenSQLC) BorrarSolicitante(id int) bool {
	filas, err := a.q.BorrarSolicitante(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// DISPOSITIVOS
// =========================================================

func (a *AlmacenSQLC) ListarDispositivos() []models.Dispositivo {
	filas, err := a.q.ListarDispositivos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Dispositivo, 0, len(filas))
	for _, f := range filas {
		out = append(out, aDispositivoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	f, err := a.q.BuscarDispositivoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Dispositivo{}, false
	}
	return aDispositivoDominio(f), true
}

func (a *AlmacenSQLC) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	f, err := a.q.CrearDispositivo(context.Background(), sqlcdb.CrearDispositivoParams{
		SolicitanteID:      int64(d.SolicitanteID),
		Marca:              d.Marca,
		Modelo:             d.Modelo,
		TipoAlmacenamiento: d.TipoAlmacenamiento,
		RamGb:              int64(d.RamGB),
		SistemaOperativo:   d.SistemaOperativo,
	})
	if err != nil {
		return models.Dispositivo{}
	}
	return aDispositivoDominio(f)
}

func (a *AlmacenSQLC) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	f, err := a.q.ActualizarDispositivo(context.Background(), sqlcdb.ActualizarDispositivoParams{
		Marca:              datos.Marca,
		Modelo:             datos.Modelo,
		TipoAlmacenamiento: datos.TipoAlmacenamiento,
		RamGb:              int64(datos.RamGB),
		SistemaOperativo:   datos.SistemaOperativo,
		ID:                 int64(id),
	})
	if err != nil {
		return models.Dispositivo{}, false
	}
	return aDispositivoDominio(f), true
}

func (a *AlmacenSQLC) BorrarDispositivo(id int) bool {
	filas, err := a.q.BorrarDispositivo(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// =========================================================
// TICKET AYUDAS
// =========================================================

func (a *AlmacenSQLC) ListarTicketAyudas() []models.TicketAyuda {
	filas, err := a.q.ListarTicketAyudas(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.TicketAyuda, 0, len(filas))
	for _, f := range filas {
		out = append(out, aTicketAyudaDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarTicketAyudaPorID(id int) (models.TicketAyuda, bool) {
	f, err := a.q.BuscarTicketAyudaPorID(context.Background(), int64(id))
	if err != nil {
		return models.TicketAyuda{}, false
	}
	return aTicketAyudaDominio(f), true
}

func (a *AlmacenSQLC) CrearTicketAyuda(t models.TicketAyuda) models.TicketAyuda {
	f, err := a.q.CrearTicketAyuda(context.Background(), sqlcdb.CrearTicketAyudaParams{
		SolicitanteID:     int64(t.SolicitanteID),
		DispositivoID:     int64(t.DispositivoID),
		DescripcionFalla:  t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido,
		EstadoTicket:      t.EstadoTicket,
	})
	if err != nil {
		return models.TicketAyuda{}
	}
	return aTicketAyudaDominio(f)
}

func (a *AlmacenSQLC) ActualizarTicketAyuda(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) {
	f, err := a.q.ActualizarTicketAyuda(context.Background(), sqlcdb.ActualizarTicketAyudaParams{
		DispositivoID:     int64(datos.DispositivoID),
		DescripcionFalla:  datos.DescripcionFalla,
		SoftwareRequerido: datos.SoftwareRequerido,
		EstadoTicket:      datos.EstadoTicket,
		ID:                int64(id),
	})
	if err != nil {
		return models.TicketAyuda{}, false
	}
	return aTicketAyudaDominio(f), true
}

func (a *AlmacenSQLC) BorrarTicketAyuda(id int) bool {
	filas, err := a.q.BorrarTicketAyuda(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// Chequeo en tiempo de compilación: AlmacenSQLC debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLC)(nil)

package storage

import (
	"context"
	"database/sql"
	"log"

	"proyecto/internal/models"
	"proyecto/internal/storage/sqlcdb"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// --- Mapeadores de sqlcdb a models ---
func aSolicitante(s sqlcdb.Solicitante) models.Solicitante {
	return models.Solicitante{
		ID:            int(s.ID),
		Nombre:        s.Nombre,
		Facultad:      s.Facultad,
		Semestre:      int(s.Semestre),
		NivelUrgencia: s.NivelUrgencia,
	}
}

func aDispositivo(d sqlcdb.Dispositivo) models.Dispositivo {
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

func aTicket(t sqlcdb.TicketAyuda) models.TicketAyuda {
	return models.TicketAyuda{
		ID:                int(t.ID),
		SolicitanteID:     int(t.SolicitanteID),
		DispositivoID:     int(t.DispositivoID),
		DescripcionFalla:  t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido,
		EstadoTicket:      t.EstadoTicket,
	}
}

// --- Solicitantes ---
func (a *AlmacenSQLC) ListarSolicitantes() []models.Solicitante {
	f, _ := a.q.ListarSolicitantes(context.Background())
	out := make([]models.Solicitante, 0, len(f))
	for _, x := range f {
		out = append(out, aSolicitante(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	f, err := a.q.BuscarSolicitantePorID(context.Background(), int64(id))
	if err != nil {
		return models.Solicitante{}, false
	}
	return aSolicitante(f), true
}

func (a *AlmacenSQLC) CrearSolicitante(s models.Solicitante) models.Solicitante {
	f, err := a.q.CrearSolicitante(context.Background(), sqlcdb.CrearSolicitanteParams{
		ID:            int64(s.ID),
		Nombre:        s.Nombre,
		Facultad:      s.Facultad,
		Semestre:      int64(s.Semestre),
		NivelUrgencia: s.NivelUrgencia,
	})
	if err != nil {
		log.Printf("Error SQLC: %v", err)
		return models.Solicitante{}
	}
	return aSolicitante(f)
}

func (a *AlmacenSQLC) ActualizarSolicitante(id int, d models.Solicitante) (models.Solicitante, bool) {
	f, err := a.q.ActualizarSolicitante(context.Background(), sqlcdb.ActualizarSolicitanteParams{
		Nombre: d.Nombre, Facultad: d.Facultad, Semestre: int64(d.Semestre), NivelUrgencia: d.NivelUrgencia, ID: int64(id),
	})
	if err != nil {
		return models.Solicitante{}, false
	}
	return aSolicitante(f), true
}

func (a *AlmacenSQLC) BorrarSolicitante(id int) bool {
	n, err := a.q.BorrarSolicitante(context.Background(), int64(id))
	return err == nil && n > 0
}

// --- Dispositivos ---
func (a *AlmacenSQLC) ListarDispositivos() []models.Dispositivo {
	f, _ := a.q.ListarDispositivos(context.Background())
	out := make([]models.Dispositivo, 0, len(f))
	for _, x := range f {
		out = append(out, aDispositivo(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	f, err := a.q.BuscarDispositivoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Dispositivo{}, false
	}
	return aDispositivo(f), true
}

func (a *AlmacenSQLC) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	f, err := a.q.CrearDispositivo(context.Background(), sqlcdb.CrearDispositivoParams{
		ID: int64(d.ID), SolicitanteID: int64(d.SolicitanteID), Marca: d.Marca, Modelo: d.Modelo,
		TipoAlmacenamiento: d.TipoAlmacenamiento, RamGb: int64(d.RamGB), SistemaOperativo: d.SistemaOperativo,
	})
	if err != nil {
		log.Printf("Error SQLC: %v", err)
		return models.Dispositivo{}
	}
	return aDispositivo(f)
}

func (a *AlmacenSQLC) ActualizarDispositivo(id int, d models.Dispositivo) (models.Dispositivo, bool) {
	f, err := a.q.ActualizarDispositivo(context.Background(), sqlcdb.ActualizarDispositivoParams{
		SolicitanteID: int64(d.SolicitanteID), Marca: d.Marca, Modelo: d.Modelo,
		TipoAlmacenamiento: d.TipoAlmacenamiento, RamGb: int64(d.RamGB), SistemaOperativo: d.SistemaOperativo, ID: int64(id),
	})
	if err != nil {
		return models.Dispositivo{}, false
	}
	return aDispositivo(f), true
}

func (a *AlmacenSQLC) BorrarDispositivo(id int) bool {
	n, err := a.q.BorrarDispositivo(context.Background(), int64(id))
	return err == nil && n > 0
}

// --- Tickets de Ayuda ---
func (a *AlmacenSQLC) ListarTickets() []models.TicketAyuda {
	f, _ := a.q.ListarTickets(context.Background())
	out := make([]models.TicketAyuda, 0, len(f))
	for _, x := range f {
		out = append(out, aTicket(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarTicketPorID(id int) (models.TicketAyuda, bool) {
	f, err := a.q.BuscarTicketPorID(context.Background(), int64(id))
	if err != nil {
		return models.TicketAyuda{}, false
	}
	return aTicket(f), true
}

func (a *AlmacenSQLC) CrearTicket(t models.TicketAyuda) models.TicketAyuda {
	f, err := a.q.CrearTicket(context.Background(), sqlcdb.CrearTicketParams{
		SolicitanteID: int64(t.SolicitanteID), DispositivoID: int64(t.DispositivoID), DescripcionFalla: t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido, EstadoTicket: t.EstadoTicket,
	})
	if err != nil {
		log.Printf("Error SQLC: %v", err)
		return models.TicketAyuda{}
	}
	return aTicket(f)
}

func (a *AlmacenSQLC) ActualizarTicket(id int, t models.TicketAyuda) (models.TicketAyuda, bool) {
	f, err := a.q.ActualizarTicket(context.Background(), sqlcdb.ActualizarTicketParams{
		SolicitanteID: int64(t.SolicitanteID), DispositivoID: int64(t.DispositivoID), DescripcionFalla: t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido, EstadoTicket: t.EstadoTicket, ID: int64(id),
	})
	if err != nil {
		return models.TicketAyuda{}, false
	}
	return aTicket(f), true
}

func (a *AlmacenSQLC) BorrarTicket(id int) bool {
	n, err := a.q.BorrarTicket(context.Background(), int64(id))
	return err == nil && n > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)

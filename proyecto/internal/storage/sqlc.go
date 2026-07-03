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
func aCita(c sqlcdb.Citum) models.Cita {
	return models.Cita{
		ID:             int(c.ID),
		SolicitanteID:  c.SolicitanteID,
		TecnicoID:      c.TecnicoID,
		Estado:         c.Estado,
		HoraAcordada:   c.HoraAcordada,
		PuntoEncuentro: c.PuntoEncuentro,
	}
}

func aPunto(p sqlcdb.PuntoEncuentro) models.PuntoEncuentro {
	return models.PuntoEncuentro{
		ID:                    int(p.ID),
		NombreLugar:           p.NombreLugar,
		FacultadPerteneciente: p.FacultadPerteneciente,
		DisponibleParaSoporte: p.DisponibleParaSoporte,
	}
}

func aSoporte(s sqlcdb.Soporte) models.Soporte {
	return models.Soporte{
		ID:              int(s.ID),
		CitaID:          int(s.CitaID),
		DispositivoID:   int(s.DispositivoID),
		Solucion:        s.Solucion,
		PiezasCambiadas: s.PiezasCambiadas,
	}
}

// --- Citas ---
func (a *AlmacenSQLC) ListarCitas() []models.Cita {
	f, _ := a.q.ListarCitas(context.Background())
	out := make([]models.Cita, 0, len(f))
	for _, x := range f {
		out = append(out, aCita(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCitaPorID(id int) (models.Cita, bool) {
	f, err := a.q.BuscarCitaPorID(context.Background(), int64(id))
	if err != nil {
		return models.Cita{}, false
	}
	return aCita(f), true
}

func (a *AlmacenSQLC) CrearCita(c models.Cita) models.Cita {
	// IMPORTANTE: Verifica el orden según tu queries.sql (schema.sql)
	f, err := a.q.CrearCita(context.Background(), sqlcdb.CrearCitaParams{
		SolicitanteID:  c.SolicitanteID,
		TecnicoID:      c.TecnicoID,
		Estado:         c.Estado,
		HoraAcordada:   c.HoraAcordada,
		PuntoEncuentro: c.PuntoEncuentro,
	})
	if err != nil {
		log.Printf("Error SQLC: %v", err) // Esto te dirá qué campo falla
		return models.Cita{}
	}
	return aCita(f)
}

func (a *AlmacenSQLC) ActualizarCita(id int, d models.Cita) (models.Cita, bool) {
	f, err := a.q.ActualizarCita(context.Background(), sqlcdb.ActualizarCitaParams{
		SolicitanteID: d.SolicitanteID, TecnicoID: d.TecnicoID, Estado: d.Estado, HoraAcordada: d.HoraAcordada, PuntoEncuentro: d.PuntoEncuentro, ID: int64(id),
	})
	if err != nil {
		return models.Cita{}, false
	}
	return aCita(f), true
}

func (a *AlmacenSQLC) BorrarCita(id int) bool {
	n, err := a.q.BorrarCita(context.Background(), int64(id))
	return err == nil && n > 0
}

// --- Puntos ---
func (a *AlmacenSQLC) ListarPuntosEncuentro() []models.PuntoEncuentro {
	f, _ := a.q.ListarPuntosEncuentro(context.Background())
	out := make([]models.PuntoEncuentro, 0, len(f))
	for _, x := range f {
		out = append(out, aPunto(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarPuntoEncuentroPorID(id int) (models.PuntoEncuentro, bool) {
	f, err := a.q.BuscarPuntoEncuentroPorID(context.Background(), int64(id))
	if err != nil {
		return models.PuntoEncuentro{}, false
	}
	return aPunto(f), true
}

func (a *AlmacenSQLC) CrearPuntoEncuentro(p models.PuntoEncuentro) models.PuntoEncuentro {
	f, _ := a.q.CrearPuntoEncuentro(context.Background(), sqlcdb.CrearPuntoEncuentroParams{NombreLugar: p.NombreLugar, FacultadPerteneciente: p.FacultadPerteneciente, DisponibleParaSoporte: p.DisponibleParaSoporte})
	return aPunto(f)
}

func (a *AlmacenSQLC) ActualizarPuntoEncuentro(id int, d models.PuntoEncuentro) (models.PuntoEncuentro, bool) {
	f, err := a.q.ActualizarPuntoEncuentro(context.Background(), sqlcdb.ActualizarPuntoEncuentroParams{NombreLugar: d.NombreLugar, FacultadPerteneciente: d.FacultadPerteneciente, DisponibleParaSoporte: d.DisponibleParaSoporte, ID: int64(id)})
	if err != nil {
		return models.PuntoEncuentro{}, false
	}
	return aPunto(f), true
}

func (a *AlmacenSQLC) BorrarPuntoEncuentro(id int) bool {
	n, err := a.q.BorrarPuntoEncuentro(context.Background(), int64(id))
	return err == nil && n > 0
}

// --- Soportes ---
func (a *AlmacenSQLC) ListarSoportes() []models.Soporte {
	f, _ := a.q.ListarSoportes(context.Background())
	out := make([]models.Soporte, 0, len(f))
	for _, x := range f {
		out = append(out, aSoporte(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSoportePorID(id int) (models.Soporte, bool) {
	f, err := a.q.BuscarSoportePorID(context.Background(), int64(id))
	if err != nil {
		return models.Soporte{}, false
	}
	return aSoporte(f), true
}

func (a *AlmacenSQLC) CrearSoporte(s models.Soporte) models.Soporte {
	f, _ := a.q.CrearSoporte(context.Background(), sqlcdb.CrearSoporteParams{CitaID: int64(s.CitaID), DispositivoID: int64(s.DispositivoID), Solucion: s.Solucion, PiezasCambiadas: s.PiezasCambiadas})
	return aSoporte(f)
}

func (a *AlmacenSQLC) ActualizarSoporte(id int, s models.Soporte) (models.Soporte, bool) {
	f, err := a.q.ActualizarSoporte(context.Background(), sqlcdb.ActualizarSoporteParams{CitaID: int64(s.CitaID), DispositivoID: int64(s.DispositivoID), Solucion: s.Solucion, PiezasCambiadas: s.PiezasCambiadas, ID: int64(id)})
	if err != nil {
		return models.Soporte{}, false
	}
	return aSoporte(f), true
}

func (a *AlmacenSQLC) BorrarSoporte(id int) bool {
	n, err := a.q.BorrarSoporte(context.Background(), int64(id))
	return err == nil && n > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)

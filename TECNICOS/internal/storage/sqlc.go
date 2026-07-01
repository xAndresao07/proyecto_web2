package storage

import (
	"context"
	"database/sql"

	"proyecto/internal/models"
	"proyecto/internal/storage/sqlcdb"
)

type AlmacenSQLC struct {
	db *sql.DB
	q  *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{db: db, q: sqlcdb.New(db)}
}

func (a *AlmacenSQLC) ListarTecnicos() []models.Tecnico {
	filas, err := a.q.ListarTecnicos(context.Background())
	if err != nil {
		return nil
	}
	var tecnicos []models.Tecnico
	for _, f := range filas {
		tecnicos = append(tecnicos, a.armarTecnicoCompleto(int(f.ID), f.Nombre, f.Reputacion))
	}
	return tecnicos
}

func (a *AlmacenSQLC) BuscarTecnicoPorID(id int) (models.Tecnico, bool) {
	f, err := a.q.BuscarTecnicoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Tecnico{}, false
	}
	return a.armarTecnicoCompleto(int(f.ID), f.Nombre, f.Reputacion), true
}

// Función auxiliar para traer servicios y horarios
func (a *AlmacenSQLC) armarTecnicoCompleto(id int, nombre string, reputacion float64) models.Tecnico {
	ctx := context.Background()
	t := models.Tecnico{ID: id, Nombre: nombre, Reputacion: reputacion}

	srvs, _ := a.q.ListarServiciosPorTecnico(ctx, int64(id))
	for _, s := range srvs {
		t.Servicios = append(t.Servicios, models.ServicioOfrecido{
			ID: int(s.ID), TecnicoID: id, NombreServicio: s.NombreServicio, NivelExperiencia: s.NivelExperiencia, TiempoEstimado: s.TiempoEstimado,
		})
	}

	hrs, _ := a.q.ListarHorariosPorTecnico(ctx, int64(id))
	for _, h := range hrs {
		t.Horarios = append(t.Horarios, models.HorarioTecnico{
			ID: int(h.ID), TecnicoID: id, DiaSemana: h.DiaSemana, HoraInicio: h.HoraInicio, HoraFin: h.HoraFin, EstadoDisponibilidad: h.EstadoDisponibilidad,
		})
	}
	return t
}

func (a *AlmacenSQLC) CrearTecnico(t models.Tecnico) models.Tecnico {
	ctx := context.Background()
	f, err := a.q.CrearTecnico(ctx, sqlcdb.CrearTecnicoParams{Nombre: t.Nombre, Reputacion: t.Reputacion})
	if err != nil {
		return models.Tecnico{}
	}
	t.ID = int(f.ID)

	for i, s := range t.Servicios {
		sf, _ := a.q.CrearServicio(ctx, sqlcdb.CrearServicioParams{TecnicoID: f.ID, NombreServicio: s.NombreServicio, NivelExperiencia: s.NivelExperiencia, TiempoEstimado: s.TiempoEstimado})
		t.Servicios[i].ID = int(sf.ID)
	}
	for i, h := range t.Horarios {
		hf, _ := a.q.CrearHorario(ctx, sqlcdb.CrearHorarioParams{TecnicoID: f.ID, DiaSemana: h.DiaSemana, HoraInicio: h.HoraInicio, HoraFin: h.HoraFin, EstadoDisponibilidad: h.EstadoDisponibilidad})
		t.Horarios[i].ID = int(hf.ID)
	}
	return t
}

func (a *AlmacenSQLC) ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool) {
	ctx := context.Background()
	_, err := a.q.ActualizarTecnico(ctx, sqlcdb.ActualizarTecnicoParams{Nombre: datos.Nombre, Reputacion: datos.Reputacion, ID: int64(id)})
	if err != nil {
		return models.Tecnico{}, false
	}

	a.q.BorrarServiciosPorTecnico(ctx, int64(id))
	a.q.BorrarHorariosPorTecnico(ctx, int64(id))

	datos.ID = id
	for i, s := range datos.Servicios {
		sf, _ := a.q.CrearServicio(ctx, sqlcdb.CrearServicioParams{TecnicoID: int64(id), NombreServicio: s.NombreServicio, NivelExperiencia: s.NivelExperiencia, TiempoEstimado: s.TiempoEstimado})
		datos.Servicios[i].ID = int(sf.ID)
	}
	for i, h := range datos.Horarios {
		hf, _ := a.q.CrearHorario(ctx, sqlcdb.CrearHorarioParams{TecnicoID: int64(id), DiaSemana: h.DiaSemana, HoraInicio: h.HoraInicio, HoraFin: h.HoraFin, EstadoDisponibilidad: h.EstadoDisponibilidad})
		datos.Horarios[i].ID = int(hf.ID)
	}
	return datos, true
}

func (a *AlmacenSQLC) BorrarTecnico(id int) bool {
	// Al borrar el técnico, SQLite borrará en cascada los servicios y horarios
	filas, err := a.q.BorrarTecnico(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

var _ TecnicoRepository = (*AlmacenSQLC)(nil)

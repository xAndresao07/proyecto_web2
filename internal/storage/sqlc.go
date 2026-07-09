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
		ID:            int64(s.ID), // El ID regresa aquí
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
		Nombre:        d.Nombre,
		Facultad:      d.Facultad,
		Semestre:      int64(d.Semestre),
		NivelUrgencia: d.NivelUrgencia,
		ID:            int64(id),
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
		ID:                 int64(d.ID), // El ID regresa aquí
		SolicitanteID:      int64(d.SolicitanteID),
		Marca:              d.Marca,
		Modelo:             d.Modelo,
		TipoAlmacenamiento: d.TipoAlmacenamiento,
		RamGb:              int64(d.RamGB),
		SistemaOperativo:   d.SistemaOperativo,
	})
	if err != nil {
		log.Printf("Error SQLC: %v", err)
		return models.Dispositivo{}
	}
	return aDispositivo(f)
}

func (a *AlmacenSQLC) ActualizarDispositivo(id int, d models.Dispositivo) (models.Dispositivo, bool) {
	f, err := a.q.ActualizarDispositivo(context.Background(), sqlcdb.ActualizarDispositivoParams{
		SolicitanteID:      int64(d.SolicitanteID),
		Marca:              d.Marca,
		Modelo:             d.Modelo,
		TipoAlmacenamiento: d.TipoAlmacenamiento,
		RamGb:              int64(d.RamGB),
		SistemaOperativo:   d.SistemaOperativo,
		ID:                 int64(id),
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
	f, _ := a.q.ListarTicketAyudas(context.Background())
	out := make([]models.TicketAyuda, 0, len(f))
	for _, x := range f {
		out = append(out, aTicket(x))
	}
	return out
}

func (a *AlmacenSQLC) BuscarTicketPorID(id int) (models.TicketAyuda, bool) {
	f, err := a.q.BuscarTicketAyudaPorID(context.Background(), int64(id))
	if err != nil {
		return models.TicketAyuda{}, false
	}
	return aTicket(f), true
}

func (a *AlmacenSQLC) CrearTicket(t models.TicketAyuda) models.TicketAyuda {
	f, err := a.q.CrearTicketAyuda(context.Background(), sqlcdb.CrearTicketAyudaParams{
		ID:                int64(t.ID), // El ID regresa aquí
		SolicitanteID:     int64(t.SolicitanteID),
		DispositivoID:     int64(t.DispositivoID),
		DescripcionFalla:  t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido,
		EstadoTicket:      t.EstadoTicket,
	})
	if err != nil {
		log.Printf("Error SQLC: %v", err)
		return models.TicketAyuda{}
	}
	return aTicket(f)
}

func (a *AlmacenSQLC) ActualizarTicket(id int, t models.TicketAyuda) (models.TicketAyuda, bool) {
	f, err := a.q.ActualizarTicketAyuda(context.Background(), sqlcdb.ActualizarTicketAyudaParams{
		SolicitanteID:     int64(t.SolicitanteID),
		DispositivoID:     int64(t.DispositivoID),
		DescripcionFalla:  t.DescripcionFalla,
		SoftwareRequerido: t.SoftwareRequerido,
		EstadoTicket:      t.EstadoTicket,
		ID:                int64(id),
	})
	if err != nil {
		return models.TicketAyuda{}, false
	}
	return aTicket(f), true
}

func (a *AlmacenSQLC) BorrarTicket(id int) bool {
	n, err := a.q.BorrarTicketAyuda(context.Background(), int64(id))
	return err == nil && n > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)


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

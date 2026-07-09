package storage

import (
	"proyecto/internal/models"
	"sync"
)

type Memoria struct {
	solicitantes  []models.Solicitante
	dispositivos  []models.Dispositivo
	tickets       []models.TicketAyuda
	tecnicos      []models.Tecnico
	citas         []models.Cita
	puntosEncuentro []models.PuntoEncuentro
	soportes      []models.Soporte
	nextSolicitanteID int
	nextDispositivoID int
	nextTicketID  int
	nextTecnicoID int
	nextCitaID    int
	nextPuntoID   int
	nextSoporteID int
	mu            sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		solicitantes:  []models.Solicitante{},
		dispositivos:  []models.Dispositivo{},
		tickets:       []models.TicketAyuda{},
		tecnicos:      []models.Tecnico{},
		citas:         []models.Cita{},
		puntosEncuentro: []models.PuntoEncuentro{},
		soportes:      []models.Soporte{},
		nextSolicitanteID: 1,
		nextDispositivoID: 1,
		nextTicketID:  1,
		nextTecnicoID: 1,
		nextCitaID:    1,
		nextPuntoID:   1,
		nextSoporteID: 1,
	}
}

// ==== SOLICITANTES ====
func (m *Memoria) SeedSolicitantes() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.solicitantes = []models.Solicitante{
		{ID: 1, Nombre: "Jandry Ejemplo", Facultad: "FACCI", Semestre: 5, NivelUrgencia: "alta"},
	}
	m.nextSolicitanteID = 2
}
func (m *Memoria) ListarSolicitantes() []models.Solicitante {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Solicitante, len(m.solicitantes))
	copy(copia, m.solicitantes)
	return copia
}
func (m *Memoria) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.solicitantes {
		if s.ID == id {
			return s, true
		}
	}
	return models.Solicitante{}, false
}
func (m *Memoria) CrearSolicitante(s models.Solicitante) models.Solicitante {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextSolicitanteID
	m.nextSolicitanteID++
	m.solicitantes = append(m.solicitantes, s)
	return s
}
func (m *Memoria) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitantes {
		if s.ID == id {
			datos.ID = id
			m.solicitantes[i] = datos
			return datos, true
		}
	}
	return models.Solicitante{}, false
}
func (m *Memoria) BorrarSolicitante(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitantes {
		if s.ID == id {
			m.solicitantes = append(m.solicitantes[:i], m.solicitantes[i+1:]...)
			return true
		}
	}
	return false
}

// ==== DISPOSITIVOS ====
func (m *Memoria) SeedDispositivos() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.dispositivos = []models.Dispositivo{
		{ID: 1, SolicitanteID: 1, Marca: "HP", Modelo: "Pavilion 14", TipoAlmacenamiento: "SSD", RamGB: 8, SistemaOperativo: "Windows 11"},
	}
	m.nextDispositivoID = 2
}
func (m *Memoria) ListarDispositivos() []models.Dispositivo {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Dispositivo, len(m.dispositivos))
	copy(copia, m.dispositivos)
	return copia
}
func (m *Memoria) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, d := range m.dispositivos {
		if d.ID == id {
			return d, true
		}
	}
	return models.Dispositivo{}, false
}
func (m *Memoria) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	m.mu.Lock()
	defer m.mu.Unlock()
	d.ID = m.nextDispositivoID
	m.nextDispositivoID++
	m.dispositivos = append(m.dispositivos, d)
	return d
}
func (m *Memoria) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, d := range m.dispositivos {
		if d.ID == id {
			datos.ID = id
			m.dispositivos[i] = datos
			return datos, true
		}
	}
	return models.Dispositivo{}, false
}
func (m *Memoria) BorrarDispositivo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, d := range m.dispositivos {
		if d.ID == id {
			m.dispositivos = append(m.dispositivos[:i], m.dispositivos[i+1:]...)
			return true
		}
	}
	return false
}

// ==== TICKETS DE AYUDA ====
func (m *Memoria) SeedTickets() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tickets = []models.TicketAyuda{
		{ID: 1, SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "No enciende", SoftwareRequerido: "Ninguno", EstadoTicket: "abierto"},
	}
	m.nextTicketID = 2
}
func (m *Memoria) ListarTickets() []models.TicketAyuda {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.TicketAyuda, len(m.tickets))
	copy(copia, m.tickets)
	return copia
}
func (m *Memoria) BuscarTicketPorID(id int) (models.TicketAyuda, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, t := range m.tickets {
		if t.ID == id {
			return t, true
		}
	}
	return models.TicketAyuda{}, false
}
func (m *Memoria) CrearTicket(t models.TicketAyuda) models.TicketAyuda {
	m.mu.Lock()
	defer m.mu.Unlock()
	t.ID = m.nextTicketID
	m.nextTicketID++
	m.tickets = append(m.tickets, t)
	return t
}
func (m *Memoria) ActualizarTicket(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.tickets {
		if t.ID == id {
			datos.ID = id
			m.tickets[i] = datos
			return datos, true
		}
	}
	return models.TicketAyuda{}, false
}
func (m *Memoria) BorrarTicket(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.tickets {
		if t.ID == id {
			m.tickets = append(m.tickets[:i], m.tickets[i+1:]...)
			return true
		}
	}
	return false
}


func (m *Memoria) ListarTecnicos() []models.Tecnico {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Tecnico, len(m.tecnicos))
	copy(copia, m.tecnicos)
	return copia
}

func (m *Memoria) BuscarTecnicoPorID(id int) (models.Tecnico, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range m.tecnicos {
		if t.ID == id {
			return t, true
		}
	}
	return models.Tecnico{}, false
}

func (m *Memoria) CrearTecnico(tecnico models.Tecnico) models.Tecnico {
	m.mu.Lock()
	defer m.mu.Unlock()

	tecnico.ID = m.nextTecnicoID
	m.nextTecnicoID++
	m.tecnicos = append(m.tecnicos, tecnico)
	return tecnico
}

// ActualizarTecnico modifica parcialmente un técnico existente.
func (m *Memoria) ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, t := range m.tecnicos {
		if t.ID == id {
			datos.ID = id // Aseguramos que el ID original no se modifique

			// Update Parcial con las nuevas estructuras
			if datos.Nombre == "" {
				datos.Nombre = t.Nombre
			}
			if datos.Reputacion == 0 {
				datos.Reputacion = t.Reputacion
			}

			// Si el cliente no envía nuevos servicios, conservamos el catálogo anterior
			if len(datos.Servicios) == 0 {
				datos.Servicios = t.Servicios
			}
			// Si el cliente no envía nuevos horarios, conservamos los anteriores
			if len(datos.Horarios) == 0 {
				datos.Horarios = t.Horarios
			}

			m.tecnicos[i] = datos
			return datos, true
		}
	}
	return models.Tecnico{}, false
}

func (m *Memoria) BorrarTecnico(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, t := range m.tecnicos {
		if t.ID == id {
			m.tecnicos = append(m.tecnicos[:i], m.tecnicos[i+1:]...)
			return true
		}
	}
	return false
}


func (m *Memoria) SeedCitas() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.citas = []models.Cita{
		{ID: 1, SolicitanteID: "est_102", TecnicoID: "tec_05", Estado: "completada", HoraAcordada: "09:00", PuntoEncuentro: "Lab CISCO"},
	}
	m.nextCitaID = 2
}
func (m *Memoria) ListarCitas() []models.Cita {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Cita, len(m.citas))
	copy(copia, m.citas)
	return copia
}
func (m *Memoria) BuscarCitaPorID(id int) (models.Cita, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, c := range m.citas {
		if c.ID == id {
			return c, true
		}
	}
	return models.Cita{}, false
}
func (m *Memoria) CrearCita(c models.Cita) models.Cita {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.nextCitaID
	m.nextCitaID++
	m.citas = append(m.citas, c)
	return c
}
func (m *Memoria) ActualizarCita(id int, datos models.Cita) (models.Cita, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.citas {
		if c.ID == id {
			datos.ID = id
			m.citas[i] = datos
			return datos, true
		}
	}
	return models.Cita{}, false
}
func (m *Memoria) BorrarCita(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.citas {
		if c.ID == id {
			m.citas = append(m.citas[:i], m.citas[i+1:]...)
			return true
		}
	}
	return false
}

// ==== PUNTOS ====
func (m *Memoria) SeedPuntosEncuentro() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.puntosEncuentro = []models.PuntoEncuentro{
		{ID: 1, NombreLugar: "Lab CISCO", FacultadPerteneciente: "FACCI", DisponibleParaSoporte: true},
	}
	m.nextPuntoID = 2
}
func (m *Memoria) ListarPuntosEncuentro() []models.PuntoEncuentro {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.PuntoEncuentro, len(m.puntosEncuentro))
	copy(copia, m.puntosEncuentro)
	return copia
}
func (m *Memoria) BuscarPuntoEncuentroPorID(id int) (models.PuntoEncuentro, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.puntosEncuentro {
		if p.ID == id {
			return p, true
		}
	}
	return models.PuntoEncuentro{}, false
}
func (m *Memoria) CrearPuntoEncuentro(p models.PuntoEncuentro) models.PuntoEncuentro {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextPuntoID
	m.nextPuntoID++
	m.puntosEncuentro = append(m.puntosEncuentro, p)
	return p
}
func (m *Memoria) ActualizarPuntoEncuentro(id int, d models.PuntoEncuentro) (models.PuntoEncuentro, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntosEncuentro {
		if p.ID == id {
			d.ID = id
			m.puntosEncuentro[i] = d
			return d, true
		}
	}
	return models.PuntoEncuentro{}, false
}
func (m *Memoria) BorrarPuntoEncuentro(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntosEncuentro {
		if p.ID == id {
			m.puntosEncuentro = append(m.puntosEncuentro[:i], m.puntosEncuentro[i+1:]...)
			return true
		}
	}
	return false
}

// ==== SOPORTES ====
func (m *Memoria) SeedSoportes() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.soportes = []models.Soporte{{ID: 1, CitaID: 1, DispositivoID: 501, Solucion: "Limpieza", PiezasCambiadas: "Ninguna"}}
	m.nextSoporteID = 2
}
func (m *Memoria) ListarSoportes() []models.Soporte {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Soporte, len(m.soportes))
	copy(copia, m.soportes)
	return copia
}
func (m *Memoria) BuscarSoportePorID(id int) (models.Soporte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.soportes {
		if s.ID == id {
			return s, true
		}
	}
	return models.Soporte{}, false
}
func (m *Memoria) CrearSoporte(s models.Soporte) models.Soporte {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextSoporteID
	m.nextSoporteID++
	m.soportes = append(m.soportes, s)
	return s
}
func (m *Memoria) ActualizarSoporte(id int, d models.Soporte) (models.Soporte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.soportes {
		if s.ID == id {
			d.ID = id
			m.soportes[i] = d
			return d, true
		}
	}
	return models.Soporte{}, false
}
func (m *Memoria) BorrarSoporte(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.soportes {
		if s.ID == id {
			m.soportes = append(m.soportes[:i], m.soportes[i+1:]...)
			return true
		}
	}
	return false
}

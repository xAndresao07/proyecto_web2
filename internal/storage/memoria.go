package storage

import (
	"proyecto/internal/models"
	"sync"
)

type Memoria struct {
	solicitantes []models.Solicitante
	dispositivos []models.Dispositivo
	tickets      []models.TicketAyuda
	nextTicketID int
	mu           sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		solicitantes: []models.Solicitante{},
		dispositivos: []models.Dispositivo{},
		tickets:      []models.TicketAyuda{},
		nextTicketID: 1,
	}
}

// ==== SOLICITANTES ====
func (m *Memoria) SeedSolicitantes() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.solicitantes = []models.Solicitante{
		{ID: 1, Nombre: "Jandry Ejemplo", Facultad: "FACCI", Semestre: 5, NivelUrgencia: "alta"},
	}
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

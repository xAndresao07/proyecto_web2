// Package storage gestiona el almacenamiento de Solicitantes y Hardware.
//
// El tipo Memoria mantiene en un solo lugar todos los datos del dominio:
// Solicitantes, Dispositivos y TicketAyudas.
package storage

import (
	"sync"

	"solicitantesYHardware/internal/models"
)

// Memoria es un almacén unificado del módulo Solicitantes y Hardware.
type Memoria struct {
	solicitantes     []models.Solicitante
	nextSolicitanteID int

	dispositivos     []models.Dispositivo
	nextDispositivoID int

	ticketAyudas     []models.TicketAyuda
	nextTicketAyudaID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		solicitantes:       []models.Solicitante{},
		nextSolicitanteID:  1,
		dispositivos:       []models.Dispositivo{},
		nextDispositivoID:  1,
		ticketAyudas:       []models.TicketAyuda{},
		nextTicketAyudaID:  1,
	}
}

// =========================================================
// SOLICITANTES
// =========================================================

// SeedSolicitantes carga solicitantes iniciales en memoria.
func (m *Memoria) SeedSolicitantes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.solicitantes = []models.Solicitante{
		{ID: 1, UsuarioID: 1, Matricula: "ULEAM-0001", Nombre: "Carlos Alberto Moreira", Facultad: "TI", Semestre: 4, NivelUrgencia: "alto"},
		{ID: 2, UsuarioID: 2, Matricula: "ULEAM-0002", Nombre: "Ricardo Villavicencio", Facultad: "Software", Semestre: 3, NivelUrgencia: "normal"},
	}
	m.nextSolicitanteID = 3
}

// ListarSolicitantes devuelve todos los solicitantes en memoria.
func (m *Memoria) ListarSolicitantes() []models.Solicitante {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Solicitante, len(m.solicitantes))
	copy(copia, m.solicitantes)
	return copia
}

// BuscarSolicitantePorID devuelve el solicitante con el ID dado (patrón comma-ok).
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

// CrearSolicitante agrega un solicitante nuevo y devuelve el registro con ID asignado.
func (m *Memoria) CrearSolicitante(s models.Solicitante) models.Solicitante {
	m.mu.Lock()
	defer m.mu.Unlock()

	s.ID = m.nextSolicitanteID
	m.nextSolicitanteID++
	m.solicitantes = append(m.solicitantes, s)
	return s
}

// ActualizarSolicitante reemplaza el solicitante con el ID dado.
func (m *Memoria) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, s := range m.solicitantes {
		if s.ID == id {
			datos.ID = id
			datos.UsuarioID = s.UsuarioID
			m.solicitantes[i] = datos
			return datos, true
		}
	}
	return models.Solicitante{}, false
}

// BorrarSolicitante elimina el solicitante con el ID dado.
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

// =========================================================
// DISPOSITIVOS
// =========================================================

// SeedDispositivos carga dispositivos iniciales que coinciden con los solicitantes de prueba.
func (m *Memoria) SeedDispositivos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.dispositivos = []models.Dispositivo{
		{ID: 1, SolicitanteID: 1, Marca: "HP", Modelo: "Pavilion 14", TipoAlmacenamiento: "HDD", RamGB: 4, SistemaOperativo: "Windows 10"},
		{ID: 2, SolicitanteID: 2, Marca: "Lenovo", Modelo: "IdeaPad 3", TipoAlmacenamiento: "SSD", RamGB: 8, SistemaOperativo: "Windows 11"},
	}
	m.nextDispositivoID = 3
}

// ListarDispositivos devuelve todos los dispositivos en memoria.
func (m *Memoria) ListarDispositivos() []models.Dispositivo {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Dispositivo, len(m.dispositivos))
	copy(copia, m.dispositivos)
	return copia
}

// BuscarDispositivoPorID devuelve el dispositivo con el ID dado (patrón comma-ok).
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

// CrearDispositivo agrega un dispositivo nuevo y devuelve el registro con ID asignado.
func (m *Memoria) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	m.mu.Lock()
	defer m.mu.Unlock()

	d.ID = m.nextDispositivoID
	m.nextDispositivoID++
	m.dispositivos = append(m.dispositivos, d)
	return d
}

// ActualizarDispositivo reemplaza el dispositivo con el ID dado.
func (m *Memoria) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, d := range m.dispositivos {
		if d.ID == id {
			datos.ID = id
			datos.SolicitanteID = d.SolicitanteID
			m.dispositivos[i] = datos
			return datos, true
		}
	}
	return models.Dispositivo{}, false
}

// BorrarDispositivo elimina el dispositivo con el ID dado.
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

// =========================================================
// TICKET AYUDAS
// =========================================================

// SeedTicketAyudas carga tickets iniciales que coinciden con los dispositivos de prueba.
func (m *Memoria) SeedTicketAyudas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.ticketAyudas = []models.TicketAyuda{
		{ID: 1, SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "PC se congela al abrir VS Code", SoftwareRequerido: "ninguno", EstadoTicket: "abierto"},
		{ID: 2, SolicitanteID: 2, DispositivoID: 2, DescripcionFalla: "Cambio de RAM solicitado", SoftwareRequerido: "ninguno", EstadoTicket: "abierto"},
	}
	m.nextTicketAyudaID = 3
}

// ListarTicketAyudas devuelve todos los tickets en memoria.
func (m *Memoria) ListarTicketAyudas() []models.TicketAyuda {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.TicketAyuda, len(m.ticketAyudas))
	copy(copia, m.ticketAyudas)
	return copia
}

// BuscarTicketAyudaPorID devuelve el ticket con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarTicketAyudaPorID(id int) (models.TicketAyuda, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range m.ticketAyudas {
		if t.ID == id {
			return t, true
		}
	}
	return models.TicketAyuda{}, false
}

// CrearTicketAyuda agrega un ticket nuevo y devuelve el registro con ID asignado.
// El estado inicial siempre se fuerza a "abierto" (la validación vive en el service).
func (m *Memoria) CrearTicketAyuda(t models.TicketAyuda) models.TicketAyuda {
	m.mu.Lock()
	defer m.mu.Unlock()

	t.ID = m.nextTicketAyudaID
	m.nextTicketAyudaID++
	m.ticketAyudas = append(m.ticketAyudas, t)
	return t
}

// ActualizarTicketAyuda reemplaza el ticket con el ID dado.
// SolicitanteID nunca cambia tras la creación (se conserva del registro original).
func (m *Memoria) ActualizarTicketAyuda(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, t := range m.ticketAyudas {
		if t.ID == id {
			datos.ID = id
			datos.SolicitanteID = t.SolicitanteID
			m.ticketAyudas[i] = datos
			return datos, true
		}
	}
	return models.TicketAyuda{}, false
}

// BorrarTicketAyuda elimina el ticket con el ID dado.
func (m *Memoria) BorrarTicketAyuda(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, t := range m.ticketAyudas {
		if t.ID == id {
			m.ticketAyudas = append(m.ticketAyudas[:i], m.ticketAyudas[i+1:]...)
			return true
		}
	}
	return false
}

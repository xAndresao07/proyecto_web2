package storage

import (
	"proyecto/internal/models"
	"sync"
)

type Memoria struct {
	citas           []models.Cita
	nextCitaID      int
	puntosEncuentro []models.PuntoEncuentro
	nextPuntoID     int
	soportes        []models.Soporte
	nextSoporteID   int
	mu              sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		citas:           []models.Cita{},
		nextCitaID:      1,
		puntosEncuentro: []models.PuntoEncuentro{},
		nextPuntoID:     1,
		soportes:        []models.Soporte{},
		nextSoporteID:   1,
	}
}

// ==== CITAS ====
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

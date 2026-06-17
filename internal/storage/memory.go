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

func NewMemoria() *Memoria {
	return &Memoria{
		citas:           []models.Cita{},
		nextCitaID:      1,
		puntosEncuentro: []models.PuntoEncuentro{},
		nextPuntoID:     1,
		soportes:        []models.Soporte{},
		nextSoporteID:   1,
	}
}

func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. Inyectando Puntos de Encuentro de la universidad
	m.puntosEncuentro = []models.PuntoEncuentro{
		{ID: 1, NombreLugar: "Laboratorio de Redes (CISCO)", FacultadPerteneciente: "FACCI", DisponibleParaSoporte: true},
		{ID: 2, NombreLugar: "Biblioteca Central ULEAM", FacultadPerteneciente: "General", DisponibleParaSoporte: true},
		{ID: 3, NombreLugar: "Sala de Lectura FCCEA", FacultadPerteneciente: "Ciencias Administrativas", DisponibleParaSoporte: true},
		{ID: 4, NombreLugar: "Auditorio Alfonso Aguilar", FacultadPerteneciente: "General", DisponibleParaSoporte: false}, // Reservado, no disponible
		{ID: 5, NombreLugar: "Cafetería Ciencias Básicas", FacultadPerteneciente: "Ciencias Básicas", DisponibleParaSoporte: true},
	}
	m.nextPuntoID = 6

	// 2. Inyectando Citas
	m.citas = []models.Cita{
		// Citas Completadas
		{ID: 1, SolicitanteID: "estudiante_102", TecnicoID: "tecnico_005", Estado: "completada", HoraAcordada: "09:00", PuntoEncuentro: "Laboratorio de Redes (CISCO)"},
		{ID: 2, SolicitanteID: "estudiante_902", TecnicoID: "tecnico_089", Estado: "completada", HoraAcordada: "10:30", PuntoEncuentro: "Sala de Lectura FCCEA"},
		// Citas Pendientes
		{ID: 3, SolicitanteID: "estudiante_304", TecnicoID: "tecnico_089", Estado: "pendiente", HoraAcordada: "14:30", PuntoEncuentro: "Biblioteca Central ULEAM"},
		{ID: 4, SolicitanteID: "estudiante_501", TecnicoID: "tecnico_012", Estado: "pendiente", HoraAcordada: "16:00", PuntoEncuentro: "Cafetería Ciencias Básicas"},
		{ID: 5, SolicitanteID: "estudiante_808", TecnicoID: "tecnico_005", Estado: "pendiente", HoraAcordada: "17:00", PuntoEncuentro: "Laboratorio de Redes (CISCO)"},
	}
	m.nextCitaID = 6

	// 3. Inyectando Soportes
	m.soportes = []models.Soporte{

		{ID: 1, CitaID: 1, DispositivoID: 5501, Solucion: "Limpieza de virus, optimización de inicio y actualización de drivers de red", PiezasCambiadas: "Ninguna"},

		{ID: 2, CitaID: 2, DispositivoID: 5502, Solucion: "Formateo e instalación de Windows 11 para correr software de contabilidad", PiezasCambiadas: "SSD Kingston M.2 500GB"},
	}
	m.nextSoporteID = 3
}

// =========================================================
// CRUD: CITAS (Usando tus campos de texto)
// =========================================================

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
	m.citas = append(m.citas, c)
	m.nextCitaID++
	return c
}

func (m *Memoria) ActualizarCita(id int, nueva models.Cita) (models.Cita, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.citas {
		if c.ID == id {
			nueva.ID = c.ID
			// Conservamos datos si vienen vacíos en el PUT
			if nueva.SolicitanteID == "" {
				nueva.SolicitanteID = c.SolicitanteID
			}
			if nueva.TecnicoID == "" {
				nueva.TecnicoID = c.TecnicoID
			}
			if nueva.Estado == "" {
				nueva.Estado = c.Estado
			}
			if nueva.HoraAcordada == "" {
				nueva.HoraAcordada = c.HoraAcordada
			}
			if nueva.PuntoEncuentro == "" {
				nueva.PuntoEncuentro = c.PuntoEncuentro
			}

			m.citas[i] = nueva
			return nueva, true
		}
	}
	return models.Cita{}, false
}

func (m *Memoria) EliminarCita(id int) bool {
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

// =========================================================
// CRUD: PUNTOS ENCUENTRO
// =========================================================

func (m *Memoria) ListarPuntos() []models.PuntoEncuentro {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.PuntoEncuentro, len(m.puntosEncuentro))
	copy(copia, m.puntosEncuentro)
	return copia
}

func (m *Memoria) BuscarPuntoPorID(id int) (models.PuntoEncuentro, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.puntosEncuentro {
		if p.ID == id {
			return p, true
		}
	}
	return models.PuntoEncuentro{}, false
}

func (m *Memoria) CrearPunto(p models.PuntoEncuentro) models.PuntoEncuentro {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextPuntoID
	m.puntosEncuentro = append(m.puntosEncuentro, p)
	m.nextPuntoID++
	return p
}

func (m *Memoria) ActualizarPunto(id int, nuevo models.PuntoEncuentro) (models.PuntoEncuentro, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.puntosEncuentro {
		if p.ID == id {
			nuevo.ID = p.ID
			if nuevo.NombreLugar == "" {
				nuevo.NombreLugar = p.NombreLugar
			}
			if nuevo.FacultadPerteneciente == "" {
				nuevo.FacultadPerteneciente = p.FacultadPerteneciente
			}
			// El booleano se sobrescribe directamente

			m.puntosEncuentro[i] = nuevo
			return nuevo, true
		}
	}
	return models.PuntoEncuentro{}, false
}

func (m *Memoria) EliminarPunto(id int) bool {
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

// =========================================================
// CRUD: SOPORTES
// =========================================================

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
	m.soportes = append(m.soportes, s)
	m.nextSoporteID++
	return s
}

func (m *Memoria) ActualizarSoporte(id int, nuevo models.Soporte) (models.Soporte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.soportes {
		if s.ID == id {
			nuevo.ID = s.ID
			if nuevo.CitaID == 0 {
				nuevo.CitaID = s.CitaID
			}
			if nuevo.DispositivoID == 0 {
				nuevo.DispositivoID = s.DispositivoID
			}
			if nuevo.Solucion == "" {
				nuevo.Solucion = s.Solucion
			}
			if nuevo.PiezasCambiadas == "" {
				nuevo.PiezasCambiadas = s.PiezasCambiadas
			}

			m.soportes[i] = nuevo
			return nuevo, true
		}
	}
	return models.Soporte{}, false
}

func (m *Memoria) EliminarSoporte(id int) bool {
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

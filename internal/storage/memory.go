package storage

import (
	"proyecto/internal/models"
	"sync"
)

// Memoria mantiene en un solo lugar todos los datos del dominio de Citas.
type Memoria struct {
	citas  []models.Cita
	nextID int
	mu     sync.Mutex
}

// NewMemoria crea un almacén vacío y listo para usar.
func NewMemoria() *Memoria {
	return &Memoria{
		citas:  []models.Cita{},
		nextID: 1,
	}
}

// Seed carga citas iniciales de prueba en memoria para facilitar los tests.
func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.citas = []models.Cita{
		{ID: 1, SolicitanteID: "estudiante_001", TecnicoID: "tecnico_089", Estado: "pendiente", HoraAcordada: "14:30", PuntoEncuentro: "Biblioteca Bloque B"},
		{ID: 2, SolicitanteID: "estudiante_002", TecnicoID: "tecnico_045", Estado: "completada", HoraAcordada: "10:00", PuntoEncuentro: "Laboratorio 1"},
	}
	m.nextID = 3
}

// ListarCitas devuelve todas las citas guardadas en memoria.
func (m *Memoria) ListarCitas() []models.Cita {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Cita, len(m.citas))
	copy(copia, m.citas)
	return copia
}

// BuscarCitaPorID devuelve la cita correspondiente al ID dado .
func (m *Memoria) BuscarCitaPorID(id int) (models.Cita, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, i := range m.citas {
		if i.ID == id {
			return i, true
		}
	}
	return models.Cita{}, false
}

// CrearCita agrega una nueva cita asignándole un ID incremental automático.
func (m *Memoria) CrearCita(cita models.Cita) models.Cita {
	m.mu.Lock()
	defer m.mu.Unlock()

	cita.ID = m.nextID
	m.citas = append(m.citas, cita)
	m.nextID++
	return cita
}

// ActualizarCita reemplaza los datos de la cita con el ID especificado.
func (m *Memoria) ActualizarCita(id int, cita models.Cita) (models.Cita, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.citas {
		if n.ID == id {
			cita.ID = n.ID

			// Protección para evitar que los campos opcionales del PUT limpien los datos existentes
			if cita.SolicitanteID == "" {
				cita.SolicitanteID = n.SolicitanteID
			}
			if cita.TecnicoID == "" {
				cita.TecnicoID = n.TecnicoID
			}
			if cita.PuntoEncuentro == "" {
				cita.PuntoEncuentro = n.PuntoEncuentro
			}
			if cita.Estado == "" {
				cita.Estado = n.Estado
			}
			if cita.HoraAcordada == "" {
				cita.HoraAcordada = n.HoraAcordada
			}

			m.citas[i] = cita
			return cita, true
		}
	}
	return models.Cita{}, false
}

// EliminarCita remueve de la lista la cita con el ID indicado.
func (m *Memoria) EliminarCita(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.citas {
		if n.ID == id {
			m.citas = append(m.citas[:i], m.citas[i+1:]...)
			return true
		}
	}
	return false
}

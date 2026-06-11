package storage

import (
	"proyecto/internal/models"
	"sync"
)

// Memoria mantiene en un solo lugar todos los datos del dominio de Intervenciones.
type Memoria struct {
	intervenciones []models.Intervencion
	nextID         int
	mu             sync.Mutex
}

// NewMemoria crea un almacén vacío y listo para usar.
func NewMemoria() *Memoria {
	return &Memoria{
		intervenciones: []models.Intervencion{},
		nextID:         1,
	}
}

// Seed carga intervenciones iniciales de prueba en memoria para facilitar los tests.
func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.intervenciones = []models.Intervencion{
		{ID: 1, SolicitanteID: "estudiante_001", TecnicoID: "tecnico_089", Estado: "pendiente", HoraAcordada: "14:30", PuntoEncuentro: "Biblioteca Bloque B"},
		{ID: 2, SolicitanteID: "estudiante_002", TecnicoID: "tecnico_045", Estado: "completada", HoraAcordada: "10:00", PuntoEncuentro: "Laboratorio 1"},
	}
	m.nextID = 3
}

// ListarIntervenciones devuelve todas las intervenciones guardadas en memoria.
func (m *Memoria) ListarIntervenciones() []models.Intervencion {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Intervencion, len(m.intervenciones))
	copy(copia, m.intervenciones)
	return copia
}

// BuscarIntervencionPorID devuelve la intervención correspondiente al ID dado .
func (m *Memoria) BuscarIntervencionPorID(id int) (models.Intervencion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, i := range m.intervenciones {
		if i.ID == id {
			return i, true
		}
	}
	return models.Intervencion{}, false
}

// CrearIntervencion agrega una nueva intervención asignándole un ID incremental automático.
func (m *Memoria) CrearIntervencion(intervencion models.Intervencion) models.Intervencion {
	m.mu.Lock()
	defer m.mu.Unlock()

	intervencion.ID = m.nextID
	m.intervenciones = append(m.intervenciones, intervencion)
	m.nextID++
	return intervencion
}

// ActualizarIntervencion reemplaza los datos de la intervención con el ID especificado.
func (m *Memoria) ActualizarIntervencion(id int, intervencion models.Intervencion) (models.Intervencion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.intervenciones {
		if n.ID == id {
			intervencion.ID = n.ID

			// Protección para evitar que los campos opcionales del PUT limpien los datos existentes
			if intervencion.SolicitanteID == "" {
				intervencion.SolicitanteID = n.SolicitanteID
			}
			if intervencion.TecnicoID == "" {
				intervencion.TecnicoID = n.TecnicoID
			}
			if intervencion.PuntoEncuentro == "" {
				intervencion.PuntoEncuentro = n.PuntoEncuentro
			}
			if intervencion.Estado == "" {
				intervencion.Estado = n.Estado
			}
			if intervencion.HoraAcordada == "" {
				intervencion.HoraAcordada = n.HoraAcordada
			}

			m.intervenciones[i] = intervencion
			return intervencion, true
		}
	}
	return models.Intervencion{}, false
}

// EliminarIntervencion remueve de la lista la intervención con el ID indicado.
func (m *Memoria) EliminarIntervencion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.intervenciones {
		if n.ID == id {
			m.intervenciones = append(m.intervenciones[:i], m.intervenciones[i+1:]...)
			return true
		}
	}
	return false
}

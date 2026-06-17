package storage

import (
	"proyecto/internal/models"
	"strconv" // Necesario para convertir el ID a string
	"sync"
)

type Memoria struct {
	tecnicos []models.Tecnico
	nextID   int
	mu       sync.Mutex
}

// NewMemoria crea un almacén vacío y listo para usar.
func NewMemoria() *Memoria {
	return &Memoria{
		tecnicos: []models.Tecnico{},
		nextID:   1,
	}
}

func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.tecnicos = []models.Tecnico{
		{ID: "1", Nombre: "Juan Pérez", Reputacion: 4.5},
		{ID: "2", Nombre: "María García", Reputacion: 4.8},
		{ID: "3", Nombre: "Carlos López", Reputacion: 4.2},
	}
	m.nextID = 4
}

func (m *Memoria) ListarTecnicos() []models.Tecnico {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Tecnico, len(m.tecnicos))
	copy(copia, m.tecnicos)
	return copia
}

func (m *Memoria) BuscarTecnicoPorID(id string) (models.Tecnico, bool) {
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

	tecnico.ID = strconv.Itoa(m.nextID)
	m.nextID++
	m.tecnicos = append(m.tecnicos, tecnico)
	return tecnico
}

// ActualizarTecnico modifica parcialmente un técnico existente.
func (m *Memoria) ActualizarTecnico(id string, datos models.Tecnico) (models.Tecnico, bool) {
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

func (m *Memoria) EliminarTecnico(id string) bool {
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

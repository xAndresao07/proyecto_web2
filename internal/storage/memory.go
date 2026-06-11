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

// ... (Aquí va el método Seed de Mario y sus métodos CRUD de intervenciones) ...
func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.tecnicos = []models.Tecnico{
		{ID: "1", Nombre: "Juan Pérez", Habilidades: []string{"Python", "Go"}, HorarioLibre: "Lunes 10:00-12:00", Reputacion: 4.5},
		{ID: "2", Nombre: "María García", Habilidades: []string{"JavaScript", "React"}, HorarioLibre: "Martes 14:00-16:00", Reputacion: 4.8},
		{ID: "3", Nombre: "Carlos López", Habilidades: []string{"Java", "Spring"}, HorarioLibre: "Miércoles 09:00-11:00", Reputacion: 4.2},
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

func (m *Memoria) ActualizarTecnico(id string, datos models.Tecnico) (models.Tecnico, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, t := range m.tecnicos {
		if t.ID == id {
			datos.ID = id // Protegemos el ID original

			// Actualización parcial
			if datos.Nombre == "" {
				datos.Nombre = t.Nombre
			}
			if len(datos.Habilidades) == 0 {
				datos.Habilidades = t.Habilidades
			}
			if datos.HorarioLibre == "" {
				datos.HorarioLibre = t.HorarioLibre
			}
			if datos.Reputacion == 0 {
				datos.Reputacion = t.Reputacion
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

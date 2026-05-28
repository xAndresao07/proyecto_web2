package storage

import (
	"proyecto/internal/models"
	"sync"
)

var (
	// Mu nos ayuda a evitar errores si dos peticiones intentan guardar al mismo tiempo
	Mu sync.Mutex

	// Intervenciones es nuestro "slice" (lista) que simula la tabla de la base de datos
	Intervenciones []models.Intervencion

	// SiguienteID es un contador para asignar un ID único a cada nueva intervención
	SiguienteID = 1
)

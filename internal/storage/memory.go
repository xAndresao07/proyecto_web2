package storage

import (
	"proyecto/internal/models"
	"sync"
)

var (
	// Mu nos ayuda a evitar errores si dos peticiones intentan guardar al mismo tiempo
	Mu sync.Mutex

	// Tecnicos es nuestro "slice" (lista) que simula la tabla de la base de datos
	Tecnicos []models.Tecnico

	// SiguienteID es un contador para asignar un ID único a cada nuevo tecnico
	SiguienteID = 1
)

package storage

import "proyecto/internal/models"

type TecnicoRepository interface {
	ListarTecnicos() []models.Tecnico
	BuscarTecnicoPorID(id int) (models.Tecnico, bool)
	CrearTecnico(t models.Tecnico) models.Tecnico
	ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool)
	BorrarTecnico(id int) bool
}
package service

import (
	"strings"
	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type TecnicoService struct {
	repo storage.TecnicoRepository
}

func NuevoTecnicoService(repo storage.TecnicoRepository) *TecnicoService {
	return &TecnicoService{repo: repo}
}

func (s *TecnicoService) Listar() []models.Tecnico {
	return s.repo.ListarTecnicos()
}

func (s *TecnicoService) Obtener(id int) (models.Tecnico, error) {
	t, ok := s.repo.BuscarTecnicoPorID(id)
	if !ok {
		return models.Tecnico{}, ErrNoEncontrado
	}
	return t, nil
}

func (s *TecnicoService) Crear(t models.Tecnico) (models.Tecnico, error) {
	if err := validarTecnico(t); err != nil {
		return models.Tecnico{}, err
	}
	t.Reputacion = 5.0 // Regla de negocio: reputación inicial
	return s.repo.CrearTecnico(t), nil
}

func (s *TecnicoService) Actualizar(id int, datos models.Tecnico) (models.Tecnico, error) {
	if err := validarTecnico(datos); err != nil {
		return models.Tecnico{}, err
	}
	actualizado, ok := s.repo.ActualizarTecnico(id, datos)
	if !ok {
		return models.Tecnico{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *TecnicoService) Borrar(id int) error {
	if !s.repo.BorrarTecnico(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarTecnico(t models.Tecnico) error {
	if strings.TrimSpace(t.Nombre) == "" {
		return ErrNombreVacio
	}
	if len(t.Servicios) == 0 {
		return ErrSinServicios
	}
	return nil
}
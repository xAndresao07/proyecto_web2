package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"
)

type SoporteService struct {
	repo storage.SoporteRepository
}

func NuevoSoporteService(repo storage.SoporteRepository) *SoporteService {
	return &SoporteService{repo: repo}
}

func (s *SoporteService) Listar() []models.Soporte {
	return s.repo.ListarSoportes()
}

func (s *SoporteService) Obtener(id int) (models.Soporte, error) {
	sop, ok := s.repo.BuscarSoportePorID(id)
	if !ok {
		return models.Soporte{}, ErrNoEncontrado
	}
	return sop, nil
}

func (s *SoporteService) Crear(sop models.Soporte) (models.Soporte, error) {
	if err := validacionSoporte(sop); err != nil {
		return models.Soporte{}, err
	}
	if strings.TrimSpace(sop.PiezasCambiadas) == "" {
		sop.PiezasCambiadas = "Ninguna"
	}
	return s.repo.CrearSoporte(sop), nil
}

func (s *SoporteService) Actualizar(id int, sop models.Soporte) (models.Soporte, error) {
	if err := validacionSoporte(sop); err != nil {
		return models.Soporte{}, err
	}
	actualizado, ok := s.repo.ActualizarSoporte(id, sop)
	if !ok {
		return models.Soporte{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *SoporteService) Borrar(id int) error {
	if !s.repo.BorrarSoporte(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionSoporte(s models.Soporte) error {
	if s.CitaID <= 0 || s.DispositivoID <= 0 || strings.TrimSpace(s.Solucion) == "" {
		return ErrNombreVacio
	}
	return nil
}

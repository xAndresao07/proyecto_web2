package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"
)

type SolicitanteService struct {
	repo storage.SolicitanteRepository
}

func NuevoSolicitanteService(repo storage.SolicitanteRepository) *SolicitanteService {
	return &SolicitanteService{repo: repo}
}

func (s *SolicitanteService) Listar() []models.Solicitante {
	return s.repo.ListarSolicitantes()
}

func (s *SolicitanteService) Obtener(id int) (models.Solicitante, error) {
	sol, ok := s.repo.BuscarSolicitantePorID(id)
	if !ok {
		return models.Solicitante{}, ErrNoEncontrado
	}
	return sol, nil
}

func (s *SolicitanteService) Crear(sol models.Solicitante) (models.Solicitante, error) {
	if err := validacionSolicitante(sol); err != nil {
		return models.Solicitante{}, err
	}
	if strings.TrimSpace(sol.NivelUrgencia) == "" {
		sol.NivelUrgencia = "media"
	}
	return s.repo.CrearSolicitante(sol), nil
}

func (s *SolicitanteService) Actualizar(id int, sol models.Solicitante) (models.Solicitante, error) {
	if err := validacionSolicitante(sol); err != nil {
		return models.Solicitante{}, err
	}
	actualizado, ok := s.repo.ActualizarSolicitante(id, sol)
	if !ok {
		return models.Solicitante{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *SolicitanteService) Borrar(id int) error {
	if !s.repo.BorrarSolicitante(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionSolicitante(s models.Solicitante) error {
	// Eliminamos 'int(s.ID) == 0 ||' de la validación
	if strings.TrimSpace(s.Nombre) == "" || strings.TrimSpace(s.Facultad) == "" || s.Semestre <= 0 {
		return ErrNombreVacio // Asegúrate de que este error devuelva el mensaje "los campos obligatorios no pueden estar vacios"
	}
	return nil
}

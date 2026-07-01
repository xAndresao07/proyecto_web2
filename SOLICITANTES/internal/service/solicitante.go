package service

import (
	"strings"

	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/storage"
)

type SolicitanteService struct {
	repo storage.SolicitanteRepositorio
}

func NewSolicitanteService(repo storage.SolicitanteRepositorio) *SolicitanteService {
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
	if err := validarSolicitante(sol); err != nil {
		return models.Solicitante{}, err
	}
	if sol.NivelUrgencia == "" {
		sol.NivelUrgencia = "normal"
	}
	return s.repo.CrearSolicitante(sol), nil
}

func (s *SolicitanteService) Actualizar(id int, datos models.Solicitante) (models.Solicitante, error) {
	if err := validarSolicitante(datos); err != nil {
		return models.Solicitante{}, err
	}
	actualizado, ok := s.repo.ActualizarSolicitante(id, datos)
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

func validarSolicitante(sol models.Solicitante) error {
	if strings.TrimSpace(sol.Matricula) == "" {
		return ErrMatriculaVacia
	}
	if strings.TrimSpace(sol.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(sol.Facultad) == "" {
		return ErrFacultadVacia
	}
	if sol.Semestre <= 0 {
		return ErrSemestreInvalido
	}
	if sol.NivelUrgencia != "" {
		switch sol.NivelUrgencia {
		case "normal", "alto", "critico":
		default:
			return ErrNivelUrgenciaInvalido
		}
	}
	return nil
}

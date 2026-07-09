package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"
)

type CitaService struct {
	repo storage.CitaRepository
}

func NuevoCitaService(repo storage.CitaRepository) *CitaService {
	return &CitaService{repo: repo}
}

func (s *CitaService) Listar() []models.Cita {
	return s.repo.ListarCitas()
}

func (s *CitaService) Obtener(id int) (models.Cita, error) {
	c, ok := s.repo.BuscarCitaPorID(id)
	if !ok {
		return models.Cita{}, ErrNoEncontrado
	}
	return c, nil
}

func (s *CitaService) Crear(c models.Cita) (models.Cita, error) {
	if err := validacionCita(c); err != nil {
		return models.Cita{}, err
	}
	if strings.TrimSpace(c.Estado) == "" {
		c.Estado = "pendiente"
	}
	return s.repo.CrearCita(c), nil
}

func (s *CitaService) Actualizar(id int, c models.Cita) (models.Cita, error) {
	if err := validacionCita(c); err != nil {
		return models.Cita{}, err
	}
	actualizado, ok := s.repo.ActualizarCita(id, c)
	if !ok {
		return models.Cita{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *CitaService) Borrar(id int) error {
	if !s.repo.BorrarCita(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionCita(c models.Cita) error {
	if strings.TrimSpace(c.SolicitanteID) == "" || strings.TrimSpace(c.TecnicoID) == "" || strings.TrimSpace(c.HoraAcordada) == "" || strings.TrimSpace(c.PuntoEncuentro) == "" {
		return ErrNombreVacio
	}
	return nil
}

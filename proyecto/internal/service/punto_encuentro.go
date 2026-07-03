package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"
)

type PuntoEncuentroService struct {
	repo storage.PuntoEncuentroRepository
}

func NuevoPuntoEncuentroService(repo storage.PuntoEncuentroRepository) *PuntoEncuentroService {
	return &PuntoEncuentroService{repo: repo}
}

func (s *PuntoEncuentroService) Listar() []models.PuntoEncuentro {
	return s.repo.ListarPuntosEncuentro()
}

func (s *PuntoEncuentroService) Obtener(id int) (models.PuntoEncuentro, error) {
	p, ok := s.repo.BuscarPuntoEncuentroPorID(id)
	if !ok {
		return models.PuntoEncuentro{}, ErrNoEncontrado
	}
	return p, nil
}

func (s *PuntoEncuentroService) Crear(p models.PuntoEncuentro) (models.PuntoEncuentro, error) {
	if err := validacionPuntoEncuentro(p); err != nil {
		return models.PuntoEncuentro{}, err
	}
	return s.repo.CrearPuntoEncuentro(p), nil
}

func (s *PuntoEncuentroService) Actualizar(id int, p models.PuntoEncuentro) (models.PuntoEncuentro, error) {
	if err := validacionPuntoEncuentro(p); err != nil {
		return models.PuntoEncuentro{}, err
	}
	actualizado, ok := s.repo.ActualizarPuntoEncuentro(id, p)
	if !ok {
		return models.PuntoEncuentro{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *PuntoEncuentroService) Borrar(id int) error {
	if !s.repo.BorrarPuntoEncuentro(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionPuntoEncuentro(p models.PuntoEncuentro) error {
	if strings.TrimSpace(p.NombreLugar) == "" || strings.TrimSpace(p.FacultadPerteneciente) == "" {
		return ErrNombreVacio
	}
	return nil
}

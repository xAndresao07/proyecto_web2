package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"
)

type DispositivoService struct {
	repo storage.DispositivoRepository
}

func NuevoDispositivoService(repo storage.DispositivoRepository) *DispositivoService {
	return &DispositivoService{repo: repo}
}

func (s *DispositivoService) Listar() []models.Dispositivo {
	return s.repo.ListarDispositivos()
}

func (s *DispositivoService) Obtener(id int) (models.Dispositivo, error) {
	d, ok := s.repo.BuscarDispositivoPorID(id)
	if !ok {
		return models.Dispositivo{}, ErrNoEncontrado
	}
	return d, nil
}

func (s *DispositivoService) Crear(d models.Dispositivo) (models.Dispositivo, error) {
	if err := validacionDispositivo(d); err != nil {
		return models.Dispositivo{}, err
	}
	return s.repo.CrearDispositivo(d), nil
}

func (s *DispositivoService) Actualizar(id int, d models.Dispositivo) (models.Dispositivo, error) {
	if err := validacionDispositivo(d); err != nil {
		return models.Dispositivo{}, err
	}
	actualizado, ok := s.repo.ActualizarDispositivo(id, d)
	if !ok {
		return models.Dispositivo{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *DispositivoService) Borrar(id int) error {
	if !s.repo.BorrarDispositivo(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionDispositivo(d models.Dispositivo) error {
	if int(d.ID) == 0 || int(d.SolicitanteID) == 0 || strings.TrimSpace(d.Marca) == "" || strings.TrimSpace(d.Modelo) == "" {
		return ErrNombreVacio
	}
	return nil
}

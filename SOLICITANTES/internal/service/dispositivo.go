package service

import (
	"strings"

	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/storage"
)

type DispositivoService struct {
	repo storage.DispositivoRepositorio
}

func NewDispositivoService(repo storage.DispositivoRepositorio) *DispositivoService {
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
	if err := validarDispositivo(d); err != nil {
		return models.Dispositivo{}, err
	}
	return s.repo.CrearDispositivo(d), nil
}

func (s *DispositivoService) Actualizar(id int, datos models.Dispositivo) (models.Dispositivo, error) {
	if err := validarDispositivoParcial(datos); err != nil {
		return models.Dispositivo{}, err
	}
	actualizado, ok := s.repo.ActualizarDispositivo(id, datos)
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

func validarDispositivo(d models.Dispositivo) error {
	if d.SolicitanteID == 0 {
		return ErrSolicitanteIDInvalido
	}
	return validarDispositivoParcial(d)
}

// validarDispositivoParcial valida los campos propios del dispositivo
// (sin exigir SolicitanteID, que en un PUT normalmente ya existe).
func validarDispositivoParcial(d models.Dispositivo) error {
	if strings.TrimSpace(d.Marca) == "" {
		return ErrMarcaVacia
	}
	if strings.TrimSpace(d.Modelo) == "" {
		return ErrModeloVacio
	}
	switch d.TipoAlmacenamiento {
	case "HDD", "SSD", "NVMe":
	default:
		return ErrTipoAlmacenamientoInvalido
	}
	if d.RamGB <= 0 {
		return ErrRamInvalida
	}
	if strings.TrimSpace(d.SistemaOperativo) == "" {
		return ErrSistemaOperativoVacio
	}
	return nil
}

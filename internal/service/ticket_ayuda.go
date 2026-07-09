package service

import (
	"proyecto/internal/models"
	"proyecto/internal/storage"
	"strings"
)

type TicketAyudaService struct {
	repo storage.TicketAyudaRepository
}

func NuevoTicketAyudaService(repo storage.TicketAyudaRepository) *TicketAyudaService {
	return &TicketAyudaService{repo: repo}
}

func (s *TicketAyudaService) Listar() []models.TicketAyuda {
	return s.repo.ListarTickets()
}

func (s *TicketAyudaService) Obtener(id int) (models.TicketAyuda, error) {
	t, ok := s.repo.BuscarTicketPorID(id)
	if !ok {
		return models.TicketAyuda{}, ErrNoEncontrado
	}
	return t, nil
}

func (s *TicketAyudaService) Crear(t models.TicketAyuda) (models.TicketAyuda, error) {
	if err := validacionTicket(t); err != nil {
		return models.TicketAyuda{}, err
	}
	if strings.TrimSpace(t.EstadoTicket) == "" {
		t.EstadoTicket = "abierto"
	}
	if strings.TrimSpace(t.SoftwareRequerido) == "" {
		t.SoftwareRequerido = "Ninguno"
	}
	return s.repo.CrearTicket(t), nil
}

func (s *TicketAyudaService) Actualizar(id int, t models.TicketAyuda) (models.TicketAyuda, error) {
	if err := validacionTicket(t); err != nil {
		return models.TicketAyuda{}, err
	}
	actualizado, ok := s.repo.ActualizarTicket(id, t)
	if !ok {
		return models.TicketAyuda{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *TicketAyudaService) Borrar(id int) error {
	if !s.repo.BorrarTicket(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validacionTicket(t models.TicketAyuda) error {
	if int(t.SolicitanteID) == 0 || int(t.DispositivoID) == 0 || strings.TrimSpace(t.DescripcionFalla) == "" {
		return ErrNombreVacio
	}
	return nil
}

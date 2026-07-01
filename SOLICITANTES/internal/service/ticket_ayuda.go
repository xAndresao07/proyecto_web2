package service

import (
	"strings"

	"solicitantesYHardware/internal/models"
	"solicitantesYHardware/internal/storage"
)

type TicketAyudaService struct {
	repo storage.TicketAyudaRepositorio
}

func NewTicketAyudaService(repo storage.TicketAyudaRepositorio) *TicketAyudaService {
	return &TicketAyudaService{repo: repo}
}

func (s *TicketAyudaService) Listar() []models.TicketAyuda {
	return s.repo.ListarTicketAyudas()
}

func (s *TicketAyudaService) Obtener(id int) (models.TicketAyuda, error) {
	t, ok := s.repo.BuscarTicketAyudaPorID(id)
	if !ok {
		return models.TicketAyuda{}, ErrNoEncontrado
	}
	return t, nil
}

// Crear registra un nuevo ticket. El estado inicial siempre se fuerza a
// "abierto", sin importar lo que venga en el body.
func (s *TicketAyudaService) Crear(t models.TicketAyuda) (models.TicketAyuda, error) {
	if t.SolicitanteID == 0 {
		return models.TicketAyuda{}, ErrSolicitanteIDInvalido
	}
	if t.DispositivoID == 0 {
		return models.TicketAyuda{}, ErrDispositivoIDInvalido
	}
	if strings.TrimSpace(t.DescripcionFalla) == "" {
		return models.TicketAyuda{}, ErrDescripcionFallaVacia
	}
	t.EstadoTicket = "abierto"
	return s.repo.CrearTicketAyuda(t), nil
}

func (s *TicketAyudaService) Actualizar(id int, datos models.TicketAyuda) (models.TicketAyuda, error) {
	if datos.DispositivoID == 0 {
		return models.TicketAyuda{}, ErrDispositivoIDInvalido
	}
	if strings.TrimSpace(datos.DescripcionFalla) == "" {
		return models.TicketAyuda{}, ErrDescripcionFallaVacia
	}
	switch datos.EstadoTicket {
	case "abierto", "en_proceso", "cerrado", "cancelado":
	default:
		return models.TicketAyuda{}, ErrEstadoTicketInvalido
	}
	actualizado, ok := s.repo.ActualizarTicketAyuda(id, datos)
	if !ok {
		return models.TicketAyuda{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *TicketAyudaService) Borrar(id int) error {
	if !s.repo.BorrarTicketAyuda(id) {
		return ErrNoEncontrado
	}
	return nil
}

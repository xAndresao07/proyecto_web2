package service

import (
	"testing"

	"solicitantesYHardware/internal/models"
)

type mockTicketRepo struct {
	crearFueLlamado bool
}

func (m *mockTicketRepo) ListarTicketAyudas() []models.TicketAyuda {
	return nil
}

func (m *mockTicketRepo) BuscarTicketAyudaPorID(id int) (models.TicketAyuda, bool) {
	return models.TicketAyuda{}, false
}

func (m *mockTicketRepo) CrearTicketAyuda(t models.TicketAyuda) models.TicketAyuda {
	m.crearFueLlamado = true
	t.ID = 1
	return t
}

func (m *mockTicketRepo) ActualizarTicketAyuda(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) {
	return models.TicketAyuda{}, false
}

func (m *mockTicketRepo) BorrarTicketAyuda(id int) bool {
	return false
}

func TestCrearTicket_RechazaFaltaDispositivo(t *testing.T) {
	mock := &mockTicketRepo{}
	svc := NewTicketAyudaService(mock)

	invalido := models.TicketAyuda{
		SolicitanteID:    1,
		DispositivoID:    0, // inválido
		DescripcionFalla: "Falla",
	}

	_, err := svc.Crear(invalido)

	if err != ErrDispositivoIDInvalido {
		t.Fatalf("se esperaba ErrDispositivoIDInvalido, se obtuvo: %v", err)
	}

	if mock.crearFueLlamado {
		t.Fatal("CrearTicketAyuda fue llamado con datos inválidos")
	}
}

func TestCrearTicket_AceptaValido(t *testing.T) {
	mock := &mockTicketRepo{}
	svc := NewTicketAyudaService(mock)

	valido := models.TicketAyuda{
		SolicitanteID:    1,
		DispositivoID:    1,
		DescripcionFalla: "Falla",
	}

	creado, err := svc.Crear(valido)

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if !mock.crearFueLlamado {
		t.Fatal("CrearTicketAyuda debió ser llamado")
	}
	if creado.EstadoTicket != "abierto" {
		t.Fatalf("estado inicial debió forzarse a abierto, pero fue %s", creado.EstadoTicket)
	}
}

func TestTicketAyuda_OtrasOperaciones(t *testing.T) {
	mock := &mockTicketRepo{}
	svc := NewTicketAyudaService(mock)

	// Listar
	lista := svc.Listar()
	if lista != nil {
		t.Fatal("Listar debio retornar nil desde el mock")
	}

	// Obtener no encontrado
	_, err := svc.Obtener(99)
	if err != ErrNoEncontrado {
		t.Fatalf("Obtener debio retornar ErrNoEncontrado, obtuvo: %v", err)
	}

	// Borrar no encontrado
	err = svc.Borrar(99)
	if err != ErrNoEncontrado {
		t.Fatalf("Borrar debio retornar ErrNoEncontrado, obtuvo: %v", err)
	}

	// Actualizar inválido
	invalido := models.TicketAyuda{
		DispositivoID:    1,
		DescripcionFalla: "Falla",
		EstadoTicket:     "inexistente",
	}
	_, err = svc.Actualizar(1, invalido)
	if err != ErrEstadoTicketInvalido {
		t.Fatalf("Actualizar debio retornar ErrEstadoTicketInvalido, obtuvo: %v", err)
	}
}

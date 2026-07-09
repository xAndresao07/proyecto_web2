package service

import (
	"testing"

	"solicitantesYHardware/internal/models"
)

type mockDispositivoRepo struct {
	crearFueLlamado bool
}

func (m *mockDispositivoRepo) ListarDispositivos() []models.Dispositivo {
	return nil
}

func (m *mockDispositivoRepo) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	return models.Dispositivo{}, false
}

func (m *mockDispositivoRepo) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	m.crearFueLlamado = true
	d.ID = 1
	return d
}

func (m *mockDispositivoRepo) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	return models.Dispositivo{}, false
}

func (m *mockDispositivoRepo) BorrarDispositivo(id int) bool {
	return false
}

func TestCrearDispositivo_RechazaTipoAlmacenamientoInvalido(t *testing.T) {
	mock := &mockDispositivoRepo{}
	svc := NewDispositivoService(mock)

	invalido := models.Dispositivo{
		SolicitanteID:      1,
		Marca:              "Dell",
		Modelo:             "Inspiron",
		TipoAlmacenamiento: "Pendrive", // inválido
		RamGB:              16,
		SistemaOperativo:   "Linux",
	}

	_, err := svc.Crear(invalido)

	if err != ErrTipoAlmacenamientoInvalido {
		t.Fatalf("se esperaba ErrTipoAlmacenamientoInvalido, se obtuvo: %v", err)
	}

	if mock.crearFueLlamado {
		t.Fatal("CrearDispositivo fue llamado con datos inválidos")
	}
}

func TestCrearDispositivo_AceptaValido(t *testing.T) {
	mock := &mockDispositivoRepo{}
	svc := NewDispositivoService(mock)

	valido := models.Dispositivo{
		SolicitanteID:      1,
		Marca:              "Dell",
		Modelo:             "Inspiron",
		TipoAlmacenamiento: "SSD",
		RamGB:              16,
		SistemaOperativo:   "Linux",
	}

	_, err := svc.Crear(valido)

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if !mock.crearFueLlamado {
		t.Fatal("CrearDispositivo debió ser llamado")
	}
}

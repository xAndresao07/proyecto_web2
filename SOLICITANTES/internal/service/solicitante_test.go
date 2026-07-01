package service

import (
	"testing"

	"solicitantesYHardware/internal/models"
)

// mockSolicitanteRepo es un DOBLE de prueba que implementa
// storage.SolicitanteRepositorio sin tocar ninguna base de datos real.
//
// Su única función es DETECTAR si el service llegó a llamar CrearSolicitante.
// Si la validación funciona bien, ese método nunca debe ejecutarse cuando
// el dato es inválido.
type mockSolicitanteRepo struct {
	crearFueLlamado bool
}

func (m *mockSolicitanteRepo) ListarSolicitantes() []models.Solicitante {
	return nil
}

func (m *mockSolicitanteRepo) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	return models.Solicitante{}, false
}

func (m *mockSolicitanteRepo) CrearSolicitante(s models.Solicitante) models.Solicitante {
	m.crearFueLlamado = true // marcamos que SÍ se intentó guardar
	s.ID = 1
	return s
}

func (m *mockSolicitanteRepo) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	return models.Solicitante{}, false
}

func (m *mockSolicitanteRepo) BorrarSolicitante(id int) bool {
	return false
}

// TestCrearSolicitante_RechazaNivelUrgenciaInvalido prueba la regla de negocio:
// un nivel_urgencia que no sea "normal", "alto" o "critico" debe ser RECHAZADO
// por el service, y el repositorio NUNCA debe llegar a guardarlo.
//
// Qué se rompería si la implementación fallara: si alguien borra el switch de
// validarSolicitante (o lo comenta "para probar rápido"), este test detecta
// que un dato basura como "urgentisimo" pasaría directo a la base de datos.
func TestCrearSolicitante_RechazaNivelUrgenciaInvalido(t *testing.T) {
	mock := &mockSolicitanteRepo{}
	svc := NewSolicitanteService(mock)

	solicitanteInvalido := models.Solicitante{
		UsuarioID:     1,
		Matricula:     "ULEAM-0099",
		Nombre:        "Estudiante de Prueba",
		Facultad:      "TI",
		Semestre:      4,
		NivelUrgencia: "urgentisimo", // valor que NO existe en el dominio
	}

	_, err := svc.Crear(solicitanteInvalido)

	// 1. El service debe devolver el error específico de la regla.
	if err != ErrNivelUrgenciaInvalido {
		t.Fatalf("se esperaba ErrNivelUrgenciaInvalido, se obtuvo: %v", err)
	}

	// 2. El repositorio NUNCA debió ser llamado: el dato inválido se frena
	//    en el service, antes de llegar a la capa de persistencia.
	if mock.crearFueLlamado {
		t.Fatal("CrearSolicitante del repositorio fue llamado con datos inválidos; " +
			"la validación debió detener el flujo antes de guardar")
	}
}

// TestCrearSolicitante_AceptaNivelUrgenciaValido es el caso de control:
// confirma que un dato BUENO sí llega al repositorio. Sin este test, el
// anterior podría estar "pasando en verde" solo porque el mock siempre
// rechaza todo, no porque la regla esté bien escrita.
func TestCrearSolicitante_AceptaNivelUrgenciaValido(t *testing.T) {
	mock := &mockSolicitanteRepo{}
	svc := NewSolicitanteService(mock)

	solicitanteValido := models.Solicitante{
		UsuarioID:     1,
		Matricula:     "ULEAM-0100",
		Nombre:        "Estudiante Valido",
		Facultad:      "TI",
		Semestre:      4,
		NivelUrgencia: "alto",
	}

	_, err := svc.Crear(solicitanteValido)

	if err != nil {
		t.Fatalf("no se esperaba error con datos válidos, se obtuvo: %v", err)
	}
	if !mock.crearFueLlamado {
		t.Fatal("CrearSolicitante del repositorio debió ser llamado con datos válidos")
	}
}

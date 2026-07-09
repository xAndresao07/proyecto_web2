package service_test

import (
	"testing"

	"proyecto/internal/models"
	"proyecto/internal/service"
)

// 1. Creamos un Mock (un doble de riesgo) de la base de datos
type mockRepo struct{}

func (m *mockRepo) ListarTecnicos() []models.Tecnico                 { return nil }
func (m *mockRepo) BuscarTecnicoPorID(id int) (models.Tecnico, bool) { return models.Tecnico{}, false }
func (m *mockRepo) ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool) {
	return models.Tecnico{}, false
}
func (m *mockRepo) BorrarTecnico(id int) bool { return false }

// Solo simulamos que "guarda" devolviendo el mismo objeto con un ID
func (m *mockRepo) CrearTecnico(t models.Tecnico) models.Tecnico {
	t.ID = 1
	return t
}

func TestCrearTecnico_ReglasDeNegocio(t *testing.T) {
	// Inyectamos el mock, ¡aislando el servicio de la BD real!
	repo := &mockRepo{}
	svc := service.NuevoTecnicoService(repo)

	// Regla 1: Falla si el nombre está vacío
	_, err := svc.Crear(models.Tecnico{Nombre: "", Servicios: []models.ServicioOfrecido{{NombreServicio: "Formateo"}}})
	if err != service.ErrNombreVacio {
		t.Errorf("se esperaba %v, se obtuvo: %v", service.ErrNombreVacio, err)
	}

	// Regla 2: Falla si no tiene servicios
	_, err = svc.Crear(models.Tecnico{Nombre: "Juan", Servicios: []models.ServicioOfrecido{}})
	if err != service.ErrSinServicios {
		t.Errorf("se esperaba %v, se obtuvo: %v", service.ErrSinServicios, err)
	}

	// Regla 3: Si todo es válido, debe asignar reputación inicial de 5.0
	nuevo, err := svc.Crear(models.Tecnico{Nombre: "Juan", Servicios: []models.ServicioOfrecido{{NombreServicio: "Formateo"}}})
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo: %v", err)
	}
	if nuevo.Reputacion != 5.0 {
		t.Errorf("se esperaba reputación de 5.0, se obtuvo: %f", nuevo.Reputacion)
	}
}

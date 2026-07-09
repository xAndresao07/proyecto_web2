package storage

import (
	"testing"
	"proyecto/internal/models"
)

func TestMemoria_Solicitantes(t *testing.T) {
	m := NuevaMemoria()
	s := models.Solicitante{Nombre: "Test"}
	
	creado := m.CrearSolicitante(s)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarSolicitantes()
	if len(lista) != 1 {
		t.Error("No se listó el solicitante")
	}

	buscado, ok := m.BuscarSolicitantePorID(creado.ID)
	if !ok || buscado.Nombre != "Test" {
		t.Error("No se encontró el solicitante o el nombre no coincide")
	}

	buscado, ok = m.BuscarSolicitantePorID(999)
	if ok {
		t.Error("No debería encontrar solicitante")
	}

	actualizado, ok := m.ActualizarSolicitante(creado.ID, models.Solicitante{Nombre: "Actualizado"})
	if !ok || actualizado.Nombre != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarSolicitante(999, models.Solicitante{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarSolicitante(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarSolicitante(999)
	if ok {
		t.Error("No debería borrar")
	}

	m.SeedSolicitantes()
	if len(m.ListarSolicitantes()) == 0 {
		t.Error("Seed falló")
	}
}

func TestMemoria_Dispositivos(t *testing.T) {
	m := NuevaMemoria()
	d := models.Dispositivo{Marca: "Test"}
	
	creado := m.CrearDispositivo(d)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarDispositivos()
	if len(lista) != 1 {
		t.Error("No se listó el dispositivo")
	}

	buscado, ok := m.BuscarDispositivoPorID(creado.ID)
	if !ok || buscado.Marca != "Test" {
		t.Error("No se encontró o no coincide")
	}

	buscado, ok = m.BuscarDispositivoPorID(999)
	if ok {
		t.Error("No debería encontrar")
	}

	actualizado, ok := m.ActualizarDispositivo(creado.ID, models.Dispositivo{Marca: "Actualizado"})
	if !ok || actualizado.Marca != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarDispositivo(999, models.Dispositivo{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarDispositivo(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarDispositivo(999)
	if ok {
		t.Error("No debería borrar")
	}

	m.SeedDispositivos()
	if len(m.ListarDispositivos()) == 0 {
		t.Error("Seed falló")
	}
}

func TestMemoria_Tickets(t *testing.T) {
	m := NuevaMemoria()
	tck := models.TicketAyuda{DescripcionFalla: "Test"}
	
	creado := m.CrearTicket(tck)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarTickets()
	if len(lista) != 1 {
		t.Error("No se listó")
	}

	buscado, ok := m.BuscarTicketPorID(creado.ID)
	if !ok || buscado.DescripcionFalla != "Test" {
		t.Error("No se encontró o no coincide")
	}

	buscado, ok = m.BuscarTicketPorID(999)
	if ok {
		t.Error("No debería encontrar")
	}

	actualizado, ok := m.ActualizarTicket(creado.ID, models.TicketAyuda{DescripcionFalla: "Actualizado"})
	if !ok || actualizado.DescripcionFalla != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarTicket(999, models.TicketAyuda{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarTicket(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarTicket(999)
	if ok {
		t.Error("No debería borrar")
	}

	m.SeedTickets()
	if len(m.ListarTickets()) == 0 {
		t.Error("Seed falló")
	}
}

func TestMemoria_Tecnicos(t *testing.T) {
	m := NuevaMemoria()
	tc := models.Tecnico{Nombre: "Test"}
	
	creado := m.CrearTecnico(tc)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarTecnicos()
	if len(lista) != 1 {
		t.Error("No se listó")
	}

	buscado, ok := m.BuscarTecnicoPorID(creado.ID)
	if !ok || buscado.Nombre != "Test" {
		t.Error("No se encontró o no coincide")
	}

	buscado, ok = m.BuscarTecnicoPorID(999)
	if ok {
		t.Error("No debería encontrar")
	}

	actualizado, ok := m.ActualizarTecnico(creado.ID, models.Tecnico{Nombre: "Actualizado"})
	if !ok || actualizado.Nombre != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarTecnico(999, models.Tecnico{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarTecnico(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarTecnico(999)
	if ok {
		t.Error("No debería borrar")
	}
}

func TestMemoria_Citas(t *testing.T) {
	m := NuevaMemoria()
	c := models.Cita{Estado: "Test"}
	
	creado := m.CrearCita(c)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarCitas()
	if len(lista) != 1 {
		t.Error("No se listó")
	}

	buscado, ok := m.BuscarCitaPorID(creado.ID)
	if !ok || buscado.Estado != "Test" {
		t.Error("No se encontró o no coincide")
	}

	buscado, ok = m.BuscarCitaPorID(999)
	if ok {
		t.Error("No debería encontrar")
	}

	actualizado, ok := m.ActualizarCita(creado.ID, models.Cita{Estado: "Actualizado"})
	if !ok || actualizado.Estado != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarCita(999, models.Cita{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarCita(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarCita(999)
	if ok {
		t.Error("No debería borrar")
	}

	m.SeedCitas()
	if len(m.ListarCitas()) == 0 {
		t.Error("Seed falló")
	}
}

func TestMemoria_Puntos(t *testing.T) {
	m := NuevaMemoria()
	p := models.PuntoEncuentro{NombreLugar: "Test"}
	
	creado := m.CrearPuntoEncuentro(p)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarPuntosEncuentro()
	if len(lista) != 1 {
		t.Error("No se listó")
	}

	buscado, ok := m.BuscarPuntoEncuentroPorID(creado.ID)
	if !ok || buscado.NombreLugar != "Test" {
		t.Error("No se encontró o no coincide")
	}

	buscado, ok = m.BuscarPuntoEncuentroPorID(999)
	if ok {
		t.Error("No debería encontrar")
	}

	actualizado, ok := m.ActualizarPuntoEncuentro(creado.ID, models.PuntoEncuentro{NombreLugar: "Actualizado"})
	if !ok || actualizado.NombreLugar != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarPuntoEncuentro(999, models.PuntoEncuentro{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarPuntoEncuentro(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarPuntoEncuentro(999)
	if ok {
		t.Error("No debería borrar")
	}
}

func TestMemoria_Soportes(t *testing.T) {
	m := NuevaMemoria()
	s := models.Soporte{Solucion: "Test"}
	
	creado := m.CrearSoporte(s)
	if creado.ID == 0 {
		t.Error("ID no asignado")
	}

	lista := m.ListarSoportes()
	if len(lista) != 1 {
		t.Error("No se listó")
	}

	buscado, ok := m.BuscarSoportePorID(creado.ID)
	if !ok || buscado.Solucion != "Test" {
		t.Error("No se encontró o no coincide")
	}

	buscado, ok = m.BuscarSoportePorID(999)
	if ok {
		t.Error("No debería encontrar")
	}

	actualizado, ok := m.ActualizarSoporte(creado.ID, models.Soporte{Solucion: "Actualizado"})
	if !ok || actualizado.Solucion != "Actualizado" {
		t.Error("No se actualizó")
	}

	_, ok = m.ActualizarSoporte(999, models.Soporte{})
	if ok {
		t.Error("No debería actualizar")
	}

	ok = m.BorrarSoporte(creado.ID)
	if !ok {
		t.Error("No se borró")
	}

	ok = m.BorrarSoporte(999)
	if ok {
		t.Error("No debería borrar")
	}
}

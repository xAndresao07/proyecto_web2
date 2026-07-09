package storage

import (
	"proyecto/internal/models"

	"gorm.io/gorm"
)

type AlmacenSQLite struct{ db *gorm.DB }

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite { return &AlmacenSQLite{db: db} }

// --- Solicitantes ---
func (a *AlmacenSQLite) ListarSolicitantes() []models.Solicitante {
	var x []models.Solicitante
	a.db.Find(&x)
	return x
}
func (a *AlmacenSQLite) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	var s models.Solicitante
	if err := a.db.First(&s, "id = ?", id).Error; err != nil {
		return models.Solicitante{}, false
	}
	return s, true
}
func (a *AlmacenSQLite) CrearSolicitante(s models.Solicitante) models.Solicitante {
	a.db.Create(&s)
	return s
}
func (a *AlmacenSQLite) ActualizarSolicitante(id int, d models.Solicitante) (models.Solicitante, bool) {
	var s models.Solicitante
	if err := a.db.First(&s, "id = ?", id).Error; err != nil {
		return models.Solicitante{}, false
	}
	d.ID = id
	a.db.Save(&d)
	return d, true
}
func (a *AlmacenSQLite) BorrarSolicitante(id int) bool {
	return a.db.Delete(&models.Solicitante{}, "id = ?", id).RowsAffected > 0
}

// --- Dispositivos ---
func (a *AlmacenSQLite) ListarDispositivos() []models.Dispositivo {
	var x []models.Dispositivo
	a.db.Find(&x)
	return x
}
func (a *AlmacenSQLite) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	var d models.Dispositivo
	if err := a.db.First(&d, "id = ?", id).Error; err != nil {
		return models.Dispositivo{}, false
	}
	return d, true
}
func (a *AlmacenSQLite) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	a.db.Create(&d)
	return d
}
func (a *AlmacenSQLite) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	var d models.Dispositivo
	if err := a.db.First(&d, "id = ?", id).Error; err != nil {
		return models.Dispositivo{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}
func (a *AlmacenSQLite) BorrarDispositivo(id int) bool {
	return a.db.Delete(&models.Dispositivo{}, "id = ?", id).RowsAffected > 0
}

// --- Tickets de Ayuda ---
func (a *AlmacenSQLite) ListarTickets() []models.TicketAyuda {
	var x []models.TicketAyuda
	a.db.Find(&x)
	return x
}
func (a *AlmacenSQLite) BuscarTicketPorID(id int) (models.TicketAyuda, bool) {
	var t models.TicketAyuda
	if err := a.db.First(&t, "id = ?", id).Error; err != nil {
		return models.TicketAyuda{}, false
	}
	return t, true
}
func (a *AlmacenSQLite) CrearTicket(t models.TicketAyuda) models.TicketAyuda {
	a.db.Create(&t)
	return t
}
func (a *AlmacenSQLite) ActualizarTicket(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) {
	var t models.TicketAyuda
	if err := a.db.First(&t, "id = ?", id).Error; err != nil {
		return models.TicketAyuda{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}
func (a *AlmacenSQLite) BorrarTicket(id int) bool {
	return a.db.Delete(&models.TicketAyuda{}, id).RowsAffected > 0
}

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.Solicitante{}).Count(&n)
	if n > 0 {
		return
	}

	solicitantes := []models.Solicitante{
		{ID: 1, Nombre: "Jandry Ejemplo", Facultad: "FACCI", Semestre: 5, NivelUrgencia: "alta"},
	}
	a.db.Create(&solicitantes)

	dispositivos := []models.Dispositivo{
		{ID: 1, SolicitanteID: 1, Marca: "HP", Modelo: "Pavilion 14", TipoAlmacenamiento: "SSD", RamGB: 8, SistemaOperativo: "Windows 11"},
	}
	a.db.Create(&dispositivos)

	tickets := []models.TicketAyuda{
		{SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "No enciende", SoftwareRequerido: "Ninguno", EstadoTicket: "abierto"},
	}
	a.db.Create(&tickets)
}

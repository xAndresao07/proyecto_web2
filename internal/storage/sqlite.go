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

// --- Usuarios ---
func (a *AlmacenSQLite) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	err := a.db.Create(&u).Error
	return u, err
}

func (a *AlmacenSQLite) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	// Buscamos por email. Si falla (error != nil), devolvemos false.
	if err := a.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
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
		// ¡Aquí está el ID: 1 que faltaba!
		{ID: 1, SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "No enciende", SoftwareRequerido: "Ninguno", EstadoTicket: "abierto"},
	}
	a.db.Create(&tickets)

}


func (a *AlmacenSQLite) ListarTecnicos() []models.Tecnico {
	var tecnicos []models.Tecnico
	// Preload carga las relaciones (los arreglos de servicios y horarios)
	a.db.Preload("Servicios").Preload("Horarios").Find(&tecnicos)
	return tecnicos
}

func (a *AlmacenSQLite) BuscarTecnicoPorID(id int) (models.Tecnico, bool) {
	var t models.Tecnico
	if err := a.db.Preload("Servicios").Preload("Horarios").First(&t, id).Error; err != nil {
		return models.Tecnico{}, false
	}
	return t, true
}

func (a *AlmacenSQLite) CrearTecnico(t models.Tecnico) models.Tecnico {
	a.db.Create(&t) // GORM inserta el técnico y automáticamente sus servicios/horarios
	return t
}

func (a *AlmacenSQLite) ActualizarTecnico(id int, datos models.Tecnico) (models.Tecnico, bool) {
	var existente models.Tecnico
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Tecnico{}, false
	}
	
	datos.ID = id
	// Session con FullSaveAssociations obliga a GORM a actualizar también los arreglos anidados
	a.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarTecnico(id int) bool {
	// Select le dice a GORM que también borre los servicios y horarios asociados a este técnico
	res := a.db.Select("Servicios", "Horarios").Delete(&models.Tecnico{ID: id})
	return res.RowsAffected > 0
}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir TecnicoRepository.
var _ TecnicoRepository = (*AlmacenSQLite)(nil)

func (a *AlmacenSQLite) ListarCitas() []models.Cita { var x []models.Cita; a.db.Find(&x); return x }
func (a *AlmacenSQLite) BuscarCitaPorID(id int) (models.Cita, bool) {
	var c models.Cita
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cita{}, false
	}
	return c, true
}
func (a *AlmacenSQLite) CrearCita(c models.Cita) models.Cita { a.db.Create(&c); return c }
func (a *AlmacenSQLite) ActualizarCita(id int, d models.Cita) (models.Cita, bool) {
	var c models.Cita
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cita{}, false
	}
	d.ID = id
	a.db.Save(&d)
	return d, true
}
func (a *AlmacenSQLite) BorrarCita(id int) bool {
	return a.db.Delete(&models.Cita{}, id).RowsAffected > 0
}

func (a *AlmacenSQLite) ListarPuntosEncuentro() []models.PuntoEncuentro {
	var x []models.PuntoEncuentro
	a.db.Find(&x)
	return x
}
func (a *AlmacenSQLite) BuscarPuntoEncuentroPorID(id int) (models.PuntoEncuentro, bool) {
	var p models.PuntoEncuentro
	if err := a.db.First(&p, id).Error; err != nil {
		return models.PuntoEncuentro{}, false
	}
	return p, true
}
func (a *AlmacenSQLite) CrearPuntoEncuentro(p models.PuntoEncuentro) models.PuntoEncuentro {
	a.db.Create(&p)
	return p
}
func (a *AlmacenSQLite) ActualizarPuntoEncuentro(id int, d models.PuntoEncuentro) (models.PuntoEncuentro, bool) {
	var p models.PuntoEncuentro
	if err := a.db.First(&p, id).Error; err != nil {
		return models.PuntoEncuentro{}, false
	}
	d.ID = id
	a.db.Save(&d)
	return d, true
}
func (a *AlmacenSQLite) BorrarPuntoEncuentro(id int) bool {
	return a.db.Delete(&models.PuntoEncuentro{}, id).RowsAffected > 0
}

func (a *AlmacenSQLite) ListarSoportes() []models.Soporte {
	var x []models.Soporte
	a.db.Find(&x)
	return x
}
func (a *AlmacenSQLite) BuscarSoportePorID(id int) (models.Soporte, bool) {
	var s models.Soporte
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Soporte{}, false
	}
	return s, true
}
func (a *AlmacenSQLite) CrearSoporte(s models.Soporte) models.Soporte { a.db.Create(&s); return s }
func (a *AlmacenSQLite) ActualizarSoporte(id int, d models.Soporte) (models.Soporte, bool) {
	var s models.Soporte
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Soporte{}, false
	}
	d.ID = id
	a.db.Save(&d)
	return d, true
}
func (a *AlmacenSQLite) BorrarSoporte(id int) bool {
	return a.db.Delete(&models.Soporte{}, id).RowsAffected > 0
}


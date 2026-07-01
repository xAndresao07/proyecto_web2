package storage

import (
	"gorm.io/gorm"

	"solicitantesYHardware/internal/models"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
//
// Los métodos tienen EXACTAMENTE las mismas firmas que los de Memoria.
// Por eso el Server y los handlers no se enteran de cuál de los dos reciben.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// SOLICITANTES
// =========================================================

func (a *AlmacenSQLite) ListarSolicitantes() []models.Solicitante {
	var solicitantes []models.Solicitante
	a.db.Find(&solicitantes)
	return solicitantes
}

func (a *AlmacenSQLite) BuscarSolicitantePorID(id int) (models.Solicitante, bool) {
	var s models.Solicitante
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitante{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitante(s models.Solicitante) models.Solicitante {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLite) ActualizarSolicitante(id int, datos models.Solicitante) (models.Solicitante, bool) {
	var existente models.Solicitante
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Solicitante{}, false
	}
	datos.ID = id
	datos.UsuarioID = existente.UsuarioID
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitante(id int) bool {
	res := a.db.Delete(&models.Solicitante{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// DISPOSITIVOS
// =========================================================

func (a *AlmacenSQLite) ListarDispositivos() []models.Dispositivo {
	var dispositivos []models.Dispositivo
	a.db.Find(&dispositivos)
	return dispositivos
}

func (a *AlmacenSQLite) BuscarDispositivoPorID(id int) (models.Dispositivo, bool) {
	var d models.Dispositivo
	if err := a.db.First(&d, id).Error; err != nil {
		return models.Dispositivo{}, false
	}
	return d, true
}

func (a *AlmacenSQLite) CrearDispositivo(d models.Dispositivo) models.Dispositivo {
	a.db.Create(&d)
	return d
}

func (a *AlmacenSQLite) ActualizarDispositivo(id int, datos models.Dispositivo) (models.Dispositivo, bool) {
	var existente models.Dispositivo
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Dispositivo{}, false
	}
	datos.ID = id
	datos.SolicitanteID = existente.SolicitanteID
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarDispositivo(id int) bool {
	res := a.db.Delete(&models.Dispositivo{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// TICKET AYUDAS
// =========================================================

func (a *AlmacenSQLite) ListarTicketAyudas() []models.TicketAyuda {
	var tickets []models.TicketAyuda
	a.db.Find(&tickets)
	return tickets
}

func (a *AlmacenSQLite) BuscarTicketAyudaPorID(id int) (models.TicketAyuda, bool) {
	var t models.TicketAyuda
	if err := a.db.First(&t, id).Error; err != nil {
		return models.TicketAyuda{}, false
	}
	return t, true
}

func (a *AlmacenSQLite) CrearTicketAyuda(t models.TicketAyuda) models.TicketAyuda {
	a.db.Create(&t)
	return t
}

func (a *AlmacenSQLite) ActualizarTicketAyuda(id int, datos models.TicketAyuda) (models.TicketAyuda, bool) {
	var existente models.TicketAyuda
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.TicketAyuda{}, false
	}
	datos.ID = id
	datos.SolicitanteID = existente.SolicitanteID
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarTicketAyuda(id int) bool {
	res := a.db.Delete(&models.TicketAyuda{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

// SembrarSiVacio inserta datos iniciales solo si aún no hay solicitantes.
// Así no duplicamos datos en cada arranque del servidor.
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.Solicitante{}).Count(&n)
	if n > 0 {
		return
	}

	solicitantes := []models.Solicitante{
		{ID: 1, UsuarioID: 1, Matricula: "ULEAM-0001", Nombre: "Carlos Alberto Moreira", Facultad: "TI", Semestre: 4, NivelUrgencia: "alto"},
		{ID: 2, UsuarioID: 2, Matricula: "ULEAM-0002", Nombre: "Ricardo Villavicencio", Facultad: "Software", Semestre: 3, NivelUrgencia: "normal"},
	}
	a.db.Create(&solicitantes)

	dispositivos := []models.Dispositivo{
		{ID: 1, SolicitanteID: 1, Marca: "HP", Modelo: "Pavilion 14", TipoAlmacenamiento: "HDD", RamGB: 4, SistemaOperativo: "Windows 10"},
		{ID: 2, SolicitanteID: 2, Marca: "Lenovo", Modelo: "IdeaPad 3", TipoAlmacenamiento: "SSD", RamGB: 8, SistemaOperativo: "Windows 11"},
	}
	a.db.Create(&dispositivos)

	tickets := []models.TicketAyuda{
		{ID: 1, SolicitanteID: 1, DispositivoID: 1, DescripcionFalla: "PC se congela al abrir VS Code", SoftwareRequerido: "ninguno", EstadoTicket: "abierto"},
		{ID: 2, SolicitanteID: 2, DispositivoID: 2, DescripcionFalla: "Cambio de RAM solicitado", SoftwareRequerido: "ninguno", EstadoTicket: "abierto"},
	}
	a.db.Create(&tickets)
}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)

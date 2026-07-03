package storage

import (
	"proyecto/internal/models"

	"gorm.io/gorm"
)

type AlmacenSQLite struct{ db *gorm.DB }

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite { return &AlmacenSQLite{db: db} }

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

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.PuntoEncuentro{}).Count(&n)
	if n > 0 {
		return
	}
	citas := []models.Cita{
		{SolicitanteID: "est_102", TecnicoID: "tec_05", Estado: "pendiente", HoraAcordada: "09:00", PuntoEncuentro: "Lab CISCO"},
	}
	a.db.Create(&citas)

	punto := []models.PuntoEncuentro{
		{NombreLugar: "Lab CISCO", FacultadPerteneciente: "FACCI", DisponibleParaSoporte: true},
	}
	a.db.Create(&punto)

	soport := []models.Soporte{
		{ID: 1, CitaID: 1, DispositivoID: 501, Solucion: "Limpieza", PiezasCambiadas: "Ninguna"},
	}
	a.db.Create(&soport)
}

package storage

import (
	"gorm.io/gorm"
	"proyecto/internal/models"
)

// AlmacenSQLite implementa la interfaz TecnicoRepository usando GORM sobre SQLite.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
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
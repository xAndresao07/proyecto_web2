package models

// PuntoEncuentro gestiona los espacios físicos disponibles en el campus.

type PuntoEncuentro struct {
	ID                    int    `json:"id" gorm:"primaryKey"`
	NombreLugar           string `json:"nombre_lugar"`
	FacultadPerteneciente string `json:"facultad_perteneciente"`
	DisponibleParaSoporte bool   `json:"disponible_para_soporte"`
}

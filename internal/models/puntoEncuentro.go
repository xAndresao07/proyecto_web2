package models

// PuntoEncuentro gestiona los espacios físicos disponibles en el campus.
type PuntoEncuentro struct {
	ID                    int    `json:"id"`
	NombreLugar           string `json:"nombre_lugar"`
	FacultadPerteneciente string `json:"facultad_perteneciente"`
	AforoMaximo           int    `json:"aforo_maximo"`
	DisponibleParaSoporte bool   `json:"disponible_para_soporte"`
}

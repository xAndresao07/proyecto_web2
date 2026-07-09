package models

// Dispositivo representa el hardware registrado por un solicitante.

type Dispositivo struct {
	ID                 int    `json:"id" gorm:"primaryKey"` // Tag de inventario
	SolicitanteID      int    `json:"solicitante_id"`
	Marca              string `json:"marca"`
	Modelo             string `json:"modelo"`
	TipoAlmacenamiento string `json:"tipo_almacenamiento"`
	RamGB              int    `json:"ram_gb"`
	SistemaOperativo   string `json:"sistema_operativo"`
}

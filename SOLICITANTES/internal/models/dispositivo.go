package models

// Dispositivo representa la computadora (laptop o PC) que está fallando.
//
// La relación con Solicitante es por ID: Dispositivo.SolicitanteID apunta a
// Solicitante.ID. Un solicitante puede registrar varios dispositivos (1:N).
type Dispositivo struct {
	ID                 int    `json:"id" gorm:"primaryKey"`
	SolicitanteID      int    `json:"solicitante_id" gorm:"not null"`
	Marca              string `json:"marca" gorm:"not null"`
	Modelo             string `json:"modelo" gorm:"not null"`
	TipoAlmacenamiento string `json:"tipo_almacenamiento" gorm:"not null"`
	RamGB              int    `json:"ram_gb" gorm:"not null"`
	SistemaOperativo   string `json:"sistema_operativo" gorm:"not null"`

	// Relación (Belongs-To)
	Solicitante *Solicitante `json:"solicitante,omitempty" gorm:"foreignKey:SolicitanteID"`

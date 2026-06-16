package models

type ServicioOfrecido struct {
	NombreServicio   string `json:"nombre_servicio"`   // Ej: Formateo, Limpieza Térmica
	NivelExperiencia string `json:"nivel_experiencia"` // Ej: Básico, Intermedio, Avanzado
	TiempoEstimado   string `json:"tiempo_estimado"`   // Ej: 2 horas
}

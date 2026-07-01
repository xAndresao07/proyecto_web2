package service

import "errors"

var (
	ErrNoEncontrado          = errors.New("recurso no encontrado")
	ErrEmailEnUso            = errors.New("el email ya está registrado")
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")

	// Solicitante
	ErrMatriculaVacia        = errors.New("el campo matricula es obligatorio")
	ErrNombreVacio           = errors.New("el campo nombre es obligatorio")
	ErrFacultadVacia         = errors.New("el campo facultad es obligatorio")
	ErrSemestreInvalido      = errors.New("el semestre debe ser mayor a 0")
	ErrNivelUrgenciaInvalido = errors.New("nivel_urgencia debe ser: normal, alto o critico")

	// Dispositivo
	ErrMarcaVacia                 = errors.New("el campo marca es obligatorio")
	ErrModeloVacio                = errors.New("el campo modelo es obligatorio")
	ErrTipoAlmacenamientoInvalido = errors.New("tipo_almacenamiento debe ser: HDD, SSD o NVMe")
	ErrRamInvalida                = errors.New("ram_gb debe ser mayor a 0")
	ErrSistemaOperativoVacio      = errors.New("el campo sistema_operativo es obligatorio")
	ErrSolicitanteIDInvalido      = errors.New("solicitante_id es obligatorio")

	// TicketAyuda
	ErrDescripcionFallaVacia = errors.New("el campo descripcion_falla es obligatorio")
	ErrDispositivoIDInvalido = errors.New("dispositivo_id es obligatorio")
	ErrEstadoTicketInvalido  = errors.New("estado_ticket debe ser: abierto, en_proceso, cerrado o cancelado")
)

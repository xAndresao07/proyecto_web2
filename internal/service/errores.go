package service

import "errors"

var (
	ErrNombreVacio           = errors.New("el nombre del técnico es obligatorio")
	ErrSinServicios          = errors.New("debe especificar al menos un servicio ofrecido")
	ErrNoEncontrado          = errors.New("técnico no encontrado")
)
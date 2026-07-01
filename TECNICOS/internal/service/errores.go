package service

import "errors"

var (
	ErrNombreVacio  = errors.New("el nombre del técnico es obligatorio")
	ErrSinServicios = errors.New("debe especificar al menos un servicio ofrecido")
	ErrNoEncontrado = errors.New("recurso no encontrado")

	// Errores de Auth
	ErrEmailEnUso            = errors.New("el email ya esta registrado")
	ErrCredencialesInvalidas = errors.New("email o contrasena incorrectos")
)

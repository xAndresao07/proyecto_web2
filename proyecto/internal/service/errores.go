package service

import "errors"

var (
	ErrNombreVacio           = errors.New("los campos obligatorios no pueden estar vacios")
	ErrNoEncontrado          = errors.New("cita no encontrada")
	ErrEmailEnUso            = errors.New("el email ya esta registrado o no existe")
	ErrCredencialesInvalidas = errors.New("email o contrasena incorrectos")
	ErrIDInvalido            = errors.New("el id proporcionado no es valido")
)

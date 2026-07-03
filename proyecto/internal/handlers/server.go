package handlers

import (
	"proyecto/internal/service"
)

type Server struct {
	Citas           *service.CitaService
	PuntosEncuentro *service.PuntoEncuentroService
	Soportes        *service.SoporteService
	Auth            *service.AuthService
}

type Deps struct {
	Citas           *service.CitaService
	PuntosEncuentro *service.PuntoEncuentroService
	Soportes        *service.SoporteService
	Auth            *service.AuthService
}

func NewServer(d Deps) *Server {
	return &Server{
		Citas:           d.Citas,
		PuntosEncuentro: d.PuntosEncuentro,
		Soportes:        d.Soportes,
		Auth:            d.Auth,
	}
}

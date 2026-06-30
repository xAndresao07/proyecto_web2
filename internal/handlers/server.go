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

func NewServer(citas *service.CitaService, puntos *service.PuntoEncuentroService, soportes *service.SoporteService, auth *service.AuthService) *Server {
	return &Server{
		Citas:           citas,
		PuntosEncuentro: puntos,
		Soportes:        soportes,
		Auth:            auth,
	}
}

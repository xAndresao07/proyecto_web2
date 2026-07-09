package handlers

import (
	"proyecto/internal/service"
)

type Server struct {
	Solicitantes *service.SolicitanteService
	Dispositivos *service.DispositivoService
	Tickets      *service.TicketAyudaService
	Auth         *service.AuthService
}

type Deps struct {
	Solicitantes *service.SolicitanteService
	Dispositivos *service.DispositivoService
	Tickets      *service.TicketAyudaService
	Auth         *service.AuthService
}

func NewServer(d Deps) *Server {
	return &Server{
		Solicitantes: d.Solicitantes,
		Dispositivos: d.Dispositivos,
		Tickets:      d.Tickets,
		Auth:         d.Auth,
	}
}

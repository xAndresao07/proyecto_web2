package handlers

import (
	"solicitantesYHardware/internal/service"
)

type Server struct {
	Solicitantes *service.SolicitanteService
	Dispositivos *service.DispositivoService
	TicketAyudas *service.TicketAyudaService
	Auth         *service.AuthService
}

func NuevoServer(
	solicitantes *service.SolicitanteService,
	dispositivos *service.DispositivoService,
	ticketAyudas *service.TicketAyudaService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Solicitantes: solicitantes,
		Dispositivos: dispositivos,
		TicketAyudas: ticketAyudas,
		Auth:         auth,
	}
}

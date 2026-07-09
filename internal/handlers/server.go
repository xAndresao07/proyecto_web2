package handlers

import (
	"proyecto/internal/service"
)

type Server struct {
	Solicitantes *service.SolicitanteService
	Dispositivos *service.DispositivoService
	Tickets      *service.TicketAyudaService
	Auth         *service.AuthService
	Tecnicos     *service.TecnicoService
	Citas        *service.CitaService
	PuntosEncuentro *service.PuntoEncuentroService
	Soportes     *service.SoporteService
}

type Deps struct {
	Solicitantes *service.SolicitanteService
	Dispositivos *service.DispositivoService
	Tickets      *service.TicketAyudaService
	Auth         *service.AuthService
	Tecnicos     *service.TecnicoService
	Citas        *service.CitaService
	PuntosEncuentro *service.PuntoEncuentroService
	Soportes     *service.SoporteService
}

func NewServer(d Deps) *Server {
	return &Server{
		Solicitantes: d.Solicitantes,
		Dispositivos: d.Dispositivos,
		Tickets:      d.Tickets,
		Auth:         d.Auth,
		Tecnicos:     d.Tecnicos,
		Citas:        d.Citas,
		PuntosEncuentro: d.PuntosEncuentro,
		Soportes:     d.Soportes,
	}
}

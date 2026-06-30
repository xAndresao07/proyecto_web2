package handlers

import "proyecto/internal/service"

// Server agrupa los servicios de los que dependen los handlers.
type Server struct {
	Tecnicos *service.TecnicoService
}

func NewServer(tecnicos *service.TecnicoService) *Server {
	return &Server{Tecnicos: tecnicos}
}

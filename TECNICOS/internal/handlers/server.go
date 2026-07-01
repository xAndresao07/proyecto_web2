package handlers

import "proyecto/internal/service"

type Server struct {
	Tecnicos *service.TecnicoService
	Auth     *service.AuthService
}

func NewServer(tecnicos *service.TecnicoService, auth *service.AuthService) *Server {
	return &Server{
		Tecnicos: tecnicos,
		Auth:     auth,
	}
}

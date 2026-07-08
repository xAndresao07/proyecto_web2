package handlers

import (
	"proyecto/internal/service"
)

type Server struct {
	Tecnicos *service.TecnicoService
	Auth     *service.AuthService
}

type Deps struct {
	Tecnicos *service.TecnicoService
	Auth     *service.AuthService
}

func NewServer(d Deps) *Server {
	return &Server{
		Tecnicos: d.Tecnicos,
		Auth:     d.Auth,
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
)

type credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}

func (s *Server) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	usuario, err := s.Auth.Registrar(creds.Email, creds.Password, creds.Rol)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, usuario)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

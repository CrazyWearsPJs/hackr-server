package handlers

import (
	"net/http"

	"github.com/CrazyWearsPJs/hackr/repo"
)

type HackrHandlers struct {
	Users *repo.UserRepo
}

func (h HackrHandlers) SetupMux() *http.ServeMux {
	server_mux := http.NewServeMux()
	server_mux.HandleFunc("/api/v1/user/login", h.LoginHandler)
	server_mux.HandleFunc("/api/v1/user/register", h.RegisterHandler)
	server_mux.HandleFunc("/api/v1/user/exercises", h.SubmissionHandler)
	return server_mux
}

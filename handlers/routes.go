package handlers

import (
	"fmt"
	"net/http"

	"github.com/CrazyWearsPJs/hackr/repo"
)

type HackrMux struct {
	Users *repo.UserRepo
}

func (h HackrMux) SetupMux() *http.ServeMux {
	server_mux := http.NewServeMux()
	server_mux.HandleFunc("/", h.IndexHandler)
	server_mux.HandleFunc("/api/v1/user/login", h.LoginHandler)
	server_mux.HandleFunc("/api/v1/user/register", h.RegisterHandler)
	server_mux.HandleFunc("/api/v1/user/exercises", h.SubmissionHandler)
	return server_mux
}

func (h HackrMux) IndexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello There")
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CrazyWearsPJs/hackr/models/user"
)

type loginRequest struct {
	Email             string `json::"email"`
	PlaintextPassword string `json::"password"`
}

func (h HackrHandlers) LoginHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		h.PostLoginHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h HackrHandlers) PostLoginHandler(res http.ResponseWriter, req *http.Request) {
	var lreq *loginRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&lreq); err != nil {
		log.Printf("Error parsing login request: %v\n", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	email := lreq.Email
	pass := lreq.PlaintextPassword

	_, err := user.New(email, pass)
	if err != nil {
		log.Printf("Error creating new user object: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	//err = r.Users.Login(u)
	res.WriteHeader(http.StatusOK)
}

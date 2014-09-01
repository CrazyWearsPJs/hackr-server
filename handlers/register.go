package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/CrazyWearsPJs/hackr/models/user"
)

type registerRequest struct {
	Email             string `json::"email"`
	PlaintextPassword string `json::"password"`
}

func (h HackrHandlers) RegisterHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		h.PostRegisterHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h HackrHandlers) PostRegisterHandler(res http.ResponseWriter, req *http.Request) {
	var rreq *registerRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&rreq); err != nil {
		log.Printf("Error parsing registration request: %v\n", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	email := rreq.Email

	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("Error parsing email address: %v\n", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if !strings.HasSuffix(email, "ucr.edu") {
		log.Println("Error must be a ucr email address")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	pass := rreq.PlaintextPassword
	u, err := user.New(email, pass)
	if err != nil {
		log.Printf("Error creating new user object: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.Users.Add(u)
	if err != nil {
		log.Printf("Error, user already exists: %v", err)
		res.WriteHeader(http.StatusForbidden)
		return
	}

	res.WriteHeader(http.StatusCreated)
}

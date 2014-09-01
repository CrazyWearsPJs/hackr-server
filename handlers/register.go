package handlers

import (
	_ "bytes"
	_ "compress/zlib"
	"encoding/json"
	_ "fmt"
	_ "io"
	"log"
	"net/http"

	"github.com/CrazyWearsPJs/hackr/models/user"
	_ "net/url"

	_ "github.com/codegangsta/negroni"
	_ "github.com/garyburd/redigo/redis"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

type registerRequest struct {
	Email             string `json::"email"`
	PlaintextPassword string `json::"password"`
}

func (h HackrMux) RegisterHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		h.PostRegisterHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h HackrMux) PostRegisterHandler(res http.ResponseWriter, req *http.Request) {
	var rreq *registerRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&rreq); err != nil {
		log.Printf("Error parsing registration request: %v\n", err)
		res.WriteHeader(http.StatusBadRequest)
	}

	email := rreq.Email
	pass := rreq.PlaintextPassword
	u, err := user.New(email, pass)
	if err != nil {
		log.Printf("Error creating new user object: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	err = h.Users.Add(u)
	if err != nil {
		log.Printf("Error, user already exists: %v", err)
		res.WriteHeader(http.StatusForbidden)
	}

	res.WriteHeader(http.StatusCreated)
}

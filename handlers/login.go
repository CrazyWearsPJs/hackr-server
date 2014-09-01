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

type loginRequest struct {
	Email             string `json::"email"`
	PlaintextPassword string `json::"password"`
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		PostLoginHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func PostLoginHandler(res http.ResponseWriter, req *http.Request) {
	var lreq *loginRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&lreq); err != nil {
		log.Printf("Error parsing login request: %v\n", err)
		res.WriteHeader(http.StatusBadRequest)
	}

	email := lreq.Email
	pass := lreq.PlaintextPassword

	_, err := user.New(email, pass)
	if err != nil {
		log.Printf("Error creating new user object: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	//err = Users.Login(u)
	res.WriteHeader(http.StatusOK)
}

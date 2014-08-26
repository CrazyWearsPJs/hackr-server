package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/codegangsta/negroni"
	_ "github.com/garyburd/redigo/redis"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	_ "net/url"
	_ "os"
)

func SetupMux() *http.ServeMux {
	server_mux := http.NewServeMux()
	server_mux.HandleFunc("/", IndexHandler)
	server_mux.HandleFunc("/login", LoginHandler)
	server_mux.HandleFunc("/api/v1/user/exercises", SubmissionHandler)
	return server_mux
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello There")
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hi")
}

type submitResponse struct {
	Status string `json::"status"`
	Error  string `json::"error"`
}

type submitRequest struct {
	Email string `json::"email"`
	Key   string `json::"key"`
	Code  string `json::"code"`
}

func SubmissionHandler(res http.ResponseWriter, req *http.Request) {
	var sreq *submitRequest

	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&sreq); err != nil {
			log.Printf("Error parsing response: %v", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		sres := submitResponse{Status: "OK", Error: "None"}

		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		res.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(res)
		if err := encoder.Encode(sres); err != nil {
			log.Printf("Error parsing response: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	default:
		res.WriteHeader(http.StatusBadRequest)
	}
}

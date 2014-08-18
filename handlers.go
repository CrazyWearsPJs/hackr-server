package main

import (
	 "fmt"
	_ "github.com/codegangsta/negroni"
	_ "github.com/garyburd/redigo/redis"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"net/http"
	_ "net/url"
	_ "os"
)

func SetupMux() *http.ServeMux {
	server_mux := http.NewServeMux()
	server_mux.HandleFunc("/", IndexHandler)
    server_mux.HandleFunc("/login", LoginHandler)
	return server_mux
}

func IndexHandler(res http.ResponseWriter, req *http.Request){
    fmt.Fprintf(res, "Hello There")
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hi")
}

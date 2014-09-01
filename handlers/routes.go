package handlers

import (
	_ "bytes"
	_ "compress/zlib"
	_ "encoding/json"
	"fmt"
	_ "io"
	_ "log"
	"net/http"

	_ "github.com/CrazyWearsPJs/hackr/models/user"
	"github.com/CrazyWearsPJs/hackr/repo"
	_ "net/url"

	_ "github.com/codegangsta/negroni"
	_ "github.com/garyburd/redigo/redis"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
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

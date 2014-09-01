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

var (
	Users *repo.UserRepo
)

func SetupMux() *http.ServeMux {
	server_mux := http.NewServeMux()
	server_mux.HandleFunc("/", IndexHandler)
	server_mux.HandleFunc("/api/v1/user/login", LoginHandler)
	server_mux.HandleFunc("/api/v1/user/register", RegisterHandler)
	server_mux.HandleFunc("/api/v1/user/exercises", SubmissionHandler)
	return server_mux
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello There")
}

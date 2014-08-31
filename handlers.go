package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/url"
	_ "os"

	_ "github.com/codegangsta/negroni"
	_ "github.com/garyburd/redigo/redis"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

const (
	Kb = 1024
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
	switch req.Method {
	case "POST":
		PostSubmissionHandler(res, req)
	default:
		res.WriteHeader(http.StatusBadRequest)
	}
}

func PostSubmissionHandler(res http.ResponseWriter, req *http.Request) {
	var sreq *submitRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&sreq); err != nil {
		log.Printf("Error parsing response: %v", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := Users.FindUserByEmail(sreq.Email)
	if err != nil || sreq.Key != u.APIKey {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	code := sreq.Code

	compressed_code, err := compressCode(code)
	if err != nil {
		log.Printf("Compressing code failure: %v\n", err)
		res.WriteHeader(http.StatusInternalServerError)
	}
	uncompress_code, err := uncompressCode(compressed_code)
	if err != nil {
		log.Printf("Uncompressing code failure: %v\n", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("Compressed Code: %v\n", compressed_code)
	fmt.Printf("Uncompressed Code: %v\n", uncompress_code)

	if len(compressed_code) > 20*Kb {
		log.Println("File is too big!")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = Users.SubmitCode(u.Email, compressed_code)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sres := submitResponse{Status: "OK", Error: "None"}

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(res)
	if err := encoder.Encode(sres); err != nil {
		log.Printf("Error parsing response: %v", err)
	}
}

func compressCode(code string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	defer gz.Close()

	_, err := gz.Write([]byte(code))
	if err != nil {
		return b.String(), err
	}

	return b.String(), nil
}

func uncompressCode(code string) (string, error) {
	b := bytes.NewBufferString(code)
	var uncompresed []byte
	gunz, err := gzip.NewReader(b)
	if err != nil {
		return "", err
	}
	defer gunz.Close()

	if _, err := gunz.Read(uncompresed); err != nil {
		return string(uncompresed), err
	}

	return string(uncompresed), nil
}

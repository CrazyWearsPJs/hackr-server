package handlers

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	_ "fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/CrazyWearsPJs/hackr/models/user"
	_ "net/url"

	_ "github.com/codegangsta/negroni"
	_ "github.com/garyburd/redigo/redis"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

const (
	Kb = 1024
)

type submitResponse struct {
	Status string `json::"status"`
	Error  string `json::"error"`
}

type submitRequest struct {
	Email string `json::"email"`
	Key   string `json::"key"`
	Code  string `json::"code"`
}

func (h *HackrMux) SubmissionHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		h.PostSubmissionHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *HackrMux) PostSubmissionHandler(res http.ResponseWriter, req *http.Request) {
	var sreq *submitRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&sreq); err != nil {
		log.Printf("Error parsing submission request: %v\n", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.Users.FindUserByEmail(sreq.Email)
	if err != nil || sreq.Key != u.APIKey {
		res.WriteHeader(http.StatusForbidden)
	}

	code := sreq.Code

	compressed_code, err := compressCode(code)
	if err != nil {
		log.Printf("Compressing code failure: %v\n", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	if len(compressed_code) > 20*Kb {
		log.Println("File is too big!")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.Users.SubmitCode(u.Email, compressed_code)
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
	compress := zlib.NewWriter(&b)

	_, err := compress.Write([]byte(code))
	if err != nil {
		return b.String(), err
	}

	compress.Close()

	return b.String(), nil
}

func uncompressCode(code string) (string, error) {
	var b bytes.Buffer
	code_buffer := bytes.NewBufferString(code)

	uncompress, err := zlib.NewReader(code_buffer)
	if err != nil {
		return "", err
	}

	io.Copy(&b, uncompress)
	uncompress.Close()

	return b.String(), nil
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CrazyWearsPJs/hackr/util"
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

func (h HackrHandlers) SubmissionHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		h.PostSubmissionHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h HackrHandlers) PostSubmissionHandler(res http.ResponseWriter, req *http.Request) {
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
		return
	}

	code := sreq.Code

	compressed_code, err := util.CompressString(code)
	if err != nil {
		log.Printf("Compressing code failure: %v\n", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(compressed_code) > 20*Kb {
		log.Println("File is too big!")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.Users.SubmitCode(sreq.Email, compressed_code)
	if err != nil {
		log.Printf("Error: Unable to submit exercise: %v!\n", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sres := submitResponse{Status: "OK", Error: "None"}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(res)
	if err := encoder.Encode(sres); err != nil {
		log.Printf("Error parsing response: %v", err)
	}
}

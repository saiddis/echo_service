package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error reading request body: %v", err))
		return
	}

	log.Printf("request: %s", reqBody)

	w.WriteHeader(200)
	w.Write(reqBody)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

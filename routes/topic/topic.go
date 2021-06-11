package topic

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Switch Methods
	switch r.Method {
	case "GET":
		w.Write([]byte(http.StatusText(http.StatusOK)))
	case "POST":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(http.StatusText(http.StatusAccepted)))
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(http.StatusText(http.StatusNoContent)))
	default:
		log.Println("Method not supported!", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

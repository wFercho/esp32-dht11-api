package api

import (
	"log"
	"net/http"

	"dht11_server/db"
)

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {

	sMux := http.NewServeMux()

	sMux.HandleFunc("GET /api/data", makeHTTPHandleFunc(s.handleGetRegisterDHT11Data))
	sMux.HandleFunc("POST /api/data", makeHTTPHandleFunc(s.handleRegisterDHT11Data))

	log.Println("ESP32 DHT11 API running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, sMux)
}

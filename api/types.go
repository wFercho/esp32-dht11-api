package api

import (
	"dht11_server/db"
	"net/http"
)

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type RegisterDHT11Request struct {
	Temperature float32 `json:"temperatura"`
	Humidity    float32 `json:"humedad"`
	Is_UDP      bool    `json:"is_UDP"`
}

type APIServer struct {
	listenAddr string
	store      db.Storage
}

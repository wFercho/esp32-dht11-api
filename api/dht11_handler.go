package api

import (
	"encoding/json"
	"log"
	"net/http"

	e "dht11_server/entities"
)

func (s *APIServer) handleRegisterDHT11Data(w http.ResponseWriter, r *http.Request) error {
	registerDHT11Req := new(RegisterDHT11Request)
	if err := json.NewDecoder(r.Body).Decode(registerDHT11Req); err != nil {
		return err
	}

	defer r.Body.Close()

	dht11Data := e.NewDHT11(registerDHT11Req.Temperature, registerDHT11Req.Humidity, registerDHT11Req.Is_UDP)

	if err := s.store.RegisterDHT11Data(dht11Data); err != nil {
		log.Fatal(err)
		return err
	}
	w.WriteHeader(http.StatusCreated)

	return nil
}

func (s *APIServer) handleGetRegisterDHT11Data(w http.ResponseWriter, r *http.Request) error {

	defer r.Body.Close()

	data, err := s.store.GetDHT11Registers()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, data)
}

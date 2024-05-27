package db

import (
	e "dht11_server/entities"
)

type Storage interface {
	RegisterDHT11Data(*e.DHT11) error
	GetDHT11Registers() (*[]e.DHT11, error)
}

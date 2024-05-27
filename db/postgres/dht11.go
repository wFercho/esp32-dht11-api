package postgres

import (
	e "dht11_server/entities"
	"fmt"
)

func (store *PostgresStore) RegisterDHT11Data(dht11 *e.DHT11) error {
	query := fmt.Sprintf(`INSERT INTO %s (temperatura, humedad, timestamp, is_udp)
	VALUES ($1, $2, $3, $4)
	RETURNING id`, DHT11_TABLE)

	id := 0

	err := store.db.QueryRow(
		query,
		dht11.Temperature,
		dht11.Humidity,
		dht11.RegisterAt,
		dht11.Is_UDP).Scan(&id)

	if err != nil {
		return err
	}

	dht11.ID = id

	return nil
}

func (store *PostgresStore) GetDHT11Registers() (*[]e.DHT11, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, DHT11_TABLE)

	rows, err := store.db.Query(query)

	if err != nil {
		return nil, err
	}

	dht11_registers := make([]e.DHT11, 0)

	for rows.Next() {
		dht11 := e.DHT11{}
		err := rows.Scan(&dht11.ID, &dht11.Temperature, &dht11.Humidity, &dht11.RegisterAt, &dht11.Is_UDP)
		if err != nil {
			return nil, err
		}
		dht11_registers = append(dht11_registers, dht11)
	}

	return &dht11_registers, nil
}

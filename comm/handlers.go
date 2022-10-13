package comm

import (
	"encoding/json"
	"fmt"

	"github.com/joelschutz/gorally/models"
)

func HandleMessage(message []byte, db *DB) (msg []byte, err error) {
	data := Payload{}
	err = json.Unmarshal(message, &data)
	if err != nil {
		return []byte("error"), fmt.Errorf("Failed to Unmarshall message: %s", err)
	}
	switch data.Action.Target {
	case "vehicle":
		return handleVehicle(data, db)
	case "driver":
		return handleDriver(data, db)
	}
	return []byte{}, nil
}

func handleVehicle(payload Payload, db *DB) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Vehicle{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.Vehicles = append(db.Vehicles, v)
	case "list":
		return json.Marshal(db.Vehicles)
	}
	return []byte("ok"), nil
}

func handleDriver(payload Payload, db *DB) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Driver{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall driver: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.Drivers = append(db.Drivers, v)
	case "list":
		d, err := json.Marshal(db.Drivers)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall drivers: %s", err)
		}
		v := Payload{
			Action: Action{
				Target: "driver",
				Method: "list",
			},
			Data: d,
		}
		p, err := json.Marshal(v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall drivers: %s", err)
		}
		return p, nil
	}
	return []byte("ok"), nil
}

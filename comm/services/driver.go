package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/models"
)

type DriverService struct {
}

func NewDriverService() *DriverService {
	return &DriverService{}
}

func (g *DriverService) HandlePayload(ctx context.Context, payload schema.Payload, db Storage) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Driver{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.AddDriver(ctx, v)
	case "update":
		v := models.Driver{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.UpdateDriver(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := db.ListDrivers(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Drivers: %s", err)
		}
		v := schema.Payload{
			Action: schema.Action{
				Target: "vehicle",
				Method: "list",
			},
			Data: d,
		}
		p, err := json.Marshal(v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Drivers: %s", err)
		}
		return p, nil
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

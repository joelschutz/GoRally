package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/models"
)

type ServiceHandler interface {
	HandlePayload(ctx context.Context, payload schema.Payload, db Storage) ([]byte, error)
}

type VehicleService struct {
}

func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

func (g *VehicleService) HandlePayload(ctx context.Context, payload schema.Payload, db Storage) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Vehicle{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.AddVehicle(ctx, v)
	case "update":
		v := models.Vehicle{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.UpdateVehicle(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := db.ListVehicles(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Vehicles: %s", err)
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
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Vehicles: %s", err)
		}
		return p, nil
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

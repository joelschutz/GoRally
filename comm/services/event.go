package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/models"
)

type EventService struct {
}

func NewEventService() *EventService {
	return &EventService{}
}

func (g *EventService) HandlePayload(ctx context.Context, payload schema.Payload, db Storage) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Event{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.AddEvent(ctx, v)
	case "update":
		v := models.Event{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.UpdateEvent(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := db.ListEvents(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Events: %s", err)
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
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Events: %s", err)
		}
		return p, nil
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

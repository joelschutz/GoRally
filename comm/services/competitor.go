package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/comm/storage"
)

type CompetitorService struct {
}

func NewCompetitorService() *CompetitorService {
	return &CompetitorService{}
}

func (g *CompetitorService) HandlePayload(ctx context.Context, payload schema.Payload, db storage.Storage) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := schema.CompetitorSchema{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Competitors: %s", err)
		}
		fmt.Println("vehicle: ", v)
		return msg, db.AddCompetitor(ctx, v)
	case "update":
		v := schema.CompetitorSchema{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Competitors: %s", err)
		}
		fmt.Println("vehicle: ", v)
		return msg, db.UpdateCompetitor(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := db.ListCompetitors(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Competitors: %s", err)
		}
		v := schema.Payload{
			Action: schema.Action{
				Target: "competitor",
				Method: "list",
			},
			Data: d,
		}
		p, err := json.Marshal(v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Competitors: %s", err)
		}
		return p, nil
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

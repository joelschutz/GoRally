package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/comm/storage"
	"github.com/joelschutz/gorally/models"
)

type TrackService struct {
}

func NewTrackService() *TrackService {
	return &TrackService{}
}

func (g *TrackService) HandlePayload(ctx context.Context, payload schema.Payload, db storage.Storage) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Track{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.AddTrack(ctx, v)
	case "update":
		v := models.Track{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		db.UpdateTrack(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := db.ListTracks(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Tracks: %s", err)
		}
		v := schema.Payload{
			Action: schema.Action{
				Target: "track",
				Method: "list",
			},
			Data: d,
		}
		p, err := json.Marshal(v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Tracks: %s", err)
		}
		return p, nil
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

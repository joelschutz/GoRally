package services

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/comm/storage"
	"github.com/joelschutz/gorally/mechanics"
	"github.com/joelschutz/gorally/models"
)

type EventService struct {
}

func NewEventService() *EventService {
	return &EventService{}
}

func saveToCsv(data models.TrackResult) {
	f, err := os.OpenFile("resultBySegment.csv", os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(f)
	for _, v := range data.TimeBySegment {
		err := w.Write([]string{fmt.Sprint(v)})
		if err != nil {
			panic(err)
		}
	}

	f2, err := os.OpenFile("VSBySecond.csv", os.O_CREATE, 0755)
	defer f2.Close()
	if err != nil {
		panic(err)
	}
	w = csv.NewWriter(f2)
	w.Write([]string{
		"speed",
		"acceleration",
		"damage",
		"location",
	})
	for _, v := range data.VStateBySecond {
		err := w.Write([]string{
			fmt.Sprint(v.Speed),
			fmt.Sprint(v.Acceleration),
			fmt.Sprint(v.Damage),
			fmt.Sprint(v.Location),
		})
		if err != nil {
			panic(err)
		}
	}
	f3, err := os.OpenFile("TSBySecond.csv", os.O_CREATE, 0755)
	defer f3.Close()
	if err != nil {
		panic(err)
	}
	w = csv.NewWriter(f3)
	w.Write([]string{
		"maxSpeed",
		"maxAcceleration",
		"distanceLeft",
		"location",
	})
	for _, v := range data.TStateBySecond {
		err := w.Write([]string{
			fmt.Sprint(v.MaxSpeed),
			fmt.Sprint(v.MaxTorque),
			fmt.Sprint(v.DistanceLeft),
			fmt.Sprint(v.Location),
		})
		if err != nil {
			panic(err)
		}
	}
}

func (g *EventService) HandlePayload(ctx context.Context, payload schema.Payload, db storage.Storage) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := schema.EventSchema{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Events: %s", err)
		}
		fmt.Println("vehicle: ", v)
		return msg, db.AddEvent(ctx, v)
	case "update":
		v := schema.EventSchema{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Events: %s", err)
		}
		fmt.Println("vehicle: ", v)
		return msg, db.UpdateEvent(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := db.ListEvents(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Events: %s", err)
		}
		v := schema.Payload{
			Action: schema.Action{
				Target: "event",
				Method: "list",
			},
			Data: d,
		}
		p, err := json.Marshal(v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Events: %s", err)
		}
		return p, nil
	case "run":
		e, err := db.GetEvent(ctx, payload.Action.Index)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to get Events: %s", err)
		}
		r := mechanics.CalcTrackTime(e.Tracks[0], e.Competitors[0].Driver, e.Competitors[0].Vehicle)
		saveToCsv(r)
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

package comm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/joelschutz/gorally/comm/services"
	"github.com/joelschutz/gorally/models"
)

type WebSocketHandler interface {
	HandleFunc(w http.ResponseWriter, r *http.Request)
	Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)
}

type Game struct {
	upgrader *websocket.Upgrader
	db       services.Storage
}

func NewGameHandler(up *websocket.Upgrader) *Game {
	return &Game{upgrader: up}
}

func (g *Game) HandleFunc(w http.ResponseWriter, r *http.Request) {
	c, err := g.upgrader.Upgrade(w, r, nil)
	fmt.Println("Connected")
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		log.Printf("recv: %s", message)
		message, err = g.HandleMessage(r.Context(), message)
		if err != nil {
			fmt.Println(err)
		}
		err = c.WriteMessage(mt, message)
		fmt.Printf("just sent: %s", message)
		if err != nil {
			log.Println("write:", err)
			break
		}
		fmt.Println("db: ", g.db)
	}
}

func (g *Game) HandleMessage(ctx context.Context, message []byte) (msg []byte, err error) {
	data := Payload{}
	err = json.Unmarshal(message, &data)
	if err != nil {
		return []byte("error"), fmt.Errorf("Failed to Unmarshall message: %s", err)
	}
	switch data.Action.Target {
	case "vehicle":
		return g.handleVehicle(ctx, data)
	case "driver":
		return g.handleDriver(ctx, data)
	case "track":
		return g.handleTrack(ctx, data)
	}

	return []byte("error"), fmt.Errorf("Target not Allowed")
}

func (g *Game) handleVehicle(ctx context.Context, payload Payload) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Vehicle{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		g.db.AddVehicle(ctx, v)
	case "update":
		v := models.Vehicle{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall vehicle: %s", err)
		}
		fmt.Println("vehicle: ", v)
		g.db.UpdateVehicle(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := g.db.ListVehicles(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall Vehicles: %s", err)
		}
		v := Payload{
			Action: Action{
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

func (g *Game) handleDriver(ctx context.Context, payload Payload) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Driver{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall driver: %s", err)
		}
		fmt.Println("driver: ", v)
		g.db.AddDriver(ctx, v)
	case "update":
		v := models.Driver{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall driver: %s", err)
		}
		fmt.Println("driver: ", v)
		g.db.UpdateDriver(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := g.db.ListDrivers(ctx)
		d, err := json.Marshal(arr)
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
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

func (g *Game) handleTrack(ctx context.Context, payload Payload) (msg []byte, err error) {
	switch payload.Action.Method {
	case "add":
		v := models.Track{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall track: %s", err)
		}
		fmt.Println("track: ", v)
		g.db.AddTrack(ctx, v)
	case "update":
		v := models.Track{}
		err := json.Unmarshal(payload.Data, &v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall track: %s", err)
		}
		fmt.Println("track: ", v)
		g.db.UpdateTrack(ctx, payload.Action.Index, v)
	case "list":
		arr, _ := g.db.ListTracks(ctx)
		d, err := json.Marshal(arr)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall tracks: %s", err)
		}
		v := Payload{
			Action: Action{
				Target: "track",
				Method: "list",
			},
			Data: d,
		}
		p, err := json.Marshal(v)
		if err != nil {
			return []byte("error"), fmt.Errorf("Failed to Unmarshall tracks: %s", err)
		}
		return p, nil
	}
	return []byte("error"), fmt.Errorf("Method not Allowed")
}

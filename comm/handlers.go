package comm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/joelschutz/gorally/comm/schema"
	"github.com/joelschutz/gorally/comm/services"
)

type WebSocketHandler interface {
	HandleFunc(w http.ResponseWriter, r *http.Request)
	Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)
}

type Game struct {
	upgrader   *websocket.Upgrader
	db         services.Storage
	vehicleSvc services.ServiceHandler
	driverSvc  services.ServiceHandler
	trackSvc   services.ServiceHandler
	eventSvc   services.ServiceHandler
}

func NewGameHandler(up *websocket.Upgrader, db services.Storage, vehicleSvc services.ServiceHandler, driverSvc services.ServiceHandler, trackSvc services.ServiceHandler, eventSvc services.ServiceHandler) *Game {
	return &Game{
		upgrader:   up,
		db:         db,
		vehicleSvc: vehicleSvc,
		driverSvc:  driverSvc,
		trackSvc:   trackSvc,
		eventSvc:   eventSvc,
	}
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
	data := schema.Payload{}
	err = json.Unmarshal(message, &data)
	if err != nil {
		return []byte("error"), fmt.Errorf("Failed to Unmarshall message: %s", err)
	}
	switch data.Action.Target {
	case "vehicle":
		return g.vehicleSvc.HandlePayload(ctx, data, g.db)
	case "driver":
		return g.driverSvc.HandlePayload(ctx, data, g.db)
	case "track":
		return g.trackSvc.HandlePayload(ctx, data, g.db)
	case "event":
		return g.eventSvc.HandlePayload(ctx, data, g.db)
	}

	return []byte("error"), fmt.Errorf("Target not Allowed")
}

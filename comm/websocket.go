// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comm

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func NewUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
}

type Echo struct {
	upgrader *websocket.Upgrader
}

func NewEchoHandler(up *websocket.Upgrader) *Echo {
	return &Echo{upgrader: up}
}

func (e *Echo) HandleFunc(w http.ResponseWriter, r *http.Request) {
	c, err := e.upgrader.Upgrade(w, r, nil)
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
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

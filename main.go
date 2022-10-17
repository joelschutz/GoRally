package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/joelschutz/gorally/comm"
	"github.com/joelschutz/gorally/comm/services"
	"github.com/joelschutz/gorally/comm/storage"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	// Create DB
	// db := storage.NewMemoryDB()
	db, err := storage.NewSqliteDB("sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	// Create Upgrader
	up := comm.NewUpgrader()
	// Create Services
	vSvc := services.NewVehicleService()
	dSvc := services.NewDriverService()
	tSvc := services.NewTrackService()
	eSvc := services.NewEventService()
	// Create Handlers
	e := comm.NewEchoHandler(up)
	g := comm.NewGameHandler(up, db, vSvc, dSvc, tSvc, eSvc)

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", e.HandleFunc)
	http.HandleFunc("/game", g.HandleFunc)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

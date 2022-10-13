package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/joelschutz/gorally/comm"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", comm.Echo)
	http.HandleFunc("/game", comm.Game)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

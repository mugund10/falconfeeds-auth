package main

import (
	"log"

	"github.com/mugund10/falconfeeds-auth/api"
)

func main() {
	listenAddress := ":8080"
	server := api.Newserver(listenAddress)
	log.Printf("server running on http://0.0.0.0%s", listenAddress)
	log.Fatal(server.Start())
}

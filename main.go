package main

import (
	"flag"
	"go-crud/server"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	routes := server.SetupRoutes()
	srv := server.SetupServer(routes)
	server.StartServer(srv)
	server.StopServer(wait, srv)
}

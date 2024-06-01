package main

import (
	"rest_app_metrics/processing"
	"rest_app_metrics/rest"
)

func main() {
	server := rest.CreateRestServer()

	rest.AddHealthEndpoints(server.RouterGroups)
	processing.AddEndpoints(server.RouterGroups)

	server.Start()
}

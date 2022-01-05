package main

import (
	"go-mongodb/delivery"
)

func main() {
	var server delivery.Routes
	server.StartGin()
}

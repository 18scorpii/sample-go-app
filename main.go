package main

import (
	"fmt"

	"github.com/18scorpii/sample-go-app/server"
)

func main() {
	fmt.Println("Starting Sample Go App")
	server.StartHttpServer()
}

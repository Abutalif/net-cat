package main

import (
	"fmt"

	"net-cat/config"
	"net-cat/usecases"
)

func main() {
	port, isServer, err := config.IsServerMode()
	if err != nil {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if isServer {
		fmt.Printf("Listening on the port %v\n", port)
		err = usecases.NewServer("localhost" + port).StartServer()
	} else {
		err = usecases.NewClient("localhost" + port).Connect()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}

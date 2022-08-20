package main

import (
	"log"
	"os"

	"net-cat/usecases"
)

func main() {
	isServer, err := IsServerMode()
	if err != nil {
		log.Println(err)
		return
	}
	if isServer {
		err = usecases.NewServer("localhost:8080", "testfile.txt").StartServer()
		log.Println("Listening on the port :8080")
	} else {
		// err = ClientMode()
	}
	if err != nil {
		log.Println(err)
		return
	}
}

// Returns True for Server mode, False for Client mode
func IsServerMode() (bool, error) {
	// hanlde flags
	args := os.Args
	if len(args) == 0 {
		return false, nil
	}

	return true, nil
}

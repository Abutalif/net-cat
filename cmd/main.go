package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"net-cat/internal/models"
)

func main() {
	isServer, err := IsServerMode()
	if err != nil {
		log.Println(err)
		return
	}
	if isServer {
		err = ServerMode()
	} else {
		err = ClientMode()
	}
	if err != nil {
		log.Println(err)
		return
	}
}

// should be in config
// Returns True for Server mode, False for Client mode
func IsServerMode() (bool, error) {
	// hanlde flags
	//-l listen
	args := os.Args
	if len(args) == 0 {
		return false, nil // error should not be nil
	}

	return true, nil
}

// should be in internal/delivery server
// Runs a Server
func ServerMode() error {
	lstn, err := net.Listen("tcp", "localhost:8080") // default or changable host
	if err != nil {
		return err
	}
	defer lstn.Close()
	welcome, err := os.ReadFile("./static/welcome.txt")
	if err != nil {
		return err
	}

	connMap := &sync.Map{}
	// map[User]net.Conn
	for {
		conn, err := lstn.Accept()
		if err != nil {
			break
		}
		_, err = conn.Write(welcome)
		if err != nil {
			break
		}
		newUser := models.User{}
		connChan := make(chan interface{})
		go HandleNewConn(conn, connChan)
		connMap.Store(newUser, conn)
	}
	// should return nil if correct, else err
	return err
}

//
func HandleNewConn(conn net.Conn, connChan chan interface{}) {
	defer conn.Close()

	conn.Write([]byte("[ENTER YOUR NAME]: ")) // do it in cycle
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// send signal two chanel he has exited
			return
		}
		if true {
			continue
		} else {
			// send messaeg
			fmt.Println("send user input two chanel")
		}

	}
}

// should be in internal/delivery/client
// Runs a Client
func ClientMode() error {
	// connect to server with a given IP
	// create a file with a penguin
	return nil
}

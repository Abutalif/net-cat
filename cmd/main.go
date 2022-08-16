package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	isServer, err := CheckMode()
	if err != nil {
		// TODO: log error
		return
	}
	if isServer {
		err = ServerMode()
	} else {
		err = ClientMode()
	}
	if err != nil {
		// TODO: log error
		return
	}
}

// Returns True for Server mode, False for Client mode
func CheckMode() (bool, error) {
	// TODO:
	// hanlde flags
	//-l listen
	// also
	args := os.Args
	if len(args) == 0 {
		// TODO: log error
		return false, nil // TODO: error should not be nil
	}

	return true, nil
}

// Runs a Server
func ServerMode() error {
	lstn, err := net.Listen("tcp", "localhost:8080") // TODO: default or changable host
	// errChan := make(chan error)
	// msgChan := make(chan string)
	if err != nil {
		// TODO: log error
		return err
	}
	defer lstn.Close()

	// TODO:
	// connMap := &sync.Map{}
	users := make(map[int]net.Conn)
	i := 0
	for {
		conn, err := lstn.Accept()
		if err != nil {
			break
		}
		users[i] = conn
		i++

		go HandleNewUser(conn)
	}
	// should return nil if correct, else err
	return err
}

func HandleNewUser(conn net.Conn) {
	defer conn.Close()
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(userInput)

	}
}

// Runs a Client
func ClientMode() error {
	// TODO: connect to server with a given IP
	// TODO: create a file with a penguin
	return nil
}

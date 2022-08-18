package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"sync"
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

// Returns True for Server mode, False for Client mode
func IsServerMode() (bool, error) {
	// hanlde flags
	args := os.Args
	if len(args) == 0 {
		return false, nil
	}

	return true, nil
}

// write down all msg

// Runs a Server
func ServerMode() error {
	lstn, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	defer lstn.Close()
	welcome, err := os.ReadFile("./static/welcome.txt")
	if err != nil {
		return err
	}

	userMap := &sync.Map{}
	// map[models.User]net.Conn
	for {
		conn, err := lstn.Accept()
		if err != nil {
			break
		}
		_, err = conn.Write(welcome)
		if err != nil {
			break
		}

		go HandleNewConn(conn, *userMap)
	}
	return err // edit
}

//
func HandleNewConn(conn net.Conn, users sync.Map) {
	defer conn.Close()

	// create a user
	// add user to sync map

	// get name cycle
	// cycle
	name := "me"
	// send prev_msg if len(prev_msg)!=nil
	users.Store(conn, name)
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		msg, err := AddTimeStamp(userInput)
		if err != nil {
			return
		}

		if err = WriteToChat(msg); err != nil {
			return
		}

		if err = SaveMessage(msg); err != nil {
			return
		}
	}
}

func AddTimeStamp(rawMsg string) (string, error) {
	return "a", nil
}

func WriteToChat(msg string) error {
	return nil
}

func SaveMessage(msg string) error {
	return nil
}

// Runs a Client
func ClientMode() error {
	return nil
}

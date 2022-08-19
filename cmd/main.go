package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
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
		fmt.Println(userMap)
		go HandleNewConn(conn, userMap)
	}
	return err // edit
}

type client struct {
	Name     string
	IsActive bool
}

func HandleNewConn(conn net.Conn, users *sync.Map) {
	defer conn.Close()
	var name string
	var err error

	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Write([]byte("Cannot Read name\n"))
			continue
		}

		_, hasUser := users.LoadOrStore(name, conn)
		if hasUser {
			conn.Write([]byte("There is already a user with such name in the chat\n"))
		} else {
			break
		}
	}

	user := &client{
		Name:     name,
		IsActive: true,
	}
	// send prev_ms	g if len(prev_msg)!=nil
	users.Store(name, conn)
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		msg := AddTimeStamp(userInput, user.Name)
		if err != nil {
			return
		}
		users.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				if _, err := conn.Write([]byte(msg)); err != nil {
					return false
				}
			}
			return true
		})

		if err = SaveMessage(msg); err != nil {
			return
		}
	}
}

func AddTimeStamp(rawMsg string, name string) string {
	timeStamp := time.Now().Format("2020-01-20 15:48:41")
	return "[" + timeStamp + "][" + strings.Replace(name, "\n", "", -1) + "]" + rawMsg
}

func SaveMessage(msg string) error {
	return nil
}

// Runs a Client
func ClientMode() error {
	return nil
}

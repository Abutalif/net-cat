package usecases

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type server struct {
	address string
	users   *sync.Map
	logfile string
}

type Connecter interface {
	StartServer() error
}

func NewServer(address, logfile string) Connecter {
	return &server{
		address: address,
		users:   &sync.Map{},
		logfile: logfile,
	}
}

func (s *server) StartServer() error {
	numOfUsers := 0
	lstn, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer lstn.Close()
	welcome, err := os.ReadFile("./static/welcome.txt")
	if err != nil {
		return err
	}
	// TODO: create message log file that appends
	for {
		conn, err := lstn.Accept()
		if err != nil {
			return err
		}
		_, err = conn.Write(welcome)
		if err != nil {
			return err
		}
		if numOfUsers <= 10 {
			numOfUsers++
			go s.handleNewConn(conn)
		} else {
			conn.Write([]byte("Chat capacity full.\nGood bye!"))
			conn.Close()
		}

	}

	// somehere there should be a channel for sending all prev messages
}

func (s *server) handleNewConn(conn net.Conn) {
	var goodName string
	var err error
	var name string
	// var messageHist []byte
	defer s.users.Delete(goodName)

	// Reading username
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Write([]byte("Cannot Read name\n"))
			continue
		}
		goodName = s.formatName(name) // check not empty
		if _, hasUser := s.users.LoadOrStore(goodName, conn); hasUser {
			conn.Write([]byte("There is already a user with such name in the chat\n"))
		} else {
			break
		}
	}

	// Sending old messages
	// oldMessages, err := os.ReadFile(s.logfile) // check availavility of the file
	// if err != nil {
	// 	log.Println(err)
	// }
	// conn.Write(oldMessages)

	// Sending enterence notification
	s.writeToChat(goodName + " has joined our chat...\n")
	s.users.Store(goodName, conn)

	// writing the messages
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			s.writeToChat(goodName + " has left our chat...\n")
			return
		}
		msg := s.addTimeStamp(userInput, name)
		s.writeToChat(msg)

		if err = s.saveMessage(msg); err != nil {
			log.Println(err)
		}
	}
}

// adding the timestamp to the message
func (s *server) addTimeStamp(rawMsg string, name string) string {
	// use fmt.sprintf
	timeStamp := time.Now().Format("2020-01-20 15:48:41") // add proper time format
	return "[" + timeStamp + "][" + strings.Replace(name, "\n", "", -1) + "]" + rawMsg
}

// formating the name to look good
func (s *server) formatName(name string) string {
	return strings.Replace(name, "\n", "", -1)
}

// saving message to the chat
func (s *server) saveMessage(msg string) error {
	if err := os.WriteFile(s.logfile, []byte(msg), 0o644); err != nil {
		return err
	}
	// close it
	// append it
	return nil
}

// writing to the chat
func (s *server) writeToChat(msg string) {
	// send all, but myself
	s.users.Range(func(key, value interface{}) bool {
		if conn, ok := value.(net.Conn); ok {
			if _, err := conn.Write([]byte(msg)); err != nil {
				return false
			}
		}
		return true
	})
}

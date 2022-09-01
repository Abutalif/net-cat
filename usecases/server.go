package usecases

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type server struct {
	syncMaster  *sync.Mutex
	users       map[string]net.Conn
	address     string
	messageHist []string
	// lastMessage chan string
}

// maybe add broadcaster

type Connecter interface {
	StartServer() error
}

func NewServer(address string) Connecter {
	return &server{
		syncMaster:  &sync.Mutex{},
		users:       make(map[string]net.Conn),
		address:     address,
		messageHist: make([]string, 0),
		// lastMessage: make(chan string),
	}
}

func (s *server) StartServer() error {
	lstn, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer lstn.Close()
	welcome, err := os.ReadFile("./static/welcome.txt")
	if err != nil {
		return err
	}

	for {
		conn, err := lstn.Accept()
		if err != nil {
			return err
		}
		_, err = conn.Write(welcome)
		if err != nil {
			return err
		}
		s.syncMaster.Lock()
		if len(s.users) < 10 {
			go s.handleNewConn(conn)
		} else {
			conn.Write([]byte("Chat capacity full.\nGood bye!"))
			conn.Close()
		}
		fmt.Println(len(s.users))
		s.syncMaster.Unlock()

	}
}

func (s *server) handleNewConn(conn net.Conn) {
	var goodName string
	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, err := reader.ReadString('\n')
		if err != nil {
			conn.Write([]byte("Cannot Read name\n"))
			continue
		}
		goodName = s.formatName(name)
		if goodName == "" {
			conn.Write([]byte("Empty name\n"))
			continue
		}

		if s.hasUser(goodName) {
			conn.Write([]byte("There is already a user with such name in the chat\n"))
			continue
		} else {
			s.addUser(goodName, conn)
			break
		}

	}

	defer s.removeUser(goodName)

	// Sending old messages
	s.sendOldMessages(conn)

	// Sending enterence notification
	s.writeToChat(goodName, goodName+" has joined our chat...\n", true)

	// writing the messages
	// conn.Write([]byte(s.addTimeStamp("", goodName)))
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			s.writeToChat(goodName, goodName+" has left our chat...\n", true)
			s.removeUser(goodName)
			return
		}

		if userInput == "\n" {
			continue
		}
		s.writeToChat(goodName, userInput, false)
	}
}

// checks if there is a user with such name
func (s *server) hasUser(name string) bool {
	s.syncMaster.Lock()
	defer s.syncMaster.Unlock()
	for key := range s.users {
		if key == name {
			return true
		}
	}
	return false
}

func (s *server) addUser(name string, conn net.Conn) {
	s.syncMaster.Lock()
	defer s.syncMaster.Unlock()
	s.users[name] = conn
}

// removes user from the chat
func (s *server) removeUser(name string) {
	s.syncMaster.Lock()
	defer s.syncMaster.Unlock()
	s.users[name].Close()
	delete(s.users, name)
}

// adding the timestamp to the message
func (s *server) addTimeStamp(name string) string {
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%v][%v]: ", timeStamp, name)
}

// send old messages
func (s *server) sendOldMessages(conn net.Conn) {
	res := ""
	for _, line := range s.messageHist {
		res = res + line
	}
	s.syncMaster.Lock()
	conn.Write([]byte(res))
	defer s.syncMaster.Unlock()
}

// formating the name to look good
func (s *server) formatName(name string) string {
	return strings.ReplaceAll(strings.ReplaceAll(name, "\n", ""), " ", "")
}

// writing to the chat
func (s *server) writeToChat(sender, msg string, statusMsg bool) {
	s.syncMaster.Lock()
	prefix := s.addTimeStamp(sender)
	defer s.syncMaster.Unlock()
	for key, val := range s.users {

		if key != sender {
			if !statusMsg {
				val.Write([]byte("\n" + prefix + msg))
			} else {
				val.Write([]byte("\n" + msg))
			}
		}
		val.Write([]byte(s.addTimeStamp(key)))

	}
	s.messageHist = append(s.messageHist, prefix+msg)
}

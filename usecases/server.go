// we dont need that!!!

package usecases

import (
	"bufio"
	"net"
	"os"
	"sync"
)

type server struct {
	locker  *sync.Mutex
	address string
	users   map[string]net.Conn
	logfile string
}

type Connecter interface {
	StartServer() error
}

func NewServer(address, logfile string) Connecter {
	users := make(map[string]net.Conn)
	return &server{
		locker:  &sync.Mutex{},
		address: address,
		users:   users,
		logfile: logfile,
	}
}

func (s server) StartServer() error {
	lstn, err := net.Listen("tcp", "localhost:8080")
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
		go s.handleNewConn(conn)
	}
}

func (s server) handleNewConn(conn net.Conn) {
	var name string
	var err error
	defer s.removeUser(name)
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Write([]byte("Cannot Read name\n"))
			continue
		}
		if s.users[name] == nil {
			break
		} else {
			conn.Write([]byte("There is already a user with such name in the chat\n"))
		}
	}

	// for {
	// 	userInput, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		return
	// 	}
	// 	// msg := AddTimeStamp(userInput, user.Name)
	// 	// if err != nil {
	// 	// 	return
	// 	// }
	// 	// users.Range(func(key, value interface{}) bool {
	// 	// 	if conn, ok := value.(net.Conn); ok {
	// 	// 		if _, err := conn.Write([]byte(msg)); err != nil {
	// 	// 			return false
	// 	// 		}
	// 	// 	}
	// 	// 	return true
	// 	// })

	// 	// if err = SaveMessage(msg); err != nil {
	// 	// 	return
	// 	// }
	// }
}

func (s server) removeUser(name string) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.users[name].Close()
	delete(s.users, name)
}

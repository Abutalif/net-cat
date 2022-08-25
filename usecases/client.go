package usecases

import (
	"bufio"
	"fmt"
	"net"
)

type client struct {
	address string
}

type Talker interface {
	Connect() error
}

func NewClient(address string) Talker {
	return &client{
		address: address,
	}
}

func (c *client) Connect() error {
	// fmt.Println("Starting client mode")
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return err
	}

	go inputReader(conn)

	for {
		inMsg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return nil
		}
		fmt.Println(inMsg)
	}
}

func inputReader(conn net.Conn) {
	outMsg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	conn.Write([]byte(outMsg))
}

package usecases

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	fmt.Println("Starting client mode")
	conn, err := net.Dial("tcp", c.address)
	msgs := make([]byte, 10000)
	if err != nil {
		return err
	}

	n, err := conn.Read(msgs)
	if err != nil {
		return err
	}

	fmt.Print(n)
	fmt.Println(string(msgs[:n]))
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		_, err = conn.Write([]byte(sc.Text()))
		if err != nil {
			return err
		}
		n, err := conn.Read(msgs)
		if err != nil {
			return err
		}
		fmt.Println(string(msgs[:n]))
	}
	return nil
}

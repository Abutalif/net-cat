package usecases

import (
	"io"
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
	var finish chan error
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return err
	}
	defer conn.Close()

	go c.srcToDst(conn, os.Stdout, finish)
	c.srcToDst(os.Stdin, conn, finish)
	return nil
}

func (c *client) srcToDst(src io.Reader, dst io.Writer, finish chan error) {
	_, err := io.Copy(dst, src)
	if err != nil {
		finish <- err
	}
	finish <- nil
}

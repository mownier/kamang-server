package tcp_server

import (
	"bufio"
    "log"
	"net"
)

// Client

type Client struct {
	conn net.Conn
	Server *server
}

func (c *Client) start() {
	reader := bufio.NewReader(c.conn)
	for {
		data, err := reader.ReadBytes(10)
		if err != nil {
			c.conn.Close()
			c.Server.onDropClient(c, err)
			return
		}
		c.Server.onReceiveData(c, data)
	}
}

func (c *Client) ReceiveResponse(r string) error {
	_, err := c.conn.Write([]byte(r + "\n"))
	return err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetConn() net.Conn {
	return c.conn
}

// Server

type server struct {
	address string
	onAcceptClient func(c *Client)
	onDropClient func(c *Client, err error)
	onReceiveData func(c *Client, data []byte)
}

func (s *server) OnAcceptClient(callback func(c *Client)) {
	s.onAcceptClient = callback
}

func (s *server) OnDropClient(callback func(c *Client, err error)) {
	s.onDropClient = callback
}

func (s *server) OnReceiveData(callback func(c *Client, data []byte)) {
	s.onReceiveData = callback
}

func (s *server) Start() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting server")
	}
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		client := &Client {
			conn: conn,
			Server: s,
		}
		go client.start()
		s.onAcceptClient(client)
	}
}

func New(addr string) *server {
	log.Println("Created server:", addr)
	svr := &server {
		address: addr,
	}
	return svr
}


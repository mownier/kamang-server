package main

import (
	"log"
	"tcp_server"
)


func main() {
	server := tcp_server.New("localhost:9000")

	server.OnAcceptClient(func(c *tcp_server.Client) {
		log.Println("Client connected:",c)
	})

	server.OnDropClient(func(c *tcp_server.Client, err error) {
		log.Println("Drop client:", c)
	})

	server.OnReceiveData(func(c *tcp_server.Client, data []byte) {
		log.Println("Data:", string(data))
		c.ReceiveResponse("haler")
	})

	server.Start()
}

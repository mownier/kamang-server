package main

import (
	"log"
	"os"
	"io"
	"strings"
	"tcp_server"
)

func main() {
	server := tcp_server.New("localhost:9001")

	server.OnAcceptClient(func(c *tcp_server.Client) {
		log.Println("Connected client:", c)
	})

	server.OnDropClient(func(c *tcp_server.Client, err error) {
		log.Println("Drop client:", c)
	})

	server.OnReceiveData(func(c *tcp_server.Client, data []byte) {
		go processRequest(c, data)
	})

	server.Start()
}

func processRequest(c *tcp_server.Client, data []byte) {
	var request = string(data)
	var parts = strings.Split(request, ",")
	var assetId = parts[1]
	var assetName = parts[0]

	file, err := os.Open(assetName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
/*
	sendBuffer := make([]byte, 1024)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		c.GetConn().Write(sendBuffer)
	}
	c.ReceiveResponse("ok,"+assetName)
*/

	n, err := io.Copy(c.GetConn(), file)
	if err != nil {
		log.Fatal(err)
	}
	var resp = "\r\n" + assetId + "," + assetName + ",\n\n"
	c.GetConn().Write([]byte(resp))
	log.Println(resp)
	log.Println(n, "bytes sent.")
}


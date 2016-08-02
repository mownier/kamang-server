package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9000")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Send text: ")
		text, _ := reader.ReadString('\n')

		// Write to socket
		conn.Write([]byte(text))

		// Listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("[server]:", message)
	}
}


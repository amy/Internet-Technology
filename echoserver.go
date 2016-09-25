package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}

	for {
		// wait for & return next valid connection
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// go routine for connection logic
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	// close connection after handleClient returns
	defer conn.Close()

	fmt.Println("HELLO")

	var message [512]byte
	for {
		n, err := conn.Read(message[0:])
		if err != nil || string(message[0:]) == "#" {
			return
		}

		message = reverse(message)

		_, err = conn.Write(message[0:n])
		if err != nil {
			fmt.Printf("ERROR:%v", err)
			return
		}
	}
}

func reverse(message [512]byte) [512]byte {

	for i, j := 0, len(message)-1; i < j; i, j = i+1, j-1 {
		message[i], message[j] = message[j], message[i]
	}

	return message
}

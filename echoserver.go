package main

import (
	"net"
	"os"
)

func main() {

	arg := os.Args[1]

	listener, err := net.Listen("tcp", ":"+arg)
	if err != nil {
		return
	}

	for {
		// wait for & return next valid connection
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// connection logic
		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	// close connection after handleClient returns
	defer conn.Close()

	message := make([]byte, 512)

	for {
		readBytes, err := conn.Read(message)
		if err != nil {
			return
		}

		reverse(message[:readBytes])

		conn.Write(message[:readBytes])
	}
}

func reverse(message []byte) {

	for i, j := 0, len(message)-1; i < j; i, j = i+1, j-1 {
		message[i], message[j] = message[j], message[i]
	}
}

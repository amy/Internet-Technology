package main

import (
	"bufio"
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

	bufr := bufio.NewReader(conn)
	message := make([]byte, 1024)

	for {
		readBytes, err := bufr.Read(message)
		if err != nil {
			return
		}

		message = reverse(message)

		conn.Write([]byte(string(message[:readBytes])))
	}
}

func reverse(message []byte) []byte {

	for i, j := 0, len(message)-1; i < j; i, j = i+1, j-1 {
		message[i], message[j] = message[j], message[i]
	}

	return message
}

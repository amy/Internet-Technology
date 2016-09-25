package main

import (
	"bufio"
	"fmt"
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

		// Go routine to handle connection
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	// close connection when handleClient returns
	defer conn.Close()

	// scan each line of client input
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Bytes()

		// returns from handleClient when terminating characters given
		if closeConnection(string(message)) {
			return
		}

		// reverse message in place
		reverse(message)

		// stupid Java client won't recognize the end of a message without a newline char
		message = append(message, "\n"...)

		// print to standard out
		fmt.Print(string(message))

		// write message to client
		conn.Write(message)
	}
}

func reverse(message []byte) {

	for i, j := 0, len(message)-1; i < j; i, j = i+1, j-1 {
		message[i], message[j] = message[j], message[i]
	}
}

func closeConnection(message string) bool {

	if message == "#" || message == "$" {
		return true
	}

	return false
}

package main

import (
	//	"bufio"
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con) // Read the streams of buffer input from current client

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n') // It reads the string till new line

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest) // Trimspace will trim the space from head and tail of the string
			if clientRequest == ":QUIT" {
				log.Println("Current client requested server to close")
				return
			} else {
				log.Println(clientRequest)
			}
		case io.EOF:
			log.Println("Connection closed for the current Client")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		// Responding to the client request
		if _, err = con.Write([]byte("Got your response!\n")); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}

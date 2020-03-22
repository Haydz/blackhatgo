package main

import (
	"io"
	"log"
	"net"
)

//echo is a handler fucntion that simply echoes received data

func echo(conn net.Conn) {
	defer conn.Close()
	// Copy data from io.Reader to io.Writer via io.Copy().
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	//Bind to Port 20080 to listen
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")

	}
	log.Println("Listening on 0.0.0.0:20080")

	for {
		//Wait for connection. Create net.Conn on conns
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		//Handle connection using goroutine
		go echo(conn)
	}
}

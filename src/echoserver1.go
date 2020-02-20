package main

import (
	"io"
	"log"
	"net"
)

//echo is a handler fucntion that simply echoes received data

func echo(conn net.Conn) {
	defer conn.Close()

	//buffer to stre received data
	b := make([]byte, 512)

	for {
		//Receive data via conn.Read into a buffer/

		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Clinet disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpeceted error")
			break
		}
		//if not break
		log.Printf("received %d byes:%s\n", string(b))

		//send data via conn.Write
		log.Println("Writing Data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data)")
		}
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

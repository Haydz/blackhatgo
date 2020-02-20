package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")

	}

	defer dst.Close()
	//run in goroutine to prevent io.Copy from blocking
	go func() {
		//Copy our sources output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)

		}
	}()
	//copy our destinations output back to our source
	if _, err = io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	//listen on port 20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bidn to port")

	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")

		}
		go handle(conn)
	}

}

package main

import (
	"fmt"
	"net"
)

func main() {

	for i := 0; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		//fmt.Println(address)
		fmt.Printf("==scanningport  %d ==\n", i)
		//connect for each port
		conn, err := net.Dial("tcp", address)

		if err != nil {
			//port is closed or filtered, do nothing
			continue
		}

		//close connection, so not left open
		conn.Close()

		fmt.Printf("***%d open***\n", i)

	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

type Results struct {
	CommandsList []string
	ID           int
	Command      string
	Output       string
	Time         string
}

func main() {

	l, err := net.Listen("tcp", "127.0.0.1:9999")
	// l, err := net.Listen("tcp", *connect)

	checkError(err)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("===ADMIN SERVER LISTENING ===")
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PARENT CONNECTED")

	//LOOPING for read and writing sockets in NORMAL Mode
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		// if STOp is typed
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Fprintf(c, text+"\n")
			fmt.Println("Closing TCP server!")

			return
		}

		var outputTest Results
		decoder := json.NewDecoder(c)
		decoder.Decode(&outputTest)
		fmt.Println("===Results===")

		fmt.Printf("Time of command execution: :%s\n", outputTest.Time)
		fmt.Println(outputTest.Output)

	}

}

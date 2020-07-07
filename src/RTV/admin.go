package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

type Results struct {
	CommandsList []string
	ID           int
	//Command string
	Output string
	Time   time.Time
}

var (
	testArray = []string{"whoami", "hostname"}
	connect   = flag.String("connect", "", "127.0.0.19999")
	commands  = flag.Bool("commands", false, "Execute a list of commands")
)

func listMode() {

	l, err := net.Listen("tcp", *connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("COMMANDS to be run")
	for _, value := range testArray {
		fmt.Println(value)
	}

	inputTest := &Results{
		ID: 1,
		// Commands:     true,
		CommandsList: testArray,
	}

	encoder := json.NewEncoder(c)
	encoder.Encode(inputTest)

	var outputTest Results

	for x := 0; x < len(inputTest.CommandsList); x++ {
		decoder := json.NewDecoder(c)

		decoder.Decode(&outputTest)

		fmt.Println("===Results===")
		now := time.Now()

		fmt.Println("current time is: ", now) // time for logging
		// fmt.Println("\ntype is ", reflect.TypeOf(now))
		fmt.Println(outputTest.ID, outputTest.Output)
	}
}

func main() {

	//parsing flags
	flag.Parse()

	if *connect == "" {
		fmt.Println("empty connection")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *commands == true {
		fmt.Println("COMMAND SET")
	}

	if *commands == true {
		listMode()

	} else {

		// PORT := "127.0.0.1:9999"
		// PORT := os.Args[1]
		l, err := net.Listen("tcp", *connect)
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
		fmt.Println(reflect.TypeOf(c))

		fmt.Println("PARENT CONNECTED")

		for {
			// PORT := ":" + arguments[1]

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
			now := time.Now()

			fmt.Printf("current time is :%s", now)
			fmt.Println(reflect.TypeOf(now))
			fmt.Println(outputTest.ID, outputTest.Output)

		}
	}

}

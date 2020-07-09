package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
)

type Results struct {
	CommandsList []string
	ID           int
	//Command string
	Output string
	Time   string
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
		// currentTime := time.Now()

		// fmt.Printf("current time is :%s", currentTime.Format("2006-01-02 15:04:05")) // time for logging
		// fmt.Println("\ntype is ", reflect.TypeOf(now))
		fmt.Println("Time of command execution: ", outputTest.Time)
		fmt.Println(outputTest.ID, outputTest.Output)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {

	//parsing flags
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\ca.pem", "C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\ca.key")
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	if *connect == "" {
		fmt.Println("empty connection")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *commands == true {
		listMode()

	} else {
		//listener, err := tls.Listen("tcp", service, &config)
		l, err := tls.Listen("tcp", *connect, &config)
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

			fmt.Printf("Time of command execution: :%s\n", outputTest.Time)
			fmt.Println(outputTest.Output)
			//outputTest.ID

		}
	}

}

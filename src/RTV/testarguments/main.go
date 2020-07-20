package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 	logging := flag.NewFlagSet("log", flag.ContinueOnError)
	// 	loggingName := logging.String("log", "default.json", "Log to a json file. eg \"-log logfile\" will log commands to logfile.json")
	// 	loggingnoName := logging.Bool("log", false, "Add log if you would like command logs written to a file")

	// 	// loggingString := *loggingName
	// 	// loggingBool := *loggingnoName

	// 	fmt.Println("LOGGING: ", *loggingName, " ", *loggingnoName)

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("enable", "", "name")
	fooCmd.Parse(os.Args[2:])

	fmt.Println(*fooEnable, *fooName)
}

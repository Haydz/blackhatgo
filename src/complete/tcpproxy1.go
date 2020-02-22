package main

import (
	"fmt"
	"log"
	"os"
)

//FooReader definines an io.Reader to read from stdin

type FooReader struct{}

//Read reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")

	return os.Stdin.Read(b)
}

//FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{}

//writes data to stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)

}

func main() {
	//Instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	//create buffer to hold input/output
	input := make([]byte, 4096)

	//Use reader ti read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")

	}
	fmt.Printf("Read %d bytes from stdin \n", s)

	//Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")

	}

	fmt.Printf("Wrote %d bytes to stdout\n", s)

}

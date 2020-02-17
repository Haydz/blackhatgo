package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	//waitgroup for go routines
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		//adds one to go routine
		wg.Add(1)
		go func(j int) {
			// decrements 1 from go routine
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)

	}
	//block program from closing before goroutines done
	wg.Wait()

}

package main

import (
	"fmt"
	"net"
	"sort"
)

//func worker(ports chan int, wg *sync.WaitGroup) {
func worker(ports, results chan int) {
	for p := range ports {
		// fmt.Println(p)
		// wg.Done()
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0

			continue
		}
		conn.Close()
		results <- p

	}
}

func main() {
	ports := make(chan int, 100)
	//var wg sync.WaitGroup

	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)

	}
	//send workers into separate go routing, result gathering
	// loop needs to start before more than 100 works can cotinue
	go func() {
		for i := 1; i <= 1024; i++ {
			//wg.Add(1)
			ports <- i
		}
	}()

	//wg.Wait()
	//close(ports)

	//result gathering loop
	// receives on results channel 1024 times,
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	//closing channels once down
	close(ports)
	close(results)
	//sort slice so in order
	sort.Ints(openports)
	//iterate over slice and print to screen
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

	//testing upload from laptop
}

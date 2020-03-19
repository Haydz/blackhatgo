package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear... a 2")
		}
	}

}

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("recovered in clean up:", r)

	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()

	// need  a way to keep program open : wg.Wait
	// and tell WG when its done
	//DEFER - put off execution until function is done

}

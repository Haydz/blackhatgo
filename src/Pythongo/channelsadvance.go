package main

import (
	"fmt"
	"time"
)

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)
	go foo(fooVal, 5)
	go foo(fooVal, 3)
	v1 := <-fooVal
	v2 := <-fooVal
	fmt.Println(v1, v2)

	for i := 0; i < 10; i++ {
		go foo(fooVal, i)
	}

	close(fooVal)
	for item := range fooVal {
		fmt.Println(item)
	}
	time.Sleep(time.Second * 2)

}

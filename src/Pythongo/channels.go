package main

import "fmt"

func foo(c chan int, someValue int) {

	c <- someValue * 5

}

func main() {

	fooVal := make(chan int)
	go foo(fooVal, 5)
	go foo(fooVal, 10)
	v1 := <-fooVal
	v2 := <-fooVal
	fmt.Println(v1)
	fmt.Println(v2)

}

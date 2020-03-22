package main

import "fmt"

func foo() {

	//defer works on first in last out
	// like the stack
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done yet?")
	fmt.Println("Doing some stuff, who knows what?")
}

func foo2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	foo()
	foo2()
}

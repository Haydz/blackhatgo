package main

import "fmt"

func main() {
	x := 15

	//memory address
	a := &x
	fmt.Println(a)

	//read actual from memory
	fmt.Println(*a)

	*a = 5 // sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)
	println(x)

	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)

	fmt.Println(&x)
	fmt.Println(a)

}

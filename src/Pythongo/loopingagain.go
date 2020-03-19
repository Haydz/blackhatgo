package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 5
	for {
		fmt.Println("do stuff", x)
		x += 3
		if x > 25 {
			break
		}
	}

	//smaller

	for x := 5; x < 25; x += 3 {
		fmt.Println(x)
	}

	a := 3

	for x := 5; a < 25; x += 3 {
		fmt.Println("do stuff again", x)
		a += 4
	}

}

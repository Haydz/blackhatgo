package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func main() {
	{
		num1, num2 := 5.6, 9.2
		fmt.Println(add(num1, num2))
	}
}

package main

import "fmt"

func main() {

	//Each number in the sequence is the sum of the two numbers that precede it. So, the sequence goes: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, and so on.
	// The mathematical equation describing it is Xn+2= Xn+1 + Xn
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34,
	a := []int{0, 1}
	//a := make(opint)
	b := 0
	c := 1

	for i := 0; i < 20; i++ {

		sum := a[b] + a[c]
		fmt.Println("SUM: ", a[b], "=", a[c])
		a = append(a, sum)
		fmt.Println(sum)
		// adding next numbers
		b++
		c++

	}

}

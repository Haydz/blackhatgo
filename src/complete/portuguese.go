package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	places := []string{"restaurante",
		"newspaper", "hotel", "hospital", "library", "school", "pharmacy",
		"city"}
	x := rand.Intn(len(places))
	fmt.Println(places[x])
	fmt.Println("Just for Pull request test")
	//test to pull back down to computer
}

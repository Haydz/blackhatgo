package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woofwoof")
}

func f() {
	fmt.Println("f function")
}

//channels for goroutines to chat
func strleng(s string, c chan int) {
	c <- len(s)
}

//encoding
type encoding struct {
	Bar string
	Baz string
}

//xml
type xmls struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}

func main() {
	// fmt.Println("hello")

	// var count = int(42)
	// ptr := &count
	// fmt.Println(*ptr)
	// *ptr = 100
	// fmt.Println(count)

	var guy = new(Person)
	guy.Name = "Dave"
	guy.SayHello()

	Greet(guy)

	var dog = new(Dog)
	Greet(dog)

	nums := []int{2, 4, 6, 8}
	for abc, val := range nums {
		fmt.Println(abc, val)
	}

	x := 5
	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	// func test(Friend interface{}){
	// 	switch v := i.(type){
	// 	case int:
	// 		fmt.Println("I'm an integer!")
	// 	case string:
	// 		fmt.Println("I'm a string!")
	// 	default:
	// 		fmt.Println("Unknown type!")
	// 	}
	// 	}

	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	c := make(chan int)
	go strleng("Salutations", c)
	go strleng("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	g := encoding{"Joe Junior", "hello shabado"}
	b, _ := json.Marshal(g)
	fmt.Println(string(b))
	fmt.Print(json.Unmarshal(b, &g))

	h := xmls{"Joe Junior", "hello shabado"}
	b, _ := xml.h
	fmt.Println(string(h))
	//fmt.Print(json.Unmarshal(b, &g))

}

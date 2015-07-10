package main

import (
	"fmt"
	"reflect"
	//"math/rand"
	//"time"
)

func add(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	//rand.Seed(time.Now().Unix())
	//var x, y int = add(rand.Intn(10), rand.Intn(10))
	//fmt.Println("hello, world", x, y)
	//test_iota()
	//test_reflect()
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// func create_tree() {
// 	t := make(chan Tree)

// }

func test_iota() {
	const (
		A = 1
		C = 2
		D = iota
		B
		E
	)

	fmt.Printf("The value of A is %v\n", A)
	fmt.Printf("The value of C is %v\n", C)
	fmt.Printf("The value of D is %v\n", D)
	fmt.Printf("The value of B is %v\n", B)
	fmt.Printf("The value of E is %v\n", E)
}

func test_reflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind(), v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println(<-c)
			fmt.Println("quit")
			return
		}
	}
}

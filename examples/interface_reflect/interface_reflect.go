package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i float32 = 1.23
	var t interface{}
	t = i
	fmt.Println("i type: ", reflect.TypeOf(i).Name())
	fmt.Println("t type: ", reflect.TypeOf(t).Name())
	p := reflect.ValueOf(t).Float()
	fmt.Println("p value: ", p)
}

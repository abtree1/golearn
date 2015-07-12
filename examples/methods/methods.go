// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
    width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
    return r.width * r.height
}

func (r *rect) set(w,h int) int {
    r.width = w 
    r.height = h
    return 0 
}

func (r rect) setl(w,h int) int {
    r.width = w
    r.height = h 
    return 0
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // Here we call the 2 methods defined for our struct.
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    r.set(20, 10)
    r.setl(30, 20) //值类型不能修改外部值
    // Go automatically handles conversion between values
    // and pointers for method calls. You may want to use
    // a pointer receiver type to avoid copying on method
    // calls or to allow the method to mutate the
    // receiving struct.
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}

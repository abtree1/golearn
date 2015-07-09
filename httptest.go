package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h)
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h)
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!")
}

// func (h Hello) ServeHTTP(
// 	w http.ResponseWriter,
// 	r *http.Request) {
// 	switch {
// 	case r.RequestURI == "/string":
// 		fmt.Fprint(w, "I'm a frayed knot.")
// 	case r.RequestURI == "/struct":
// 		fmt.Fprint(w, "I'm a frayed struct.")
// 	default:
// 		fmt.Fprint(w, "hello")
// 	}
// }

func main() {
	// var h Hello
	// err := http.ListenAndServe("localhost:4000", h)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.Handle("/", &Hello{})
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

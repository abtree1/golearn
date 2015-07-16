package main

import (
	"database/sql"
	"fmt"
	"reflect"
	//"math/rand"
	//"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
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

	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci(c, quit)

	test_db()
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

var Table = map[string]string{
	"id":   "int",
	"name": "string",
	"pwd":  "string",
	"age":  "int",
}

func test_db() {
	db, err := sql.Open("mysql", "root:@/go_test?charset=utf8")
	if err != nil {
		fmt.Println("Open Sql Error", err.Error())
		db.Close()
	}

	rows, select_err := db.Query("SELECT * FROM users")
	if select_err != nil {
		fmt.Println("Open Sql Error", select_err.Error())
		db.Close()
	}

	var id int
	var name string
	var pwd string
	var age int
	refs := make([]interface{}, 4)
	rets := make([]interface{}, 4)
	refs[0] = &id
	refs[1] = &name
	refs[2] = &pwd
	refs[3] = &age
	for rows.Next() {
		//rows.Scan(&id, &name, &pwd, &age)
		rows.Scan(refs...)
		// fmt.Println("from db id:", id, " name:", name, " pwd:", pwd, " age:", age)
		rets[0] = *refs[0].(*int)
		rets[1] = *refs[1].(*string)
		rets[2] = *refs[2].(*string)
		rets[3] = *refs[3].(*int)
		fmt.Println("from db id:", rets[0].(int), " name:", rets[1].(string), " pwd:", rets[2].(string), " age:", rets[3].(int))
	}
	rows.Close()
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

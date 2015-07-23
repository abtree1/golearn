package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	//"strconv"
	//"math/rand"
	//"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/speedata/goxlsx"
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

	// keys := make([]string, 0, len(TABLE_LIST))
	// for k, _ := range TABLE_LIST {
	// 	keys = append(keys, k)
	// }
	// fmt.Println("TableName: ", keys[0], keys[1])

	//test_db()
	ini_parser()
}

func ini_parser() {
	file_path := "src/gs_tmp/static/test.ini"
	file, err := os.Open(file_path)
	ini_map := make(map[string]string)
	if err != nil {
		fmt.Println("open file error: ", err)
		file.Close()
		return
	}
	buf := bufio.NewReader(file)
	title := ""
	for {
		l, err := buf.ReadString('\n')
		if err == io.EOF {
			if len(l) == 0 {
				fmt.Println("empty file")
				file.Close()
				break
			}
		} else if err != nil {
			fmt.Println("read file error: ", err.Error())
			file.Close()
			return
		}
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		} else if strings.HasPrefix(l, "#") || strings.HasPrefix(l, ";") {
			continue
		} else if strings.HasPrefix(l, "[") {
			title = l
			fmt.Println("title", title)
			continue
		}
		i := strings.Index(l, "=")
		key := strings.TrimSpace(string([]byte(l)[:i]))
		value := strings.TrimSpace(string([]byte(l)[i+1:]))
		value = strings.TrimPrefix(value, "\"")
		value = strings.TrimSuffix(value, "\"")
		//fmt.Println("key: ", key, "value: ", value)
		if title == "" {
			ini_map[key] = value
		} else {
			key = title + key
			//fmt.Println("key", key)
			ini_map[key] = value
		}
	}
	fmt.Println("ini_map: ", ini_map)
	for _, v := range ini_map {
		r, _ := regexp.Compile("%\\(.*\\)")
		if r.MatchString(v) {
			fmt.Println("match", v)
			ss := r.FindStringSubmatch(v)
			fmt.Println(len(ss), ss[0])
			ii := r.FindStringSubmatchIndex(v)
			fmt.Println(ii)
			str := ""
			if ii[0] > 0 {
				str = string([]byte(v)[:ii[0]])
			}
			s := strings.TrimPrefix(ss[0], "%(")
			s = strings.TrimSuffix(s, ")")
			str += ini_map[s]
			if ii[1] < len(v) {
				str += string([]byte(v)[ii[1]:])
			}
			fmt.Println("regexp str: ", str)
		}
	}
}

func excel_parser() {
	file_path := "src/gs_tmp/static/test.xlsx"
	xlsx, err := goxlsx.OpenFile(file_path)
	if err != nil {
		fmt.Println("open error: ", err)
		panic(err)
	}
	i := xlsx.NumWorksheets()
	fmt.Println("NumWorksheets: ", i)
	work_sheet, err := xlsx.GetWorksheet(0)
	if err != nil {
		fmt.Println("open work sheet error: ", err)
		return
	}
	fmt.Println("file name", work_sheet.Name)
	fmt.Println("MaxRow: ", work_sheet.MaxRow, "MaxColumn: ", work_sheet.MaxColumn)
	fmt.Println("MinRow: ", work_sheet.MinRow, "MinColumn: ", work_sheet.MinColumn)
	for j := work_sheet.MinRow; j < work_sheet.MaxRow; j++ {
		for i := work_sheet.MinColumn; i <= work_sheet.MaxColumn; i++ {
			str := work_sheet.Cell(i, j)
			fmt.Println("column: ", i, "row: ", j, "value: ", str)
		}
	}
}

type Msg struct {
	handler <-chan int
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

var TABLE_LIST = map[string]map[string]string{
	"users": map[string]string{
		"id":   "int",
		"name": "string",
		"pwd":  "string",
		"age":  "string",
	},
	"user_conns": map[string]string{
		"id":      "int",
		"phone":   "string",
		"mobile":  "string",
		"email":   "string",
		"qq":      "string",
		"user_id": "int",
	},
}

func test_db() {
	db, err := sql.Open("mysql", "root:@/go_test?charset=utf8")
	if err != nil {
		fmt.Println("Open Sql Error", err.Error())
		db.Close()
	}

	params := make([]interface{}, 0, 1)
	params = append(params, 1)
	rows, select_err := db.Query("SELECT * FROM users where id=?", params...)
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

package main

import (
	"fmt"
	"net"
	"strconv"
	//"strings"
)

type Test struct {
	A int
	B rune
	C string
	D bool
}

func main() {
	// m := &Test{
	// 	A: 1,
	// 	B: 'a',
	// 	C: "test",
	// 	D: true,
	// }
	// exit := make(chan bool)

	// go func() {
	// 	fmt.Println("test m A =", m.A, " B=", m.B, " C=", m.C, " D=", m.D)
	// 	exit <- true
	// }()

	//Client()

	// s := make([]byte, 512)
	// index := copy(s[0:], "this is a test!")
	// fmt.Println(index)
	// //s = append(s, "this is a test")
	// fmt.Println("slices: len", len(s), " cap", cap(s))
	// index = copy(s[index:], "this is a test2!")
	// fmt.Println(index)
	// //s = append(s, "this is a test2!")
	// fmt.Println("slices: len", len(s), " cap", cap(s))
	// fmt.Println("string: ", string(s[0:index]))
	s := []byte{}
	// str := "this is a test!"
	// c := len(str)
	// s = strconv.AppendInt(s, int64(c), 10)
	// s = strconv.AppendQuote(s, str)
	s = strconv.AppendInt(s, 123, 10)
	// s = strconv.AppendBool(s, true)
	// s = strconv.AppendFloat(s, 1.23, 'f', -1, 64)
	fmt.Println("string:", string(s), " len:", len(s))

	// c := len("this is a test!")
	// str := strconv.Itoa(c) + "this is a test!"
	// str = str + strconv.Itoa(123)
	// str = str + strconv.FormatBool(true)
	// str = str + strconv.FormatFloat(1.23, 'f', -1, 64)

	// fmt.Println("reader")
	// reader := strings.NewReader(str)
	// b, _ := reader.ReadByte()
	// fmt.Println(strconv.Atoi(string(b)))

	// strconv.FormatFloat(f, fmt, prec, bitSize)
	//fmt.Println("bool: ", strconv.FormatBool(true))
	//fmt.Println("2 int: ", strconv.FormatInt(123, 10))
	//fmt.Println("36 int: ", strconv.FormatInt(123, 2))
	//fmt.Println("2 uint: ", strconv.FormatUint(123, 10))
	//fmt.Println("36 uint: ", strconv.FormatUint(123, 2))
	// fmt.Println("float b: ", strconv.FormatFloat(1.230, 'b', -1, 64))
	// fmt.Println("float E: ", strconv.FormatFloat(1.230, 'E', -1, 64))
	// fmt.Println("float e: ", strconv.FormatFloat(1.230, 'e', -1, 64))
	//fmt.Println("float f: ", strconv.FormatFloat(1.2301231, 'f', 6, 64))
	// fmt.Println("float G: ", strconv.FormatFloat(1.23012321, 'G', 6, 64))
	// fmt.Println("float g: ", strconv.FormatFloat(1.230123213, 'g', 6, 64))

	//<-exit
}

func Show(s string) string {
	return s
}

func Client() {
	client, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(Show("服务端连接失败"), Show(err.Error()))
		return
	}
	defer client.Close()
	buf := make([]byte, 1024)
	for i := 0; i < 10; i++ {
		client.Write([]byte("你好,服务端!"))
		c, err := client.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(Show(string(buf[0:c])))
	}
	client.Write([]byte("exit"))
}

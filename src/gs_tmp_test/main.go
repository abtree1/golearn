package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	//"reflect"
	//"strconv"
	//"strings"

	. "gs_tmp/utils"
)

type Test struct {
	A int
	B rune
	C string
	D bool
}

var clients = make(map[int]chan *Test)

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

	clients[0] = make(chan *Test)
	clients[1] = make(chan *Test)
	clients[2] = make(chan *Test)
	fmt.Println("map:", clients)

	//Client()

	// s := make([]byte, 0, 512)
	// index := copy(s[0:], "this is a test!")
	// fmt.Println(index)
	// str := "this is a test"
	// c := int32(len(str))
	//fmt.Printf("Type: %T,%i,%i", c, len(int), len(int32))
	// bytes := []byte{
	// 	byte(c >> 24),
	// 	byte(c >> 16),
	// 	byte(c >> 8),
	// 	byte(c),
	// }
	// s = append(s, bytes...)
	// s = append(s, str...)
	// fmt.Println("slices: len", len(s), " cap", cap(s))
	// index = copy(s[index:], "this is a test2!")
	// fmt.Println(index)
	// str = "this is a test2!"
	// c = int32(len(str))
	// bytes = []byte{
	// 	byte(c >> 24),
	// 	byte(c >> 16),
	// 	byte(c >> 8),
	// 	byte(c),
	// }
	// s = append(s, bytes...)
	// s = append(s, str...)
	// fmt.Println("slices: len", len(s), " cap", cap(s))
	// fmt.Println("string: ", string(s[0:len(s)]))

	// index := 0
	// buf := s[index:index+4]
	// index = index + 4
	// d := uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	// fmt.Println("back c:", int(d))
	// str = string(s[index:index+int(d)])
	// index = index+int(d)
	// fmt.Println("back str:", str)
	// buf = s[index:index+4]
	// index = index + 4
	// d = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	// fmt.Println("back c:", int(d))
	// str = string(s[index:index+int(d)])
	// index = index+int(d)
	// fmt.Println("back str:", str)
	//s := []byte{}
	// str := "this is a test!"
	// c := len(str)
	// s = strconv.AppendInt(s, int64(c), 10)
	// s = strconv.AppendQuote(s, str)
	//s = strconv.AppendInt(s, 123, 10)
	// s = strconv.AppendBool(s, true)
	// s = strconv.AppendFloat(s, 1.23, 'f', -1, 64)
	//fmt.Println("string:", string(s), " len:", len(s))

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

	// var i float32 = 1.23
	// var t interface{}
	// t = i
	// reflect.TypeOf(t)
	// fmt.Println("i type: ", reflect.TypeOf(i).Name())
	// fmt.Println("t type: ", reflect.TypeOf(t).Name())
	// p := reflect.ValueOf(t).Field(0)
	// fmt.Println("p= ", p)
	//fmt.Println("back= ", float32(p))
	// var i int16 = 2
	// var j int64 = int64(i)
	// fmt.Println("j= ", j)
	// buff := make([]byte, 2)
	// buff[0] = byte(j >> 8)
	// buff[1] = byte(j)
	// ret := uint16(buff[0]<<8) | uint16(buff[1])
	// fmt.Println("ret=", int16(ret))
}

func Client() {
	client, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(Show("服务端连接失败"), Show(err.Error()))
		return
	}
	defer client.Close()

	for i := 0; i < 10; i++ {
		bak := BuffFactory([]byte{})
		bak.WriteInt32(LOGIN_PARAM)
		bak.WriteString("你好，服务器!\r\n")
		bak.WriteBool(true)
		bak.WriteFloat32(1.23)
		bak.CompleteBuff()
		client.Write(bak.Data)

		head := make([]byte, 2)
		io.ReadFull(client, head)
		size := binary.BigEndian.Uint16(head)
		data := make([]byte, size)
		io.ReadFull(client, data)
		buff := BuffFactory(data)
		i32 := buff.ReadInt32()
		str := buff.ReadString()
		fmt.Println("category=", i32, " params=", str)
	}
	exit := BuffFactory([]byte{})
	exit.WriteInt32(EXIT_PARAM)
	exit.WriteString("你好，服务器!\r\n")
	exit.WriteBool(true)
	exit.WriteFloat32(4.56)
	exit.CompleteBuff()
	client.Write(exit.Data)
}

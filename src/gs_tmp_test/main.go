package main

import (
	"fmt"
	"net"
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

	Client()

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
	// buf := make([]byte, 1024)
	// for i := 0; i < 10; i++ {
	// 	client.Write([]byte("你好,服务端!"))
	// 	c, err := client.Read(buf)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	fmt.Println(Show(string(buf[0:c])))
	// }
	// client.Write([]byte("exit"))
	for i := 0; i < 10; i++ {
		bak := &Buffer{
			cur_p: 0,
			max_p: 0,
			data:  []byte{},
		}
		bak.WriteInt16(0)
		bak.WriteInt32(1)
		bak.WriteString("你好，服务器!\r\n")
		bak.WriteBool(true)
		bak.WriteFloat32(1.23)
		bak.Replace(0, bak.max_p)
		client.Write(bak.data)

		head := make([]byte, 2)
		io.ReadFull(client, head)
		size := binary.BigEndian.Uint16(head)
		data := make([]byte, size-2)
		io.ReadFull(client, data)
		buff = &Buffer{
			cur_p: 0,
			max_p: size - 2,
			data:  data,
		}
		str := buff.ReadString()
		fmt.Println(str)
	}
	exit := &Buffer{
		cur_p: 0,
		max_p: 0,
		data:  []byte{},
	}
	exit.WriteInt16(0)
	exit.WriteInt32(0)
	exit.WriteString("你好，服务器!\r\n")
	exit.WriteBool(true)
	exit.WriteFloat32(4.56)
	exit.Replace(0, exit.max_p)
	client.Write(bak.data)
}

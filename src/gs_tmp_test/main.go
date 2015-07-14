package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	. "gs_tmp/utils"
)

func main() {

	Client()
}

func Client() {
	client, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(Show("服务端连接失败"), Show(err.Error()))
		return
	}
	defer client.Close()

	login := BuffFactory([]byte{})
	login.WriteInt32(LOGIN_PARAM)
	login.WriteString("你好，服务器!\r\n")
	login.WriteBool(true)
	login.WriteFloat32(1.23)
	login.CompleteBuff()
	client.Write(login.Data)

	head := make([]byte, 2)
	io.ReadFull(client, head)
	size := binary.BigEndian.Uint16(head)
	data := make([]byte, size)
	io.ReadFull(client, data)
	buff := BuffFactory(data)
	i32 := buff.ReadInt32()
	str := buff.ReadString()
	fmt.Println("category=", i32, " params=", str)

	for i := 0; i < 10; i++ {
		test_message(client)
	}

	exit := BuffFactory([]byte{})
	exit.WriteInt32(EXIT_PARAM)
	exit.CompleteBuff()
	client.Write(exit.Data)
}

func test_message(client *net.TCPConn) {
	buff := BuffFactory([]byte{})
	buff.WriteInt32(TEST_PARAM)
	buff.WriteString("你好，服务器!\r\n")
	buff.WriteBool(true)
	buff.WriteFloat32(1.23)
	buff.CompleteBuff()
	client.Write(buff.Data)

	head := make([]byte, 2)
	io.ReadFull(client, head)
	size := binary.BigEndian.Uint16(head)
	data := make([]byte, size)
	io.ReadFull(client, data)
	buff = BuffFactory(data)
	i32 := buff.ReadInt32()
	str := buff.ReadString()
	fmt.Println("category=", i32, " params=", str)
}

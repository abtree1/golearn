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

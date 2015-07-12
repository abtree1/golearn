package connections

import (
	"fmt"
	"net"
	"io"
	"binary"
	"strconv"

	. "gs_tmp/utils"
)

func Server() {
	exit := make(chan bool)
	ip := net.ParseIP("127.0.0.1")
	addr := net.TCPAddr{ip, 8888, "0:0:0:0:0:0:0:1"}
	go func() {
		listen, err := net.ListenTCP("tcp", &addr)
		if err != nil {
			fmt.Println(Show("初始化失败"), Show(err.Error()))
			exit <- true
			return
		}
		fmt.Println(Show("正在监听..."))
		for {
			client, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(Show("客户端连接"), Show(client.RemoteAddr().String()))

			go ClientRead(client)
			// go func() {
			// 	for {
			// 		msg, err := client.Read()
			// 		client.Write(msg)
			// 	}
			// }

			// data := make([]byte, 1024)
			// c, err := client.Read(data)
			// if err != nil {
			// 	fmt.Println(Show(err.Error()))
			// }
			// fmt.Println(Show(string(data[0:c])))
			// client.Write([]byte("你好客户端!\r\n"))
			// client.Close()
		}
	}()
	<-exit
	fmt.Println(Show("服务端关闭!"))
}

func ClientRead(client *net.TCPConn) {
	// data := make([]byte, 1024)
	// for {
	// 	c, err := client.Read(data)
	// 	if err != nil {
	// 		fmt.Println(Show(err.Error()))
	// 	}
	// 	str := string(data[0:c])
	// 	fmt.Println(Show(str))
	// 	if str == "exit" {
	// 		break
	// 	}
	// 	client.Write([]byte("你好客户端!\r\n"))
	// }
	for {
		head := make([]byte, 2)
		io.ReadFull(client, head)
		size := binary.BigEndian.Uint16(head)
		data := make([]byte, size -2)
		io.ReadFull(client, data)
		buff = &Buffer{
			cur_p:0,
			max_p:size -2,
			data:data
		}
		i32 := buff.ReadInt32()
		str := buff.ReadStrng()
		b := buff.ReadBool()
		f32 := buff.ReadFloat32()
		s := []byte{}
		s = strconv.AppendInt(s, int64(i32), 10)
		s = strconv.AppendQuote(s, str)
		s = strconv.AppendBool(s, b)
		s = strconv.AppendFloat(s, f32, 'f', -1, 32)
		fmt.Println("resecive:  ", s)
		if i32 == 0{
			break
		}

		bak := &Buffer{
			cur_p:0
			max_p:0
			data:[]byte{}
		}
		bak.WriteInt16(0)
		bak.WriteString("你好客户端!\r\n")
		bak.Replace(0, bak.max_p)
		client.Write(bak.data)
	}
	client.Close()
}

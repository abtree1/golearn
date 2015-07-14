package connections

import (
	"fmt"
	"io"
	"net"
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
		}
	}()
	<-exit
	fmt.Println(Show("服务端关闭!"))
}

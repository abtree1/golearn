package connections

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	. "gs_tmp/utils"
)

var playerids = make(map[*net.TCPConn]int)

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

func ClientRead(client *net.TCPConn) {
	head := make([]byte, 2)
	for {
		io.ReadFull(client, head)
		size := binary.BigEndian.Uint16(head)
		data := make([]byte, size)
		io.ReadFull(client, data)
		buff := BuffFactory(data)
		category := buff.ReadInt32()
		var player_id int
		if category == LOGIN_PARAM {
			player_id = SessionLogin(client, buff)
			playerids[client] = player_id
		} else if category == EXIT_PARAM {
			player_id = playerids[client]
			HandleRequest(player_id, category, buff)
			delete(playerids, player_id)
			break
		} else {
			player_id = playerids[client]
			HandleRequest(player_id, category, buff)
		}
	}
	client.Close()
}

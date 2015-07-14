package connections

import (
	"encoding/binary"
	"fmt"
	"net"

	"gs_tmp/controllers"
	. "gs_tmp/utils"
)

func ClientRead(conn *net.TCPConn) {
	head := make([]byte, 2)
	var handler chan *Msg
	for {
		io.ReadFull(conn, head)
		size := binary.BigEndian.Uint16(head)
		data := make([]byte, size)
		io.ReadFull(conn, data)
		buff := BuffFactory(data)
		category := buff.ReadInt32()
		if category == LOGIN_PARAM {
			handler = SessionLogin(conn, buff)
		} else if category == EXIT_PARAM {
			HandleRequest(handler, category, buff)
			break
		} else {
			HandleRequest(handler, category, buff)
		}
	}
	fmt.Println(Show("客户端退出!"), Show(conn.RemoteAddr().String()))
	conn.Close()
}

func SessionLogin(buff *Buffer) chan *Msg {
	// player_id = 1
	handler := make(chan *Msg)
	c := &Client{
		Client:  client,
		Handler: handler,
	}
	// clients[player_id] = c

	go controllers.RunController(c)
	HandleRequest(handler, LOGIN_PARAM, buff)

	return handler
}

func HandleRequest(handler chan<- *Msg, category int, buff *Buffer) {
	msg := &Msg{
		Category: category,
		Buff:     buff,
	}
	fmt.Println("Receive Type:", category)
	handler <- msg
}

// func WrapClient(player_id int, msg *Msg) {
// 	clients[player_id].Handler <- msg
// }

// func SendResponse(player_id int, buff *Buffer) {
// 	clients[player_id].Client.Write(buff.Data)
// }

package connections

import (
	"encoding/binary"
	"fmt"
	"io"
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
		if category == PROTOCOL_LOGIN_PARAM {
			handler = SessionLogin(conn, buff)
		} else if category == PROTOCOL_EXIT_PARAM {
			HandleRequest(handler, category, buff)
			close(handler)
			break
		} else {
			HandleRequest(handler, category, buff)
		}
	}
	fmt.Println(Show("客户端退出!"), Show(conn.RemoteAddr().String()))
	conn.Close()
}

func SessionLogin(client *net.TCPConn, buff *Buffer) chan *Msg {
	handler := make(chan *Msg)
	go controllers.RunController(client, handler)
	HandleRequest(handler, PROTOCOL_LOGIN_PARAM, buff)

	return handler
}

func HandleRequest(handler chan<- *Msg, category int32, buff *Buffer) {
	msg := &Msg{
		Category: category,
		Buff:     buff,
	}
	fmt.Println("Receive Type:", category)
	handler <- msg
}

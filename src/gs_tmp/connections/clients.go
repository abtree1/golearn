package connections

import (
	"fmt"
	"net"

	"gs_tmp/helpers"
	. "gs_tmp/utils"
)

var clients = make(map[int]*Client)

func SessionLogin(client *net.TCPConn, buff *Buffer) int {
	player_id = len(playerids) + 1
	c := &Client{
		Client:  client,
		Handler: make(chan *Msg),
	}
	clients[player_id] = c

	go helpers.RunController(c.Handler)
	HandleRequest(player_id, LOGIN_PARAM, buff)

	return player_id
}

func RunClient(c chan *Msg) {
	msg := <-c
}

func HandleRequest(player_id int, category int, buff *Buffer) {
	msg := &Msg{
		PlayerId: player_id,
		Category: category,
		Buff:     buff,
	}
	clients[player_id].Handler <- msg
}

func WrapClient(player_id int, msg *Msg) {
	clients[player_id].Handler <- msg
}

func SendResponse(player_id int, buff *Buffer) {
	clients[player_id].Client.Write(buff.Data)
}

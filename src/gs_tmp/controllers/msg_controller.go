package controllers

import (
	"net"

	. "gs_tmp/utils"
)

func RunController(msg *Msg, proxy chan<- *Msg) {
	hand := make(chan *Msg)
	c := &Client{
		Handler: hand,
		Proxy:   proxy,
	}
	c.conn_observer(msg.PlayerId)
	c.HandleMsg(msg)
	for {
		select {
		case msg := <-c.Handler:
			if c.HandleMsg(msg) {
				return
			}
		default: // do nothing
		}
	}
}

func (client *Client) conn_client(msg *Msg) {
	cli := make(chan *net.TCPConn, 1)
	back := &Msg{
		Handler: cli,
	}
	msg.Handler.(chan *Msg) <- back
	back1 := &Msg{
		Handler: client.Handler,
	}
	msg.Handler.(chan *Msg) <- back1
	client.Client = <-cli
}

func (client *Client) conn_observer(player_id int) {
	add := &Msg{
		PlayerId: player_id,
		Category: PROXY_ADD_PLAYER,
		Handler:  client.Handler,
	}
	client.Proxy <- add
}

func (client *Client) HandleMsg(msg *Msg) bool {
	switch msg.Category {
	case PROTOCOL_LOGIN_PARAM:
		client.conn_client(msg)
		client.Login(msg)
	case PROTOCOL_TEST_PARAM:
		client.Test(msg)
	case PROTOCOL_EXIT_PARAM:
		client.LoginOut(msg)
		return true
	case PROTOCOL_WRAP_PARAM:
		client.Wrap(msg)
	case PROXY_GET_INFO:
		client.GetInfo(msg)
	}
	return false
}

func (client *Client) SendClient(buff *Buffer) {
	buff.CompleteBuff()
	client.Client.Write(buff.Data)
}

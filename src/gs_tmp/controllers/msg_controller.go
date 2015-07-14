package controllers

import (
	"net"

	. "gs_tmp/utils"
)

type Client struct {
	Client  *net.TCPConn
	Handler <-chan *Msg
}

func RunController(client *net.TCPConn, handler <-chan *Msg) {
	c := &Client{
		Client:  client,
		Handler: handler,
	}
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

func (client *Client) HandleMsg(msg *Msg) bool {
	switch msg.Category {
	case LOGIN_PARAM:
		client.Login(msg.Buff)
	case TEST_PARAM:
		client.Test(msg.Buff)
	case EXIT_PARAM:
		return true
	}
	return false
}

func (client *Client) SendClient(buff *Buffer) {
	buff.CompleteBuff()
	client.Client.Write(buff.Data)
}

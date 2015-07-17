package controllers

import (
	"net"

	"gs_tmp/observer"
	. "gs_tmp/utils"
)

func RunController(client *net.TCPConn, handler <-chan *Msg) {
	proxy := make(chan *ObMsg)
	c := &Client{
		Client:  client,
		Handler: handler,
		Proxy:   proxy,
	}
	add2observer(proxy)
	for {
		select {
		case msg := <-c.Handler:
			if c.HandleMsg(msg) {
				return
			}
		case obmsg := <-c.Proxy:
			c.HandleProxy(obmsg)
		default: // do nothing
		}
	}
}

func add2observer(proxy chan<- *ObMsg) {
	buff := BuffNoClose()
	buff.WriteInt32(PROXY_ADD_PLAYER)
	add := &ObMsg{
		PlayerId: 0,
		Buff:     buff,
		Handler:  proxy,
	}
	observer.Proxy(add)
}

func (client *Client) HandleMsg(msg *Msg) bool {
	switch msg.Category {
	case PROTOCOL_LOGIN_PARAM:
		client.Login(msg.Buff)
	case PROTOCOL_TEST_PARAM:
		client.Test(msg.Buff)
	case PROTOCOL_EXIT_PARAM:
		return true
	}
	return false
}

func (client *Client) HandleProxy(msg *observer.ObMsg) {

}

func (client *Client) SendClient(buff *Buffer) {
	buff.CompleteBuff()
	client.Client.Write(buff.Data)
}

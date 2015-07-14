package controllers

import (
	. "gs_tmp/utils"
)

func RunController(c *Client) {
	for {
		select {
		case msg := <-c.Handler:
			if c.handle_msg(msg) {
				return
			}
		default: // do nothing
		}
	}
}

func (client *Client) handle_msg(msg *Msg) bool {
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

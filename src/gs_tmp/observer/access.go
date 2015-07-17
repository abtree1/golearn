package observer

import (
	. "gs_tmp/utils"
)

var clients = map[int]chan<- *ObMsg{}
var writes = make(chan *ObMsg)

func RunObserver() {
	for {
		write := <-writes
		switch write.player_id {
		case 0:
			write.handler()
		default:
			write.proxy()
		}
	}
}

func (msg *ObMsg) handler() {
	category := msg.Buff.ReadInt32()
	switch category {
	case PROXY_ADD_PLAYER:
		msg.add_player()
	}
}

func (msg *ObMsg) add_player() {
	clients[msg.PlayerId] = msg.Handler
}

func (msg *ObMsg) proxy() {
	clients[msg.player_id] <- msg
}

func Proxy(msg *ObMsg) {
	writes <- msg
}

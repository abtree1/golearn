package helpers

import (
	"gs_tmp/connections"
	"gs_tmp/controllers"
	. "gs_tmp/utils"
)

func RunController(c chan *Msg) {
	for {
		select {
		case msg := <-c:
			if handle_msg(msg) {
				return
			}
		default: // do nothing
		}
	}
}

func handle_msg(msg *Msg) bool {
	switch msg.Category {
	case LOGIN_PARAM:
		controllers.Login(msg.PlayerId, msg.Buff)
	case EXIT_PARAM:
		return true
	}
	return false
}

func WrapClient(player_id int, buff *Buffer) {
	connections.WrapClient(player_id, buff)
}

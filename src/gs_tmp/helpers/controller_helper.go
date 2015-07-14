package helpers

// import (
// 	"gs_tmp/controllers"
// 	. "gs_tmp/utils"
// )

// func RunController(c *Client) {
// 	for {
// 		select {
// 		case msg := <-c.Handler:
// 			if c.handle_msg(msg) {
// 				return
// 			}
// 		default: // do nothing
// 		}
// 	}
// }

// func (client *Client) handle_msg(msg *Msg) bool {
// 	switch msg.Category {
// 	case LOGIN_PARAM:
// 		controllers.Login(msg.PlayerId, msg.Buff)
// 	case TEST_PARAM:
// 		controllers.Test(msg.PlayerId, msg.Buff)
// 	case EXIT_PARAM:
// 		return true
// 	}
// 	return false
// }

// func (client *Client) SendClient(buff *Buffer) {
// 	buff.CompleteBuff()
// 	c.Client.Write(buff.Data)
// }

// func WrapClient(player_id int, buff *Buffer) {
// 	buff.CompleteBuff()
// 	connections.WrapClient(player_id, buff)
// }

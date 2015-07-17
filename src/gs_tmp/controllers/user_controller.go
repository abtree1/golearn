package controllers

import (
	"fmt"
	"net"

	"gs_tmp/models"
	. "gs_tmp/utils"
)

type Client struct {
	Client     *net.TCPConn
	PlayerData map[string]*models.TableDb
	Handler    <-chan *Msg
}

func (client *Client) Login(buff *Buffer) {
	id := int(buff.ReadInt32())
	client.PlayerData = models.LoadAllPlayerData(id)

	user := client.PlayerData["users"].Data[0]
	id = user["id"].(int)
	name := user["name"].(string)
	pwd := user["pwd"].(string)
	age := user["age"].(int)
	fmt.Println("resecive user: id = ", id, " name=", name, " pwd=", pwd, " age=", age)

	user_conn := client.PlayerData["user_conns"].Data[0]
	id = user_conn["id"].(int)
	phone := user_conn["phone"].(string)
	mobile := user_conn["mobile"].(string)
	email := user_conn["email"].(string)
	qq := user_conn["qq"].(string)
	user_id := user_conn["user_id"].(int)
	fmt.Println("resecive user_conn: id = ", id, " phone=", phone, " mobile=", mobile, " email=", email, " qq=", qq, " user_id=", user_id)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(PROTOCOL_LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

func (client *Client) Test(buff *Buffer) {
	str := buff.ReadString()
	b := buff.ReadBool()
	f32 := buff.ReadFloat32()
	fmt.Println("resecive: str= ", str, " b=", b, " f32=", f32)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(PROTOCOL_LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

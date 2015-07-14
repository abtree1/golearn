package controllers

import (
	"fmt"

	"gs_tmp/helpers"
	. "gs_tmp/utils"
)

func (client *Client) Login(buff *Buffer) {
	str := buff.ReadString()
	b := buff.ReadBool()
	f32 := buff.ReadFloat32()
	fmt.Println("resecive: str= ", str, " b=", b, " f32=", f32)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

func (client *Client) Test(buff *Buffer) {
	str := buff.ReadString()
	b := buff.ReadBool()
	f32 := buff.ReadFloat32()
	fmt.Println("resecive: str= ", str, " b=", b, " f32=", f32)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

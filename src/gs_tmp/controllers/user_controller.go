package controllers

import (
	. "gs_tmp/utils"
)

func Login(player_id int, buff *Buffer) {
	str := buff.ReadString()
	b := buff.ReadBool()
	f32 := buff.ReadFloat32()
	fmt.Println("resecive: i32= ", player_id, " str= ", str, " b=", b, " f32=", f32)
}

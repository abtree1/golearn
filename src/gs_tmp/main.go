package main

import (
	"fmt"

	"gs_tmp/connections"
	"gs_tmp/models"
)

func main() {
	exit := make(chan bool)
	go models.RunDb(exit)
	go connections.Server(exit)
	<-exit
	fmt.Println("服务端关闭!")
}

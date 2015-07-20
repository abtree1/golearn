package main

import (
	"fmt"

	"gs_tmp/connections"
	"gs_tmp/models"
	"gs_tmp/observer"
)

func main() {
	exit := make(chan bool)
	go models.RunDb(exit)
	go connections.Server(exit)
	go observer.RunObserver()
	<-exit
	fmt.Println("服务端关闭!")
}

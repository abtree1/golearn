package main

import (
	"gs_tmp/connections"
	"gs_tmp/models"
)

func main() {
	connections.Server()
	models.InitDb()
}

package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	. "gs_tmp/utils"
)

type TableDb struct {
	Static int
	Data   map[string]interface{}
}

var db *sql.DB

func InitDb() {
	var err error
	db, err := sql.Open("mysql", "root:@/go_test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func LoadAllPlayerData() {

}

func (this *TableDb) LoadData() {

}

package models

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	. "gs_tmp/utils"
)

type TableDb struct {
	Static int
	Name   string
	Data   []map[string]interface{}
}

type Operation struct {
	table_name string
	sql        string
	params     []interface{}
	back       bool
	handler    chan<- []*TableDb
}

var writes = make(chan *Operation)

func InitDb() {
	db, err := sql.Open("mysql", "root:@/go_test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	go run_db(db)
}

func run_db(db *sql.DB) {
	for {
		write := <-writes
		if write.back {
			write.call_handler(db)
		} else {
			write.cast_handler(db)
		}
	}
}

func (write *Operation) cast_handler(db *sql.DB) error {
	_, err := db.Exec(write.sql, write.params...)
	return err
}

func (write *Operation) call_handler(db *sql.DB) error {
	rows, err := db.Query(write.sql, write.params...)
	if err != nil {
		rows.Close()
		return err
	}
	stable := TABLE_LIST[write.table_name]
	length := len(stable)
	refs := make([]interface{}, length)
	i := 0
	for _, v := range stable {
		switch v {
		case "int":
			var id int
			refs[i] = &id
		case "string":
			var sd string
			refs[i] = &sd
		}
		i++
	}
	dts := make([]map[string]interface{})
	for rows.Next() {
		rows.Scan(refs...)
		i = 0
		db_data := make(map[string]interface{})
		for k, v := range stable {
			switch v {
			case "int":
				db_data[k] = *refs[i].(*int)
			case "string":
				db_data[k] = *refs[i].(*string)
			}
			i++
		}
		append(dts, db_data)
		//fmt.Println("from db id:", rets[0].(int), " name:", rets[1].(string), " pwd:", rets[2].(string), " age:", rets[3].(int))
	}
	fmt.Println("table_name:", write.table_name, " datas:", dts)
	tb := &TableDb{
		Static: TABLE_LOAD,
		Name:   write.table_name,
		Data:   dts,
	}
	write.handler <- tb
	return nil
}

func LoadAllPlayerData(playerid int) []*TableDb {
	length := len(TABLE_LIST)
	params := []int{playerid}
	h := make(chan *TableDb)
	for k, _ := range TABLE_LIST {
		if k == "users" {
			ps := []string{k, "id"}
			s := LoadTable(ps)
			fmt.Println("sql users:", s)
			op := &Operation{
				table_name: k,
				sql:        s,
				params:     params,
				back:       true,
				handler:    h,
			}
			writes <- op
		} else {
			ps := []string{k, "user_id"}
			s := LoadTable(ps)
			fmt.Println("sql user_conns:", s)
			op := &Operation{
				table_name: k,
				sql:        s,
				params:     params,
				back:       true,
				handler:    h,
			}
			writes <- op
		}
	}
	datas := make(map[string]*TableDb)
	for i := 0; i < length; i++ {
		data := <-h
		datas[data.Name] = data
	}
	fmt.Println("sql tables:", datas)
	h.Close()
	return datas
}

func LoadData(table_name string, id int) {

}

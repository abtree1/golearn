package utils

func LoadTable(params []string) string {
	length := len(params)
	sql := "select * from "
	if length == 1 {
		sql += params[0]
	} else {
		sql += params[0] + " where "
		for i := 1; i < length; i++ {
			sql += params[i] + "=$" + i
		}
	}
	return sql
}

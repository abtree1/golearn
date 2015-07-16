package utils

const (
	PROTOCOL_EXIT_PARAM = iota
	PROTOCOL_LOGIN_PARAM
	PROTOCOL_LOGIN_BAK
	PROTOCOL_TEST_PARAM
)

const (
	TABLE_GET = iota
	TABLE_LOAD
	TABLE_SET
	TABLE_DEL
	TABLE_FIND
	TABLE_SELECT
	TABLE_PERSIST
)

const (
	TABLE_LIST = map[string]map[string]string{
		"users": map[string]string{
			"id":   "int",
			"name": "string",
			"pwd":  "string",
			"age":  "string",
		},
		"user_conns": map[string]string{
			"id":      "int",
			"phone":   "string",
			"mobile":  "string",
			"email":   "string",
			"qq":      "string",
			"user_id": "int",
		},
	}
)

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
	TABLE_LIST = map[string][]string{
		"users":      []string{"id", "name", "pwd", "age"},
		"user_conns": []string{"id", "phone", "mobile", "email", "qq", "user_id"},
	}
)

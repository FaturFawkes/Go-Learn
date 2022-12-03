package database

var conn string

func init() {
	conn = "MySQL"
}

func GetDatabase() string {
	return conn
}
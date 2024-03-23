package config

// Конфигурация для подключения к базе данных
const (
	host     = "host=postgres2"
	port     = "port=5432"
	user     = "user=adm"
	password = "password=pwd"
	dbname   = "dbname=aleshina"
	sslmode  = "sslmode=disable"
)

const ConnStr = host + " " + port + " " + user + " " + password + " " + dbname + " " + sslmode

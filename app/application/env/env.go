package env

import "os"

type EnvKey string

const (
	MysqlUser     EnvKey = "MYSQL_USER"
	MysqlPassword EnvKey = "MYSQL_PASSWORD"
	MysqlDatabase EnvKey = "MYSQL_DATABASE"
	MysqlServerIp EnvKey = "MYSQL_SERVER_IP"
	MysqlPort     EnvKey = "MYSQL_PORT"
	JwtSecret     EnvKey = "JWT_SECRET"
)

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

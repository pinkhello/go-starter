package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-starter/config"
	"go.uber.org/fx"
	"net/url"
)

var Module = fx.Provide(NewMysqlDB)

func NewMysqlDB(config config.Config) (*sql.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	return sql.Open(`mysql`, dsn)
}

package libs

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"go-starter/config"
	"net/url"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func NewXorm(config config.Config) *xorm.Engine {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	val := url.Values{}
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	engine, err := xorm.NewEngine(config.Database.Driver, dsn)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	//engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(10)
	return engine
}

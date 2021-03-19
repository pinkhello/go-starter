package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	configFile = "config/config.yml"
	configType = "yml"
)

type (
	Config struct {
		Debug          bool     `mapstructure:"debug"`
		ContextTimeout int      `mapstructure:"contextTimeout"`
		Server         Server   `mapstructure:"server"`
		Database       Database `mapstructure:"database"`
		Nsq            Nsq      `mapstructure:"nsq"`
	}

	Server struct {
		Address string `mapstructure:"address"`
	}

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	}

	Nsq struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}
)

func NewConfig() Config {
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return *conf
}

func InitConfig() {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}

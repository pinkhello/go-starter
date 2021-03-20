package utils

import (
	"github.com/spf13/viper"
	"time"
)

func NewTimeoutContext() time.Duration {
	timeout := time.Duration(viper.GetInt("contextTimeout")) * time.Second

	return timeout
}

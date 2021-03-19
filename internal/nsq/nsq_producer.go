package nsq

import (
	"context"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"go-starter/config"
	"go.uber.org/fx"
)

var ProducerModule = fx.Provide(NewProducer)

func NewProducer(lifecycle fx.Lifecycle, config config.Config) (*nsq.Producer, error) {
	address := fmt.Sprintf("%s:%d", config.Nsq.Host, config.Nsq.Port)
	producer, err := nsq.NewProducer(address, nsq.NewConfig())
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logrus.Print("Start Nsq Producer....")
			go func() {
				err := producer.Ping()
				if err != nil {
					logrus.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Print("Stopping Nsq Producer....")
			producer.Stop()
			return nil
		},
	})
	return producer, err
}

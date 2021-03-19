package nsq

import (
	"context"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"go-starter/config"
	"go-starter/internal/service"
	"go.uber.org/fx"
)

var (
	ConsumerModule = fx.Invoke(NewConsumers)
)

func NewConsumers(lifecycle fx.Lifecycle, config config.Config, groupService service.BusinessGroupService) (*nsq.Consumer, error) {

	consumer, err := buildConsumer(groupService)

	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%d", config.Nsq.Host, config.Nsq.Port)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logrus.Print("Start Nsq Consumers....")
			go func() {
				err := consumer.ConnectToNSQD(address)
				if err != nil {
					logrus.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Print("Stopping Nsq Consumers....")
			consumer.Stop()
			return nil
		},
	})

	return consumer, nil

}

type (
	AConsumerhandler struct {
		groupService service.BusinessGroupService
	}
)

func (c *AConsumerhandler) HandleMessage(message *nsq.Message) error {
	//todo 处理消息
	logrus.Info(fmt.Sprintf("%s => %s", "NSQ_A", string(message.Body)))
	return nil
}

func buildConsumer(groupService service.BusinessGroupService) (*nsq.Consumer, error) {
	//todo 也可以 config 配置
	consumer, err := nsq.NewConsumer("test", "test-channel", nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(&AConsumerhandler{
		groupService: groupService,
	})
	return consumer, err
}

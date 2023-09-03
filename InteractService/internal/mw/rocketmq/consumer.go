package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"interact_service/internal/constant"
	"sync"
)

func InitConsumer() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(constant.RocketMQGroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{constant.RocketMQNameServerAddr})),
	)
	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to new pullConsumer: %s", err), nil)
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)

	defer func() {
		if err := recover(); err != nil {
			rlog.Fatal(fmt.Sprintf("panic: %s", err), nil)
		}
	}()
	topic := constant.DBFavoriteWriterTopic
	err = c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, message := range msgs {
			msg := string(message.Body)
			fmt.Println(msg)
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to subscribe: %s", err), nil)
	}

	err = c.Start()

	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to new pullConsumer: %s", err), nil)
	}

	wg.Wait()
}

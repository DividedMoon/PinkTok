package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"interact_service/internal/constant"
)

var p rocketmq.Producer

func InitProducer() {
	p, _ = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{constant.RocketMQNameServerAddr})),
		producer.WithRetry(2),
		producer.WithGroupName(constant.RocketMQGroupName),
	)

	err := p.Start()
	if err != nil {
		return
	}
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
	}
}

func SendMessage(topic string, message string) {

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(message),
	}

	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
}

package rocketmq

import (
	"interact_service/internal/constant"
	"testing"
)

func TestRocketMQ(t *testing.T) {
	InitProducer()
	SendMessage(constant.DBFavoriteWriterTopic, "hello")
	InitConsumer()

}

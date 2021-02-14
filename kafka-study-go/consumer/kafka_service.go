package consumer

import (
	"errors"
	"fmt"
	"strings"
)

type KafkaService struct {
	Hosts      []string
	handlerMap map[string]*KafkaConsumer
}
type KafkaMessageCallback func(topic, key, value string, offset int64)

type KafkaMessageHandler func(topic, key, value string, offset int64)

func NewKafkaService(hosts []string) (*KafkaService, error) {
	if len(hosts) <= 0 {
		return nil, errors.New(fmt.Sprintf("发生异常: %v", "无法检测到hosts内容"))
	}
	return &KafkaService{
		Hosts:      hosts,
		handlerMap: map[string]*KafkaConsumer{},
	}, nil
}

func (k *KafkaService) StartConsume(group string, topic []string, callback KafkaMessageCallback) (bool, error) {
	handler := &KafkaConsumer{
		Id:       GetHandlerIds(group, topic...),
		Hosts:    k.Hosts,
		Group:    group,
		Topic:    topic,
		Callback: callback,
	}
	k.handlerMap[handler.Id] = handler
	ok, err := handler.StartConsume()
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (k *KafkaService) StopService() {
	for _, handler := range k.handlerMap {
		handler.Stop()
	}
}

func (k *KafkaService) KafKaServiceHandler(handler KafkaMessageHandler) KafkaMessageCallback {
	return func(topic, key, value string, offset int64) {
		handler(topic, key, value, offset)
	}
}

func GetHandlerIds(group string, topics ...string) string {
	return group + "-" + strings.Join(topics, "-")
}

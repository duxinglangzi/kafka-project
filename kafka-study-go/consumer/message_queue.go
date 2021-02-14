package consumer

import (
	"os"
	"time"
)

const (
	MQ_INIT_FAILED = 1
)

var kafkaService *KafkaService

func InitMQ(hosts []string) {
	kafkaService, err := NewKafkaService(hosts)
	if err != nil {
		time.Sleep(time.Second)
		os.Exit(MQ_INIT_FAILED)
	}
	// 加载监听
	kafkaService.InitKafkaHandler()
}

func StopMQ() {
	if kafkaService != nil {
		kafkaService.StopService()
	}
}

package consumer

import (
	"fmt"
	"log"
)

func (k *KafkaService) InitKafkaHandler() {
	_, err := k.StartConsume("kafka_group", []string{"test-kafka-topic"}, k.KafKaServiceHandler(KafkaTestHandler))
	if err != nil {
		fmt.Println("Kafka Consumer failed, error log : " + err.Error())
	}
}

func KafkaTestHandler(topic string, key string, value string, offset int64) {
	log.Printf("打印topic信息, topic name: %v , key: %v , offset:%v , value: %v \n", topic, key, offset, value)
}

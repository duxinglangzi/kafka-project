package producer

import (
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

type KafkaProducer struct {
	Hosts []string
	Init  bool
	sync.Mutex
	sarama.SyncProducer
}

var kafka_producer = KafkaProducer{}

func SetKafkaSetting(hosts []string) {
	fmt.Printf("当前连接地址: %v \n", hosts)
	kafka_producer.Hosts = hosts
}

func GetKafkaProducer() (*KafkaProducer, error) {
	if kafka_producer.SyncProducer == nil {
		err := kafka_producer.init()
		if err != nil {
			return &kafka_producer, err
		}
	}
	return &kafka_producer, nil
}

func (kafka *KafkaProducer) init() error {
	kafka.Lock()
	defer kafka.Unlock()
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(kafka.Hosts, cfg)
	if err != nil {
		return err
	}
	kafka_producer.Init = true
	kafka_producer.SyncProducer = p
	return nil
}

func (kafka *KafkaProducer) SendByHashPartition(topic, key string, data string) (int32, int64, interface{}, error) {
	if !kafka.Init {
		return 0, 0, nil, errors.New("not init kafka producer before")
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(data),
	}
	partition, offset, err := kafka.SendMessage(msg)
	if err != nil {
		return partition, offset, data, err
	}
	return partition, offset, data, nil
}

package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type KafkaConsumer struct {
	Id         string
	cancelList []context.CancelFunc
	consume    sarama.ConsumerGroup
	Hosts      []string
	Group      string
	Topic      []string
	Callback   KafkaMessageCallback
}

func (k *KafkaConsumer) StartConsume() (bool, error) {
	if len(k.Hosts) <= 0 || len(k.Topic) <= 0 || k.Group == "" || k.Callback == nil {
		return false, errors.New("Kafka Config Error")
	}
	keepConsumer := false
	cfg := sarama.NewConfig()
	master, err := sarama.NewConsumerGroup(k.Hosts, k.Group, cfg)
	if err != nil {
		return false, err
	}

	defer func() {
		if !keepConsumer {
			master.Close()
		}
	}()
	k.consume = master
	consumer := Consumer{
		ready:    make(chan bool),
		callback: k.Callback,
	}

	ctx, cancel := context.WithCancel(context.Background())
	k.cancelList = append(k.cancelList, cancel)

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := k.consume.Consume(ctx, k.Topic, &consumer); err != nil {
				fmt.Sprintf("Error from consumer: %v", err)
				time.Sleep(time.Second * 10) //wait for 10 second to retry
				continue
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")
	keepConsumer = true
	return true, nil
}

func (k *KafkaConsumer) Stop() {
	for _, c := range k.cancelList {
		c()
	}
	if k.consume != nil && k.consume.Close() != nil {
		k.consume.Close()
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready    chan bool
	callback KafkaMessageCallback
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		if consumer.callback != nil {
			consumer.callback(message.Topic, string(message.Key), string(message.Value), message.Offset)
		}
		fmt.Sprintf("Message Queue claimed: timestamp = %v, topic = %s, offset = %d, value = %s", message.Timestamp, message.Topic, message.Offset, string(message.Value))
		session.MarkMessage(message, "")
	}

	return nil
}

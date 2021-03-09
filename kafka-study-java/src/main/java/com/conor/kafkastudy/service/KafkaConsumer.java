package com.conor.kafkastudy.service;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import java.util.Optional;

/**
 * <p>Description: </p>
 * <p>@Author conor  2021/1/30 </p>
 */
@Component
public class KafkaConsumer {


    @KafkaListener(topics = {"test-kafka-topic"}, groupId = "test-consumer-group")
    public void listen(ConsumerRecord<?, ?> record) {

        if (record.value() != null) {
            System.out.println("[>>>>>>>>>>>>>>>>] offset : "+ record.offset());
            System.out.println("[>>>>>>>>>>>>>>>>] topic : "+ record.topic());
            System.out.println("[>>>>>>>>>>>>>>>>] toString : "+ record.toString());
        }

//        Optional.ofNullable(record.value())
//                .ifPresent(message -> {
//                    System.out.println("【+++++++++++++++++ record = {} 】" + record);
//                    System.out.println("【+++++++++++++++++ message = {}】" + message);
//                });
    }

    @KafkaListener(topics = {"filebeats-topic"}, groupId = "filebeats-comsumer-group")
    public void listenKafkaFilebeats(ConsumerRecord<?, ?> record) {

        if (record.value() != null) {
            System.out.println("[>>>>>>>>>>>>>>>>] filebeats-offset : "+ record.offset());
            System.out.println("[>>>>>>>>>>>>>>>>] filebeats-topic : "+ record.topic());
            System.out.println("[>>>>>>>>>>>>>>>>] filebeats-toString : "+ record.toString());
        }

//        Optional.ofNullable(record.value())
//                .ifPresent(message -> {
//                    System.out.println("【+++++++++++++++++ record = {} 】" + record);
//                    System.out.println("【+++++++++++++++++ message = {}】" + message);
//                });
    }

}

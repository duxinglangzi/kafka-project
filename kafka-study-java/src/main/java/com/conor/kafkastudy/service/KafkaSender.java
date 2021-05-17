package com.conor.kafkastudy.service;

import com.conor.kafkastudy.entity.Message;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.SendResult;
import org.springframework.stereotype.Component;
import org.springframework.util.concurrent.ListenableFuture;

import java.time.LocalDateTime;
import java.util.concurrent.ExecutionException;

/**
 * <p>Description: </p>
 * <p>@Author conor  2021/1/30 </p>
 */
@Component
public class KafkaSender {

    private final KafkaTemplate<String, String> kafkaTemplate;

    //构造器方式注入 kafkaTemplate
    public KafkaSender(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    private Gson gson = new GsonBuilder().create();

    public void send(String msg) {
        Message message = new Message();

        message.setId(System.currentTimeMillis());
        message.setMsg(msg);
        message.setSendTime(LocalDateTime.now());
        //对 topic = hello2 的发送消息
//        kafkaTemplate.send("test-kafka-topic",gson.toJson(message));
        ListenableFuture<SendResult<String, String>> sfp_push = kafkaTemplate.send("sfp_push", gson.toJson(message));
        try {
            String s = sfp_push.get().toString();
            System.out.println("<<<<<<< 测试返回值 :   "+s);
        } catch (InterruptedException | ExecutionException e) {
            e.printStackTrace();
        }

    }

}

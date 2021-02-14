package com.conor.kafkastudy;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(basePackages = { "com.conor" })
public class KafkaStudyApplication {

    public static void main(String[] args) {

        SpringApplication.run(KafkaStudyApplication.class, args);
    }

}

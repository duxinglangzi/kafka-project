package com.conor.kafkastudy.controller;

import com.conor.kafkastudy.service.KafkaSender;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

/**
 * <p>Description: </p>
 * <p>@Author conor  2021/1/30 </p>
 */
@RestController
public class AppController {


    @Autowired
    private KafkaSender kafkaSender;


    @GetMapping("sendMessage")
    public void sendMessage(@RequestParam("msg") String msg){
        kafkaSender.send(msg);
    }
}
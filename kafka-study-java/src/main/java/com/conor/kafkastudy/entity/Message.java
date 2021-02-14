package com.conor.kafkastudy.entity;

import java.time.LocalDateTime;

/**
 * <p>Description: </p>
 * <p>@Author conor  2021/1/30 </p>
 */
public class Message {

    private Long id;

    private String msg;

    private LocalDateTime sendTime;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public LocalDateTime getSendTime() {
        return sendTime;
    }

    public void setSendTime(LocalDateTime sendTime) {
        this.sendTime = sendTime;
    }
}

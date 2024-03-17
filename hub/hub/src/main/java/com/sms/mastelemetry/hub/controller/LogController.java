package com.sms.mastelemetry.hub.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Controller;

import java.security.Principal;

@Controller
public class LogController  {

    private SimpMessagingTemplate template;

    @Autowired
    public LogController(SimpMessagingTemplate template) {
        this.template = template;
    }

    /* /app/queue/log */
    @MessageMapping("/log")
    public void listUsers(@Payload String payload, Principal p) throws Exception {
        String clientName = p.getName();
        if(clientName != null) {
            template.convertAndSend("/topic/log/" + clientName, payload);
        } else {
            template.convertAndSend("/topic/log/orphaned", payload);
        }
    }



}

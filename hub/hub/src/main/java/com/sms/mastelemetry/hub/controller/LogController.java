package com.sms.mastelemetry.hub.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.handler.annotation.DestinationVariable;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Controller;

import java.security.Principal;
import java.util.List;
import java.util.ArrayList;

import com.sms.mastelemetry.hub.components.TopicRegistry;

@Controller
public class LogController  {

    private SimpMessagingTemplate template;
    private TopicRegistry topicRegistry;

    @Autowired
    public LogController(SimpMessagingTemplate template, TopicRegistry topicRegistry) {
        this.template = template;
        this.topicRegistry = topicRegistry;
    }

    /* /app/log */
    @MessageMapping("/log/{podname}")
    public void listUsers(@DestinationVariable("podname") String podname, @Payload String payload, Principal p) throws Exception {
        String clientName = p.getName();
        String topicName = "/topic/log/" + clientName + "/" + podname;
        if( ! topicRegistry.getTopics().contains(topicName) ) {
            topicRegistry.addTopic(topicName);
        }
        if(clientName != null) {
            template.convertAndSend(topicName, payload);
        } else {
            template.convertAndSend("/topic/log/orphaned", payload);
        }
    }



}

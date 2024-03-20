package com.sms.mastelemetry.hub.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.messaging.simp.user.SimpUser;
import org.springframework.messaging.simp.user.SimpUserRegistry;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.messaging.simp.SimpMessagingTemplate;

import com.fasterxml.jackson.databind.util.JSONPObject;
import com.sms.mastelemetry.hub.models.User;
import com.sms.mastelemetry.hub.components.TopicRegistry;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@RestController
public class UserController  {

    private SimpUserRegistry userRegistry; 
    private TopicRegistry topicRegistry;

    @Autowired
    public UserController(SimpUserRegistry userRegistry, TopicRegistry topicRegistry) {
        this.userRegistry = userRegistry;
        this.topicRegistry = topicRegistry;
    }


    @CrossOrigin(origins = "*")
    @RequestMapping(path="/users/{userid}", method = RequestMethod.GET)
    public ResponseEntity<List<String>> getQueuesForUser(@PathVariable("userid") String userId) {
        List<String> result = topicRegistry.getTopics().stream()
            .filter(element -> element.contains(userId))
            .collect(Collectors.toList());
        return ResponseEntity
            .ok()
            .body(result);
    }

    @CrossOrigin(origins = "*")
    @RequestMapping(path="/users", method = RequestMethod.GET)
    public ResponseEntity<List<User>> listUsers() {
        List<User> result = new ArrayList<>(); 
        for(SimpUser simpUser : this.userRegistry.getUsers()) {
            User user = new User(); 
            user.setNodeName(simpUser.getName());
            result.add(user);
        }
        if(result.isEmpty()) {
            return ResponseEntity
                .notFound()
                .build();
        }
        return ResponseEntity
            .ok()
            .body(result);
    }


}

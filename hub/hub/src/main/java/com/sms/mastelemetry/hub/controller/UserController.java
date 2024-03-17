package com.sms.mastelemetry.hub.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.messaging.simp.user.SimpUser;
import org.springframework.messaging.simp.user.SimpUserRegistry;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.fasterxml.jackson.databind.util.JSONPObject;
import com.sms.mastelemetry.hub.models.User;

import java.util.ArrayList;
import java.util.List;

@RestController
public class UserController  {

    private SimpUserRegistry userRegistry; 

    @Autowired
    public UserController(SimpUserRegistry userRegistry) {
        this.userRegistry = userRegistry;
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

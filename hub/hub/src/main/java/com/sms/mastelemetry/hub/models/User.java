package com.sms.mastelemetry.hub.models;


public class User {
    public String nodeName;
    public String queueName;
    public User() { }
    public User(String nodeName) {
        this.nodeName = nodeName;
    }
    public String getNodeName() {
        return this.nodeName;
    }
    public void setNodeName(String client) {
        this.nodeName = client;
    }
}

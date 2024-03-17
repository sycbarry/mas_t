
package com.sms.mastelemetry.hub.models;


public class LogMessage {
    private String instance; 
    private String message; 
    public LogMessage(String instance, String message) {
        this.instance = instance;
        this.message = message;
    }
    public LogMessage() {}
    public String getFromInstance() { return this.instance; } 
    public void setFromInstance(String instance) { this.instance = instance; } 
    public String getMessage() { return this.message; } 
    public void setMessage(String message) { this.message = message; } 
}

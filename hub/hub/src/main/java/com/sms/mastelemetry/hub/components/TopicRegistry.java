package com.sms.mastelemetry.hub.components;

import org.springframework.stereotype.Component;
import java.util.List;
import java.util.ArrayList;

@Component
public class TopicRegistry {

    private List<String> topicList = new ArrayList<>();

    public List<String> getTopics() {
        return topicList;
    }

    public void addTopic(String topic) {
        if( ! topicList.contains(topic) ) {
            topicList.add(topic);
        }
    }

}


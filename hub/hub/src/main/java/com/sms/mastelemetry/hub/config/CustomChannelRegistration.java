
package com.sms.mastelemetry.hub.config;

import org.springframework.messaging.Message;
import org.springframework.messaging.MessageChannel;
import org.springframework.messaging.simp.stomp.StompCommand;
import org.springframework.messaging.simp.stomp.StompHeaderAccessor;
import org.springframework.messaging.support.ChannelInterceptor;
import org.springframework.messaging.support.MessageHeaderAccessor;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

public class CustomChannelRegistration implements ChannelInterceptor {
        @Override
        public Message<?> preSend(Message<?> message, MessageChannel channel) {
            StompHeaderAccessor accessor = 
                        MessageHeaderAccessor.getAccessor(message, StompHeaderAccessor.class);
            if(StompCommand.CONNECT.equals(accessor.getCommand())) {
                Authentication user = 
                    SecurityContextHolder.getContext().getAuthentication();
                accessor.setUser(user); 
            }
            return message;
        }
}

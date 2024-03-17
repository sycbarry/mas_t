
package com.sms.mastelemetry.hub.config;

import java.security.Principal;
import java.util.Map;
import java.util.UUID;

import org.springframework.http.server.ServerHttpRequest;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.server.support.DefaultHandshakeHandler;

import com.sms.mastelemetry.hub.models.AnonymousPrincipal;

public class CustomHandShakeHandler extends DefaultHandshakeHandler {

    @Override
    protected Principal determineUser(ServerHttpRequest request,WebSocketHandler wsHandler, Map<String, Object> attributes) {
        Principal principal = request.getPrincipal();
        if (principal == null) {
            principal = new AnonymousPrincipal();
            String uniqueName = UUID.randomUUID().toString();
            ((AnonymousPrincipal) principal).setName(uniqueName);
        }
        return principal;
    }

}

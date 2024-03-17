package com.sms.mastelemetry.hub.config;

import java.util.Map;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.server.HandshakeInterceptor;


@Order(Ordered.HIGHEST_PRECEDENCE + 1000000)
public class HttpHandShakeInterceptor implements HandshakeInterceptor {

	@Override
	public boolean beforeHandshake(ServerHttpRequest request, ServerHttpResponse response, 
            WebSocketHandler wsHandler, Map attributes) throws Exception {
		return true;
	}

	public void afterHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler, Exception ex) {
        if(request.getHeaders() == null) return;
        if(request.getHeaders().get("Sec-WebSocket-Protocol") == null) return;
        String protocol = (String) request.getHeaders().get("Sec-WebSocket-Protocol").get(0);
        if(protocol == null) return;
        response.getHeaders().add("Sec-WebSocket-Protocol", protocol);
	}
}

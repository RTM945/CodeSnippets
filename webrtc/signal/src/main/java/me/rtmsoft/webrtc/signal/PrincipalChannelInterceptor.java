package me.rtmsoft.webrtc.signal;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.Message;
import org.springframework.messaging.MessageChannel;
import org.springframework.messaging.simp.SimpMessageHeaderAccessor;
import org.springframework.messaging.simp.stomp.StompCommand;
import org.springframework.messaging.simp.stomp.StompHeaderAccessor;
import org.springframework.messaging.support.ChannelInterceptor;
import org.springframework.messaging.support.MessageHeaderAccessor;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.WebSocketSession;

import java.io.IOException;
import java.util.LinkedList;
import java.util.Map;

@Component
public class PrincipalChannelInterceptor implements ChannelInterceptor {

    private static final Logger LOGGER = LoggerFactory.getLogger(PrincipalChannelInterceptor.class);

    @Autowired
    private PeerManager peerManager;

    @Autowired
    private SessionManager sessionManager;

    @Override
    public Message<?> preSend(Message<?> message, MessageChannel channel) {
        StompHeaderAccessor accessor = MessageHeaderAccessor.getAccessor(message, StompHeaderAccessor.class);
        if (accessor != null && accessor.getCommand() == StompCommand.CONNECT) {
            String sessionId = accessor.getSessionId();
            Object raw = message
                    .getHeaders()
                    .get(SimpMessageHeaderAccessor.NATIVE_HEADERS);
            if (raw instanceof Map) {
                try {
                    String token = ((LinkedList) ((Map) raw).get("token")).get(0).toString();
                    String type = ((LinkedList) ((Map) raw).get("type")).get(0).toString();
                    if ("offer".equals(type) || "answer".equals(type)) {
                        if (!peerManager.isRegistered(token)) {
                            Peer peer = new Peer(token, type, sessionId);
                            peerManager.add(peer);
                            accessor.setUser(peer);
                            LOGGER.info(sessionId + " set principal " + peer.getName());
                        }else{
                            sessionManager.close(sessionId, CloseStatus.SERVICE_RESTARTED);
                        }
                    }else {
                        sessionManager.close(sessionId, CloseStatus.PROTOCOL_ERROR);
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                    sessionManager.close(sessionId, CloseStatus.PROTOCOL_ERROR);
                }
            }
        }

        return message;
    }
}

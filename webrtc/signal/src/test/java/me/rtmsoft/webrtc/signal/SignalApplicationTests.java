package me.rtmsoft.webrtc.signal;

import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.messaging.converter.StringMessageConverter;
import org.springframework.messaging.simp.stomp.StompFrameHandler;
import org.springframework.messaging.simp.stomp.StompHeaders;
import org.springframework.messaging.simp.stomp.StompSession;
import org.springframework.messaging.simp.stomp.StompSessionHandlerAdapter;
import org.springframework.util.concurrent.ListenableFuture;
import org.springframework.web.socket.WebSocketHttpHeaders;
import org.springframework.web.socket.client.standard.StandardWebSocketClient;
import org.springframework.web.socket.messaging.WebSocketStompClient;
import org.springframework.web.socket.sockjs.client.SockJsClient;
import org.springframework.web.socket.sockjs.client.WebSocketTransport;

import java.lang.reflect.Type;
import java.util.Collections;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class SignalApplicationTests {

    private static final Logger LOGGER = LoggerFactory.getLogger(SignalApplicationTests.class);

    @LocalServerPort
    int port;

    @Test
    void testConnection() throws Exception {
        String baseUrl = "ws://localhost:" + port + "/signalling";
        String token = UUID.randomUUID().toString();
        StompSession answer = createPeer(baseUrl, token, "answer");
        StompSession offer = createPeer(baseUrl, token, "offer");
        offer.send("/app/sdp", "offer's sdp");
        TimeUnit.SECONDS.sleep(1);
    }

    @Test
    void testWrongType() {
        String baseUrl = "ws://localhost:" + port + "/signalling";
        WebSocketStompClient stompClient = new WebSocketStompClient(new SockJsClient(Collections.singletonList(new WebSocketTransport(new StandardWebSocketClient()))));
        stompClient.setMessageConverter(new StringMessageConverter());
        StompHeaders stompHeaders = new StompHeaders();
        stompHeaders.add("token", UUID.randomUUID().toString());
        stompHeaders.add("type", "wrong type");
        stompClient.connect(baseUrl, (WebSocketHttpHeaders) null, stompHeaders, new StompSessionHandlerAdapter() {
            @Override
            public void handleTransportError(StompSession session, Throwable exception) {
                exception.printStackTrace();
            }
        });
    }

    StompSession createPeer(String baseUrl, String token, String type) throws Exception {
        WebSocketStompClient stompClient = new WebSocketStompClient(new SockJsClient(Collections.singletonList(new WebSocketTransport(new StandardWebSocketClient()))));
        stompClient.setMessageConverter(new StringMessageConverter());
        StompHeaders stompHeaders = new StompHeaders();
        stompHeaders.add("token", token);
        stompHeaders.add("type", type);
        ListenableFuture<StompSession> future = stompClient.connect(baseUrl, (WebSocketHttpHeaders) null, stompHeaders, new StompSessionHandlerAdapter() {});
        StompSession stompSession = future.get(1, TimeUnit.SECONDS);
        stompSession.subscribe("/user/queue/onsdp", new StompFrameHandler() {
            @Override
            public Type getPayloadType(StompHeaders headers) {
                return String.class;
            }

            @Override
            public void handleFrame(StompHeaders headers, Object payload) {
                LOGGER.info(type + " receive " + payload);
                if(type.equals("answer")){
                    stompSession.send("/app/sdp", "answer's sdp");
                }
            }
        });
        return stompSession;
    }
}

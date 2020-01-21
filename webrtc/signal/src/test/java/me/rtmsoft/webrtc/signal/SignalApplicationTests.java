package me.rtmsoft.webrtc.signal;

import me.rtmsoft.webrtc.signal.vo.PairVO;
import me.rtmsoft.webrtc.signal.vo.SdpVO;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
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
import java.net.URI;
import java.net.URISyntaxException;
import java.util.Collections;
import java.util.concurrent.TimeUnit;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class SignalApplicationTests {

    private Logger logger = LoggerFactory.getLogger(SignalApplicationTests.class);

    @Autowired
    private TestRestTemplate restTemplate;

    @LocalServerPort
    int port;

    @Test
    void contextLoads() {
    }

    String user = "test";
    String token = "test_token";
    String offer = "test_offer";
    String answer = "test_answer";

    String wsOnErrors = "/queue/errors";
    String wsOffer = "/queue/offer";
    String wsRegister = "/app/register";

    @Test
    void testAll() throws Exception {
        //1.register
        String baseUrl = "ws://localhost:" + port + "/signalling";
        WebSocketStompClient stompClient = new WebSocketStompClient(new SockJsClient(Collections.singletonList(new WebSocketTransport(new StandardWebSocketClient()))));
        StompHeaders stompHeaders = new StompHeaders();
        stompHeaders.add("user", user);
        ListenableFuture<StompSession> future = stompClient.connect(baseUrl, (WebSocketHttpHeaders) null, stompHeaders, new StompSessionHandlerAdapter() {});
        StompSession stompSession = future.get(1, TimeUnit.SECONDS);
        stompSession.subscribe(wsOnErrors, new WSErrorHandler(stompSession));
        stompSession.subscribe(wsOffer, new WSOfferHandler(stompSession));
        stompSession.send(wsRegister + "/" + token, null);

        //2.pair
        baseUrl = "http://localhost:" + port + "/pair";
        URI uri = new URI(baseUrl);
        PairVO pairVO = new PairVO();
        pairVO.setName(user);
        pairVO.setToken(token);

        HttpHeaders headers = new HttpHeaders();
        HttpEntity<PairVO> pairRequest = new HttpEntity<>(pairVO, headers);
        ResponseEntity<String> result = restTemplate.postForEntity(uri, pairRequest, String.class);
        if (result.getStatusCode() == HttpStatus.OK) {
            baseUrl = "http://localhost:" + port + "/offer";
            uri = new URI(baseUrl);
            SdpVO sdpVO = new SdpVO();
            sdpVO.setSdp(offer);
            headers = new HttpHeaders();
            HttpEntity<SdpVO> offerRequest = new HttpEntity<>(sdpVO, headers);
            result = restTemplate.postForEntity(uri, offerRequest, String.class);
        }
        logger.info(result.toString());
    }

    private class WSErrorHandler implements StompFrameHandler {

        StompSession stompSession;

        public WSErrorHandler(StompSession stompSession) {
            this.stompSession = stompSession;
        }

        @Override
        public Type getPayloadType(StompHeaders headers) {
            return String.class;
        }

        @Override
        public void handleFrame(StompHeaders headers, Object payload) {
            logger.info("{} errors --> {}", stompSession.getSessionId(), payload);
        }
    }

    private class WSOfferHandler implements StompFrameHandler {

        StompSession stompSession;

        public WSOfferHandler(StompSession stompSession) {
            this.stompSession = stompSession;
        }

        @Override
        public Type getPayloadType(StompHeaders headers) {
            return SdpVO.class;
        }

        @Override
        public void handleFrame(StompHeaders headers, Object payload) {
            logger.info("{} 收到offer --> {}", stompSession.getSessionId(), payload.toString());
            StompSession.Receiptable receipt = stompSession.send("/app/answer", answer);
            receipt.addReceiptLostTask(() -> logger.info("{} lost answer", stompSession.getSessionId()));
        }
    }
}

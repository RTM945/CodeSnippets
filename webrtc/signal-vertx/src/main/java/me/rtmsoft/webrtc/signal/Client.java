package me.rtmsoft.webrtc.signal;

import io.vertx.core.http.ServerWebSocket;
import io.vertx.core.json.Json;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public class Client {

    private String token;
    private ServerWebSocket ws;
    private CountDownLatch latch;

    String answer;

    public Client(String token, ServerWebSocket ws) {
        this.token = token;
        this.ws = ws;
    }

    public void send(Dto dto) {
        ws.writeTextMessage(Json.encode(dto));
        latch = new CountDownLatch(1);
    }

    public String getAnswer() throws InterruptedException {
        if (latch != null) {
            latch.await(1, TimeUnit.SECONDS);
        }
        return answer;
    }

    public void setAnswer(String answer) {
        this.answer = answer;
        if (latch != null) {
            latch.countDown();
        }
    }
}

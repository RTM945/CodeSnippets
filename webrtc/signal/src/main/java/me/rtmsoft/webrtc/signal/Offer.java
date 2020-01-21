package me.rtmsoft.webrtc.signal;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

public class Offer {

    private BlockingQueue<String> response = new LinkedBlockingQueue<>(1);

    private final String user;
    private final long start;

    public Offer(String user) {
        this.user = user;
        this.start = System.currentTimeMillis();
    }

    public String getUser() {
        return user;
    }

    public long getStart() {
        return start;
    }

    public void setResponse(String sdp) throws InterruptedException {
        response.put(sdp);
    }

    public String getAnswer(int timeout, TimeUnit unit) throws InterruptedException {
        return response.poll(timeout, unit);
    }
}

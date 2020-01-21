package me.rtmsoft.webrtc.signal;

import org.springframework.stereotype.Service;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

@Service
public class OfferManager {

    private final Map<String, Offer> offers = new ConcurrentHashMap<>();

    public OfferManager() {
        ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor();
        scheduler.scheduleAtFixedRate(new Runnable() {
            @Override
            public void run() {
                offers.keySet().forEach(k -> {
                    Offer offer = offers.get(k);
                    if (System.currentTimeMillis() - offer.getStart() >= 2000) {
                        offers.remove(k);
                    }
                });
            }
        }, 1, 1, TimeUnit.SECONDS);
    }

    public void add(Offer request) {
        offers.put(request.getUser(), request);
    }

    public void onAnswer(String user, String sdp) throws Exception {
        Offer offer = offers.get(user);
        if (offer != null) {
            offer.setResponse(sdp);
        }
    }
}

package me.rtmsoft.webrtc.signal;

import java.security.Principal;

public class Peer implements Principal {

    public static final String OFFER_SUFFIX = "_offer";
    public static final String ANSWER_SUFFIX = "_answer";

    private final String token;

    private final String type;

    private final String sessionId;

    public Peer(String token, String type, String sessionId) {
        this.token = token;
        this.type = type;
        this.sessionId = sessionId;
    }

    public String getToken() {
        return token;
    }

    public String getType() {
        return type;
    }

    public String getSessionId() {
        return sessionId;
    }

    @Override
    public String getName() {
        return token + "_" + type;
    }
}

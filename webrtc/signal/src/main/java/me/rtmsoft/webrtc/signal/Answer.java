package me.rtmsoft.webrtc.signal;

import java.security.Principal;

public class Answer implements Principal {

    private final String name;

    private String token;

    public Answer(String name) {
        this.name = name;
    }

    @Override
    public String getName() {
        return name;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    @Override
    public String toString() {
        return "Answer{" +
                "name='" + name + '\'' +
                ", token='" + token + '\'' +
                '}';
    }
}

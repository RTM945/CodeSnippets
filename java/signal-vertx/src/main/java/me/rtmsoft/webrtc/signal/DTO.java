package me.rtmsoft.webrtc.signal;

import com.fasterxml.jackson.annotation.JsonInclude;

public class DTO {

    private int evt;

    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private String token;

    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private String sdp;

    public int getEvt() {
        return evt;
    }

    public void setEvt(int evt) {
        this.evt = evt;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public String getSdp() {
        return sdp;
    }

    public void setSdp(String sdp) {
        this.sdp = sdp;
    }

    public DTO() {
    }

    public DTO(int evt, String token, String sdp) {
        this.evt = evt;
        this.token = token;
        this.sdp = sdp;
    }

    @Override
    public String toString() {
        return "DTO{" +
                "evt=" + evt +
                ", token='" + token + '\'' +
                ", sdp='" + sdp + '\'' +
                '}';
    }
}

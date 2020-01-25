package me.rtmsoft.webrtc.signal;

import com.fasterxml.jackson.annotation.JsonInclude;

public class Dto {

    private int evt;

    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private String token;

    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private String sdp;

    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private String msg;

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

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public Dto() {
    }

    public Dto(int evt, String token, String sdp, String msg) {
        this.evt = evt;
        this.token = token;
        this.sdp = sdp;
        this.msg = msg;
    }

    @Override
    public String toString() {
        return "DTO{" +
                "evt=" + evt +
                ", token='" + token + '\'' +
                ", sdp='" + sdp + '\'' +
                ", msg='" + msg + '\'' +
                '}';
    }
}

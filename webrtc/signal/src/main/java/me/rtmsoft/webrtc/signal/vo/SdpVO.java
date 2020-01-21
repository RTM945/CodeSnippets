package me.rtmsoft.webrtc.signal.vo;

public class SdpVO {

    private String sdp;

    public String getSdp() {
        return sdp;
    }

    public void setSdp(String sdp) {
        this.sdp = sdp;
    }

    @Override
    public String toString() {
        return "SdpVO{" +
                "sdp='" + sdp + '\'' +
                '}';
    }
}

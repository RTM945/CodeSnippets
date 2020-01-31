package me.rtmsoft.webrtc.signal;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Controller;

import java.security.Principal;

@Controller
public class SignalController {

    @Autowired
    private PeerManager peerManager;

    @Autowired
    private SimpMessagingTemplate messagingTemplate;

    @MessageMapping("/sdp")
    public void sdp(@Payload String sdp, Principal principal) {
        Peer peer = (Peer) principal;
        Peer pair = peerManager.pair(peer);
        if(pair != null) {
            messagingTemplate.convertAndSendToUser(pair.getName(), "/queue/onsdp", sdp);
        }
    }

    @MessageMapping("/candidate")
    public void candidate(@Payload String candidate, Principal principal) {
        Peer peer = (Peer) principal;
        Peer pair = peerManager.pair(peer);
        if(pair != null) {
            messagingTemplate.convertAndSendToUser(pair.getName(), "/queue/oncandidate", candidate);
        }
    }
}

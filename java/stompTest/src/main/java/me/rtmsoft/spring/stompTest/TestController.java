package me.rtmsoft.spring.stompTest;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.messaging.simp.annotation.SendToUser;
import org.springframework.stereotype.Controller;

import java.security.Principal;

@Controller
public class TestController {

    private static final Logger LOGGER = LoggerFactory.getLogger(TestController.class);

    @Autowired
    private SimpMessagingTemplate messagingTemplate;

    @MessageMapping("/say")
//    @SendToUser
    public void echo(@Payload String msg, Principal principal) {
        LOGGER.info("{} say {}", principal.getName(), msg);
//        return msg;
        messagingTemplate.convertAndSendToUser(principal.getName(), "/queue/echo", msg);
    }
}

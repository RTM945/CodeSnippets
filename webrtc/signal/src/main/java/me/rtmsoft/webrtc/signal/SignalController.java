package me.rtmsoft.webrtc.signal;

import me.rtmsoft.webrtc.signal.vo.PairVO;
import me.rtmsoft.webrtc.signal.vo.SdpVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.messaging.handler.annotation.DestinationVariable;
import org.springframework.messaging.handler.annotation.MessageExceptionHandler;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.messaging.simp.annotation.SendToUser;
import org.springframework.stereotype.Controller;
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpSession;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.TimeUnit;

@Controller
public class SignalController {

    Map<String, Answer> answers = new ConcurrentHashMap<>();

    @Autowired
    private SimpMessagingTemplate messagingTemplate;

    @Autowired
    private OfferManager offerManager;

    //answer端注册
    @MessageMapping("/register/{token}")
    public void register(@DestinationVariable String token, Answer answer) throws Exception {
        if(answer == null || StringUtils.isEmpty(answer.getName())) {
            throw new Exception("name can't be empty");
        }
        if (answers.get(answer.getName()) != null) {
            throw new Exception("duplicate name");
        }else{
            answer.setToken(token);
            answers.put(answer.getName(), answer);
        }
    }

    @MessageMapping("/answer")
    public void answer(@Payload String sdp, Answer answer) throws Exception {
        if (!StringUtils.isEmpty(sdp)) {
            offerManager.onAnswer(answer.getName(), sdp);
        }
    }

    @MessageExceptionHandler
    @SendToUser("/user/errors")
    public String handleException(Throwable exception) {
        return exception.getMessage();
    }

    //将offer和answer配对
    @RequestMapping(path = "/pair", method = RequestMethod.POST, consumes = "application/json", produces = "application/json")
    public ResponseEntity<String> pair(@RequestBody PairVO pairVO, HttpSession session) {
        Answer answer = answers.get(pairVO.getName());
        if (answer != null && answer.getToken().equals(pairVO.getToken())) {
            //配对成功写入session
            session.setAttribute("user", pairVO.getName());
            return ResponseEntity.status(HttpStatus.OK).build();
        }
        return ResponseEntity.status(HttpStatus.I_AM_A_TEAPOT).build();
    }

    @RequestMapping(path = "/offer", method = RequestMethod.POST, consumes = "application/json", produces = "application/json")
    @ResponseBody
    public ResponseEntity<String> offer(@RequestBody SdpVO sdpVO, HttpSession session) {
        String user = (String) session.getAttribute("user");
        String answer = "";
        HttpStatus code = HttpStatus.OK;
        if (StringUtils.isEmpty(user)) {
            code = HttpStatus.UNAUTHORIZED;
        } else {
            //发送给answer
            Offer offer = new Offer(user);
            offerManager.add(offer);
            messagingTemplate.convertAndSendToUser(user, "/queue/offer", sdpVO);
            //等待answer的sdp
            try{
                answer = offer.getAnswer(2, TimeUnit.SECONDS);
                if (StringUtils.isEmpty(answer)) {
                    code = HttpStatus.SERVICE_UNAVAILABLE;
                }
            } catch (Exception e) {
                code = HttpStatus.BAD_REQUEST;
            }
        }

        if (StringUtils.isEmpty(user)) {
            code = HttpStatus.UNAUTHORIZED;
        }
        return ResponseEntity.status(code).body(answer);
    }
}

package me.rtmsoft.webrtc.signal;

import io.netty.util.internal.StringUtil;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Promise;
import io.vertx.core.http.ServerWebSocket;
import io.vertx.core.json.Json;
import io.vertx.core.logging.Logger;
import io.vertx.core.logging.LoggerFactory;
import io.vertx.ext.web.Router;
import io.vertx.ext.web.RoutingContext;
import io.vertx.ext.web.handler.BodyHandler;
import io.vertx.ext.web.handler.SessionHandler;
import io.vertx.ext.web.sstore.SessionStore;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class Server extends AbstractVerticle {

    private static final Logger LOGGER = LoggerFactory.getLogger(Server.class);

    private static final Map<String, Client> CLIENTS = new ConcurrentHashMap<>();

    public static final int OFFER = 0;
    public static final int ANSWER = 1;
    public static final int ERROR = 2;

    @Override
    public void start(Promise<Void> startPromise) {
        SessionStore sessionStore = SessionStore.create(vertx);
        Router root = Router.router(vertx);
        root.route().handler(SessionHandler.create(sessionStore));
        root.route().handler(BodyHandler.create(false));
        root.route("/signalling").handler(this::handleSignal);
        root.post("/offer").handler(this::offer);
        root.errorHandler(500, rc -> {
            System.err.println("Handling failure");
            Throwable failure = rc.failure();
            if (failure != null) {
                failure.printStackTrace();
            }
        });
        vertx.createHttpServer()
                .requestHandler(root)
                .listen(9090, ar -> {
                   LOGGER.info("signalling starts at 9090");
                    startPromise.complete();
                });
    }

    private void offer(RoutingContext context) {
        Dto dto = Json.decodeValue(context.getBodyAsString(), Dto.class);
        Client client = CLIENTS.get(dto.getToken());
        if (client != null) {
            dto.setEvt(OFFER);
            client.send(dto);
            try {
                dto = new Dto();
                String answer = client.getAnswer();
                if (StringUtil.isNullOrEmpty(answer)) {
                    dto.setEvt(ERROR);
                    dto.setMsg("can't pair");
                }else {
                    dto.setEvt(ANSWER);
                    dto.setSdp(answer);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
            context.response()
                    .putHeader("content-type", "application/json")
                    .end(Json.encode(dto));
        }
    }

    private void handleSignal(RoutingContext context) {
        ServerWebSocket ws = context.request().upgrade();
        final String token = ws.headers().get("token");
        if (StringUtil.isNullOrEmpty(token)) {
            ws.close((short) 1001, "no token header");
            return;
        }
        if (CLIENTS.get(token) != null) {
            ws.close((short) 1002, "duplicate token");
            return;
        }
        LOGGER.info("register " + token);
        final Client client = new Client(token, ws);
        CLIENTS.put(token, client);
        ws.textMessageHandler(msg -> {
            LOGGER.info("receive answer");
            Dto dto = Json.decodeValue(msg, Dto.class);
            LOGGER.info(dto);
            if (dto.getEvt() == ANSWER) {
                client.setAnswer(dto.getSdp());
            }
        });
        ws.closeHandler(event -> CLIENTS.remove(token));
    }
}

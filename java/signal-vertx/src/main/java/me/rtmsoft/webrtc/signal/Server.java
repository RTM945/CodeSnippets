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

    private static final Map<String, ServerWebSocket> CLIENTS = new ConcurrentHashMap<>();

    public static final int OFFER = 0;
    public static final int ANSWER = 1;

    @Override
    public void start(Promise<Void> startPromise) {
        SessionStore sessionStore = SessionStore.create(vertx);
        Router root = Router.router(vertx);
        root.route().handler(SessionHandler.create(sessionStore));
        root.route().handler(BodyHandler.create(false));
        root.route("/signalling").handler(this::handleSignal);
        root.post("/offer").handler(this::offer);
        root.errorHandler(500, rc -> {
            LOGGER.error("Handling failure");
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
        DTO dto = Json.decodeValue(context.getBodyAsString(), DTO.class);
        ServerWebSocket ws = CLIENTS.get(dto.getToken());
        if (ws != null) {
            dto.setEvt(OFFER);
            ws.writeTextMessage(Json.encode(dto)).textMessageHandler(msg -> {
                DTO answer = Json.decodeValue(msg, DTO.class);
                context.response()
                        .putHeader("content-type", "application/json")
                        .end(Json.encode(answer));
            });
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
        CLIENTS.put(token, ws);
        ws.closeHandler(event -> CLIENTS.remove(token));
    }
}

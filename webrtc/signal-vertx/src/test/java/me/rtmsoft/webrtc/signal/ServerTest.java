package me.rtmsoft.webrtc.signal;

import io.vertx.core.Vertx;
import io.vertx.core.buffer.Buffer;
import io.vertx.core.http.*;
import io.vertx.core.json.Json;
import io.vertx.core.logging.Logger;
import io.vertx.core.logging.LoggerFactory;
import io.vertx.ext.unit.Async;
import io.vertx.ext.unit.TestContext;
import io.vertx.ext.unit.junit.VertxUnitRunner;
import io.vertx.ext.web.client.HttpResponse;
import io.vertx.ext.web.client.WebClient;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;

import java.io.IOException;

@RunWith(VertxUnitRunner.class)
public class ServerTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(ServerTest.class);

    private Vertx vertx;

    @Before
    public void setup(TestContext testContext) throws IOException {
        vertx = Vertx.vertx();
        vertx.deployVerticle(Server.class.getName(), testContext.asyncAssertSuccess());
    }

    @After
    public void tearDown(TestContext testContext) {
        vertx.close(testContext.asyncAssertSuccess());
    }

    @Test
    public void testAll(TestContext testContext) {
        final Async async = testContext.async();

        HttpClient client = vertx.createHttpClient();
        WebSocketConnectOptions wsOps = new WebSocketConnectOptions();
        wsOps.setHost("localhost").setPort(9090).setURI("/signalling").addHeader("token", "test");
        client.webSocket(wsOps, ar -> {
            if(ar.succeeded()) {
                WebSocket ws = ar.result();
                ws.textMessageHandler(msg -> {
                    Dto dto = Json.decodeValue(msg, Dto.class);
                    LOGGER.info(dto);
                    if (dto.getEvt() == Server.OFFER) {
                        dto.setEvt(Server.ANSWER);
                        dto.setSdp("test answer");
                        ws.writeTextMessage(Json.encode(dto));
                        LOGGER.info("send answer");
                    }
                });
            }else {
                ar.cause().printStackTrace();
            }
        });

        WebClient offer = WebClient.create(vertx);
        Dto dto = new Dto();
        dto.setEvt(Server.OFFER);
        dto.setToken("test");
        dto.setSdp("test offer");
        offer.post(9090, "localhost", "/offer").sendJson(dto, ar -> {
            if (ar.succeeded()) {
                HttpResponse<Buffer> response = ar.result();
                LOGGER.info(response.body().toString());
                async.complete();
            } else {
                ar.cause().printStackTrace();
            }
        });
    }
}

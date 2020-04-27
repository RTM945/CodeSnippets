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
import java.util.concurrent.TimeUnit;

@RunWith(VertxUnitRunner.class)
public class ServerTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(ServerTest.class);

    private Vertx vertx;

    private String token = "test";

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
    public void testAll(TestContext testContext) throws InterruptedException {
        final Async async = testContext.async();

        HttpClient client = vertx.createHttpClient();
        WebSocketConnectOptions wsOps = new WebSocketConnectOptions();
        wsOps.setHost("localhost").setPort(9090).setURI("/signalling").addHeader("token", token);
        client.webSocket(wsOps, ar -> {
            if(ar.succeeded()) {
                WebSocket ws = ar.result();
                ws.textMessageHandler(msg -> {
                    DTO dto = Json.decodeValue(msg, DTO.class);
                    if (dto.getEvt() == Server.OFFER) {
                        dto.setEvt(Server.ANSWER);
                        dto.setSdp("test answer");
                        ws.writeTextMessage(Json.encode(dto));
                    }
                });
            }else {
                ar.cause().printStackTrace();
            }
        });
        TimeUnit.SECONDS.sleep(1);
        WebClient offer = WebClient.create(vertx);
        DTO dto = new DTO();
        dto.setEvt(Server.OFFER);
        dto.setToken("test");
        dto.setSdp("test offer");
        LOGGER.info("offer ----> " + dto);
        offer.post(9090, "localhost", "/offer").sendJson(dto, ar -> {
            if (ar.succeeded()) {
                HttpResponse<Buffer> response = ar.result();
                LOGGER.info("<---- answer " + Json.decodeValue(response.body(), DTO.class));
                async.complete();
            } else {
                ar.cause().printStackTrace();
            }
        });
    }
}

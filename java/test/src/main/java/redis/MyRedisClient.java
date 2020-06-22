package redis;

import io.lettuce.core.RedisClient;
import io.lettuce.core.RedisURI;
import io.lettuce.core.api.StatefulRedisConnection;
import io.lettuce.core.codec.RedisCodec;

public class MyRedisClient {

    private static final String HOST = "localhost";

    private static final int PORT = 6379;

    private static final RedisClient CLIENT = RedisClient.create(RedisURI.create(HOST, PORT));

    private static final RedisCodec<String, Object> CODEC = new SerializedObjectCodec();

    public static StatefulRedisConnection<String, Object> getConnection() {
        return CLIENT.connect(CODEC);
    }

    public static <K, V> StatefulRedisConnection<K, V> getConnection(RedisCodec<K, V> codec) {
        return CLIENT.connect(codec);
    }
}

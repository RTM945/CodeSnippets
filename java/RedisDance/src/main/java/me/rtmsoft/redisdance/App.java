package me.rtmsoft.redisdance;

import io.lettuce.core.RedisClient;
import io.lettuce.core.api.StatefulRedisConnection;
import io.lettuce.core.api.sync.RedisCommands;

/**
 * Hello world!
 */
public class App {
    public static void main(String[] args) {
        RedisClient redisClient = RedisClient.create("redis://localhost:6379/0");
        StatefulRedisConnection<String, String> connection = redisClient.connect();
        RedisCommands<String, String> syncCommands = connection.sync();

        syncCommands.set("key", "Hello, Redis!");
        System.out.println(syncCommands.get("key"));
        syncCommands.del("key");
        System.out.println(syncCommands.get("key"));
        connection.close();
        redisClient.shutdown();
    }
}

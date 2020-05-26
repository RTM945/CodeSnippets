package me.rtmsoft.redisdance;

import io.lettuce.core.SetArgs;
import io.lettuce.core.api.StatefulRedisConnection;
import io.lettuce.core.api.sync.RedisCommands;
import io.lettuce.core.internal.LettuceAssert;

public class RedisOps {

    static void set(String key, String value, SetArgs setArgs) {
        try (StatefulRedisConnection<String, Object> conn = MyRedisClient.getConnection()) {
            RedisCommands<String, Object> commands = conn.sync();
            String result;
            if (setArgs == null) {
                result = commands.set(key, value);
            } else {
                result = commands.set(key, value, setArgs);
            }
            LettuceAssert.isTrue("OK".equalsIgnoreCase(result), "fail");
        }
    }

    static void set(String key, String value) {
        set(key, value, null);
    }

    static void setex(String key, String value, long timeout) {
        SetArgs setArgs = SetArgs.Builder.ex(timeout);
        set(key, value, setArgs);
    }

    static void setnx(String key, String value) {
        SetArgs setArgs = SetArgs.Builder.nx();
        set(key, value, setArgs);
    }

    static void setnxex(String key, String value, long timeout) {
        SetArgs setArgs = SetArgs.Builder.nx().ex(timeout);
        set(key, value, setArgs);
    }

    static Object get(String key) {
        try (StatefulRedisConnection<String, Object> conn = MyRedisClient.getConnection()) {
            RedisCommands<String, Object> commands = conn.sync();
            return commands.get(key);
        }
    }
}
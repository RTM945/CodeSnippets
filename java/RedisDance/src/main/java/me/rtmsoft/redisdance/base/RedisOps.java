package me.rtmsoft.redisdance.base;

import io.lettuce.core.ScriptOutputType;
import io.lettuce.core.SetArgs;
import io.lettuce.core.api.StatefulRedisConnection;
import io.lettuce.core.api.sync.RedisCommands;
import io.lettuce.core.codec.StringCodec;
import io.lettuce.core.internal.LettuceAssert;

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicReference;
import java.util.function.Function;

public class RedisOps {

    private static void set(String key, String value, SetArgs setArgs) {
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

    public static void set(String key, String value) {
        set(key, value, null);
    }

    public static void setex(String key, String value, long timeout) {
        SetArgs setArgs = SetArgs.Builder.ex(timeout);
        set(key, value, setArgs);
    }

    public static void setnx(String key, String value) {
        SetArgs setArgs = SetArgs.Builder.nx();
        set(key, value, setArgs);
    }

    public static void setnxex(String key, String value, long timeout) {
        SetArgs setArgs = SetArgs.Builder.nx().ex(timeout);
        set(key, value, setArgs);
    }

    public static Object get(String key) {
        try (StatefulRedisConnection<String, Object> conn = MyRedisClient.getConnection()) {
            RedisCommands<String, Object> commands = conn.sync();
            return commands.get(key);
        }
    }

    public static long del(String... key) {
        try (StatefulRedisConnection<String, Object> conn = MyRedisClient.getConnection()) {
            RedisCommands<String, Object> commands = conn.sync();
            return commands.del(key);
        }
    }

    private static final ConcurrentHashMap<String, String> scriptLoaded = new ConcurrentHashMap<>();

    public static Object evalsha(String script, ScriptOutputType type, String[] keys, String... args) {
        // must use StringCodec.UTF8 load lua script
        scriptLoaded.computeIfAbsent(script, s -> {
            try (StatefulRedisConnection<String, String> conn = MyRedisClient.getConnection(StringCodec.UTF8)) {
                RedisCommands<String, String> commands = conn.sync();
                return commands.scriptLoad(script);
            }
        });

        try (StatefulRedisConnection<String, Object> conn = MyRedisClient.getConnection()) {
            RedisCommands<String, Object> commands = conn.sync();
            return commands.evalsha(scriptLoaded.get(script), type, keys, args);
        }
    }
}
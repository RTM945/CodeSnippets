package me.rtmsoft.redisdance;

import me.rtmsoft.redisdance.base.RedisOps;

import java.time.Duration;
import java.time.temporal.ChronoUnit;
import java.util.Random;

public class VerificationCode {

    private static final String VERIFICATION_KEY_PREFIX = "VERIFICATION_KEY_PREFIX_";

    public static void gen(String customKey) throws Exception {
        int code = new Random().nextInt(10000);
        RedisOps.setex(VERIFICATION_KEY_PREFIX + customKey, String.valueOf(code), Duration.of(5, ChronoUnit.MINUTES).toMillis());
    }

    public static boolean verify(String customKey, String code) throws Exception {
        String key = VERIFICATION_KEY_PREFIX + customKey;
        if (code.equals(RedisOps.get(key))) {
            RedisOps.del(key);
            return true;
        }
        return false;
    }

}

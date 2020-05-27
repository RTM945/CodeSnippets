package me.rtmsoft.redisdance;

import io.lettuce.core.ScriptOutputType;
import me.rtmsoft.redisdance.base.LuaScript;
import me.rtmsoft.redisdance.base.RedisOps;

import java.util.UUID;
import java.util.concurrent.TimeUnit;

public class RedLock {

    private final String clientId = UUID.randomUUID().toString();

    private static final String KEY_PREFIX = "redlock_";

    private final LuaScript unlockScript = new LuaScript("unlock", "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end");

    private String lockKey;
    private long expireAfter;

    public RedLock(String lockKey, long expireAfter) {
        this.lockKey = lockKey;
        this.expireAfter = expireAfter;
    }

    public boolean tryLock(long time, TimeUnit timeUnit) {
        long now = System.currentTimeMillis();
        try {
            long expire = now + TimeUnit.MILLISECONDS.convert(time, timeUnit);
            boolean acquired;
            while (!(acquired = obtainLock()) && System.currentTimeMillis() < expire) { //NOSONAR
                Thread.sleep(100); //NOSONAR
            }
            return acquired;
        } catch (Exception e) {
//            e.printStackTrace();
        }
        return false;
    }

    public boolean unlock() {
        String[] keys = {KEY_PREFIX + lockKey};
        Object result = RedisOps.evalsha(unlockScript, ScriptOutputType.INTEGER, keys, clientId);
        return "1".equals(String.valueOf(result));
    }

    private boolean obtainLock() {
        try {
            RedisOps.setnxex(KEY_PREFIX + lockKey, clientId, expireAfter);
            return true;
        } catch (Exception ignore) {
            return false;
        }
    }
}

package redis;

import io.lettuce.core.ScriptOutputType;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

/**
 * 使用redis实现分布式锁
 */
public class RedisLock {

    private final String clientId = UUID.randomUUID().toString();

    //key前缀
    private static final String KEY_PREFIX = "redlock_";

    //去锁的lua脚本
    //只有匹配key才能把锁记录删除
    private final LuaScript unlockScript = new LuaScript("unlock", "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end");

    //一个锁的key
    //只有匹配key才能解锁
    private String lockKey;

    //锁过期时间
    private long expireAfter;

    public RedisLock(String lockKey, long expireAfter) {
        this.lockKey = lockKey;
        this.expireAfter = expireAfter;
    }

    //阻塞的上锁操作
    public void lock() {
        while (true) {
            try {
                //循环执行直到获取锁
                while (!obtainLock()) {
                    Thread.sleep(100);
                }
            } catch (InterruptedException ignore) {

            }
            break;
        }
    }

    /**
     * 尝试获取锁
     * @param time 尝试时间
     * @param timeUnit 单位
     * @return true 获取成功 false获取失败
     */
    public boolean tryLock(long time, TimeUnit timeUnit) {
        long now = System.currentTimeMillis();
        try {
            long expire = now + TimeUnit.MILLISECONDS.convert(time, timeUnit);
            boolean acquired;
            // 加上间隔时间防止消耗太多cpu资源
            while (!(acquired = obtainLock()) && System.currentTimeMillis() < expire) { //NOSONAR
                Thread.sleep(100); //NOSONAR
            }
            return acquired;
        } catch (Exception e) {
//            e.printStackTrace();
        }
        return false;
    }

    //解锁操作
    public boolean unlock() {
        String[] keys = {KEY_PREFIX + lockKey};
        Object result = RedisOps.evalsha(unlockScript, ScriptOutputType.INTEGER, keys, clientId);
        return "1".equals(String.valueOf(result));
    }

    //获得锁操作
    private boolean obtainLock() {
        try {
            // SET key value NX EX expire
            RedisOps.setnxex(KEY_PREFIX + lockKey, clientId, expireAfter);
            return true;
        } catch (Exception ignore) {
            return false;
        }
    }
}

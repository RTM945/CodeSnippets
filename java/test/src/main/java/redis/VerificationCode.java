package redis;

import java.time.Duration;
import java.time.temporal.ChronoUnit;
import java.util.concurrent.ThreadLocalRandom;

/**
 * 用redis实现验证码操作
 */
public class VerificationCode {

    //验证码前缀
    private static final String VERIFICATION_KEY_PREFIX = "VERIFICATION_KEY_PREFIX_";

    static int min = 10000;
    static int max = 99999;

    //过期时间
    static long timeout = Duration.of(5, ChronoUnit.MINUTES).toMillis();

    //生成验证码
    public static String gen(String customKey) {
        String code = String.valueOf(ThreadLocalRandom.current().nextInt(min, max + 1));
        try{
            //SET key value EX timeout
            RedisOps.setex(VERIFICATION_KEY_PREFIX + customKey, code, timeout);
            return code;
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }

    //判断验证码是否正确
    public static boolean verify(String customKey, String code) {
        String key = VERIFICATION_KEY_PREFIX + customKey;
        if (code.equals(RedisOps.get(key))) {
            //RedisOps.del(key);
            return true;
        }
        return false;
    }

}

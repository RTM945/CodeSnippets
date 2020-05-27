package me.rtmsoft.redisdance;

import io.lettuce.core.LettuceStrings;
import io.lettuce.core.RedisClient;
import io.lettuce.core.RedisURI;
import io.lettuce.core.ScriptOutputType;
import io.lettuce.core.api.StatefulRedisConnection;
import io.lettuce.core.api.sync.RedisCommands;
import io.lettuce.core.codec.RedisCodec;
import io.lettuce.core.output.StatusOutput;
import io.lettuce.core.protocol.Command;
import io.lettuce.core.protocol.CommandArgs;
import io.lettuce.core.protocol.CommandType;
import io.lettuce.core.protocol.ProtocolKeyword;
import me.rtmsoft.redisdance.base.RedisOps;
import me.rtmsoft.redisdance.base.SerializedObjectCodec;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

public class RedLockTest {

    private final String unlockScript = "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end";

    @Test
    public void testScript() throws Exception{
        RedisOps.set("xixi", "haha");
        Object result = RedisOps.evalsha(unlockScript, ScriptOutputType.INTEGER, new String[]{"xixi"}, "haha");
        System.out.println(result);
        System.out.println(RedisOps.get("xixi"));
    }

    Runnable tryLockTester(String name, List<Integer> list) {
        return () -> {
            for (int i = 0; i < 5; i++) {
                RedLock lock = new RedLock("test", 5000);
                boolean acquired = false;
                try {
                    acquired = lock.tryLock(1, TimeUnit.SECONDS);
                    if (acquired) {
                        System.out.println(name + " get lock");
                        list.add(1);
                        Thread.sleep(1000);
                    }
                }catch (Exception e) {
                    e.printStackTrace();
                } finally {
                    if(acquired){
                        lock.unlock();
                        System.out.println(name + " unlock");
                    }
                }
            }
        };
    }

    @Test
    public void testConcurrency() throws Exception{
        List<Integer> list = new ArrayList<>();
        Runnable a = tryLockTester("a", list);
        Runnable b = tryLockTester("b", list);
        new Thread(a).start();
        new Thread(b).start();
        Thread.sleep(10000);
        System.out.println(list.size());
    }

}

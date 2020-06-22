package redis;

import io.lettuce.core.ScriptOutputType;
import org.junit.Test;

import java.util.Random;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public class RedLockTest {

    @Test
    public void testScript() throws Exception{
        RedisOps.set("xixi", "haha");
        LuaScript unlockScript = new LuaScript("unlock", "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end");
        Object result = RedisOps.evalsha(unlockScript, ScriptOutputType.INTEGER, new String[]{"xixi"}, "haha");
        System.out.println(result);
        System.out.println(RedisOps.get("xixi"));
    }

    Runnable tryLockTester(String name, CountDownLatch count) {
        return () -> {
            while (true){
                if(count.getCount() <= 0) break;
                RedisLock lock = new RedisLock("test", 100);
                boolean acquired = false;
                try {
                    acquired = lock.tryLock(1, TimeUnit.SECONDS);
                    if (acquired) {
                        System.out.println(name + " get lock");
                        Thread.sleep(new Random().nextInt(1000));
                        count.countDown();
                    }
                }catch (Exception e) {
                    e.printStackTrace();
                } finally {
                    if(acquired){
                        boolean unlock = lock.unlock();
                        System.out.println(name + " unlock " + unlock);
                    }
                }

                try {
                    Thread.sleep(new Random().nextInt(1000));
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        };
    }

    @Test
    public void testTryLock() throws Exception{
        RedisOps.del("redlock_test");
        CountDownLatch count = new CountDownLatch(10);
        Runnable a = tryLockTester("a", count);
        Runnable b = tryLockTester("b", count);
        new Thread(a).start();
        new Thread(b).start();
        count.await();
        Thread.sleep(1000);
    }

    Runnable lockTester(String name, CountDownLatch latch, int[] container) {
        return () -> {
            for (int i = 0; i < 50; i++) {
                RedisLock lock = new RedisLock("test", 100);
                lock.lock();
                System.out.println(name + " get lock");
                container[0] += 1;
                lock.unlock();
                System.out.println(name + " unlock");
                latch.countDown();
                Thread.yield();
            }
        };
    }

    @Test
    public void lockTest() throws Exception{
        RedisOps.del("redlock_test");
        CountDownLatch count = new CountDownLatch(100);
        int[] container = {0};
        Runnable a = lockTester("a", count, container);
        Runnable b = lockTester("b", count, container);
        new Thread(a).start();
        new Thread(b).start();
        count.await();
        System.out.println(container[0]);
    }
}

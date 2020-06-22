package thread;

import org.junit.Test;

import java.util.concurrent.CountDownLatch;

public class WorkerQueueTest {

    @Test
    public void test() throws Exception {
        WorkerQueue queue = new WorkerQueue(8);
        Thread.sleep(1000);
        for (int i = 0; i < 20; i++) {
            queue.put(i);
        }
        Thread.sleep(10000);
    }

}

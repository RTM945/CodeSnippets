package thread;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

/**
 * https://tech.meituan.com/2020/04/02/java-pooling-pratice-in-meituan.html
 */
public class ThreadPoolExecutorTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(ThreadPoolExecutorTest.class);

    public static ThreadPoolExecutor defaultExec() {
        return new ThreadPoolExecutor(2, 5,
                60, TimeUnit.SECONDS,
                new LinkedBlockingQueue<>(10),
                new NamedThreadFactory("test"),
                new ThreadPoolExecutor.AbortPolicy());
    }

    public static void testThreadPoolExecutor() throws Exception {
        ThreadPoolExecutor exec = defaultExec();
        for (int i = 0; i < 15; i++) {
            int id = i;
            exec.execute(() -> {
                // 通过id顺序可以发现
                // 当任务数大于corePoolSize时
                // 先存入queue
                // 当任务数大于queueSize小于queueSize+maxPoolSize时
                // 超过queueSize的任务直接创建线程
                // 这就会出现后加入的任务先执行的现象
                threadPoolStatus(exec, "创建任务" + id);
                try {
                    TimeUnit.SECONDS.sleep(10);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            });
        }
        Thread.sleep(1000);
        threadPoolStatus(exec, "改变之前");
        exec.setCorePoolSize(10);
        exec.setMaximumPoolSize(10);
        Thread.sleep(1000);
        threadPoolStatus(exec, "改变之后");
        //shutdown停止接收任务，会等待队列中任务完成
        //shutdownNow停止接收任务，但会中断所有执行中的任务
        exec.shutdown();
        while (!exec.awaitTermination(1, TimeUnit.SECONDS)) {
            LOGGER.info("执行中");
        }
        threadPoolStatus(exec, "结束");
    }

    static void threadPoolStatus(ThreadPoolExecutor exec, String tag) {
        BlockingQueue<Runnable> queue = exec.getQueue();
        LOGGER.info("{}-{}-:核心线程数:{} " +
                        "活动线程数:{} " +
                        "最大线程数:{} " +
                        "任务完成数:{} " +
                        "队列大小:{} " +
                        "排队线程数:{} " +
                        "队列剩余大小:{}",
                Thread.currentThread().getName(),
                tag,
                exec.getCorePoolSize(),
                exec.getActiveCount(),
                exec.getMaximumPoolSize(),
                exec.getCompletedTaskCount(),
                queue.size() + queue.remainingCapacity(),
                queue.size(),
                queue.remainingCapacity());
    }

    public static void main(String[] args) throws Exception{
        // 该测试可以看出
        // 如果任务数大于corePoolSize小于等于maxPoolSize
        // 或创建所有任务
        // 如果任务数大于maxPoolSize
        // 会创建maxPoolSize个任务，剩下的放入queue等待
        // 如果任务数超过queueSize+maxPoolSize
        // 执行RejectedExecutionHandler
        testThreadPoolExecutor();
    }

}

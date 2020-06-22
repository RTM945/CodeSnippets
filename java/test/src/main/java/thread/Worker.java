package thread;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingQueue;

public class Worker implements Runnable {

    private Logger logger;

    private final ExecutorService executorService;
    Integer id;
    protected BlockingQueue<Integer> queue;

    public Worker(int id) {
        this.id = id;
        NamedThreadFactory factory = new NamedThreadFactory("worker-" + id);
        this.executorService = Executors.newSingleThreadExecutor(factory);
        queue = new LinkedBlockingQueue<>();
        logger = LoggerFactory.getLogger("worker-" + id);
    }

    @Override
    public void run() {
        logger.info("worker-{} starts work", id);
        while (true) {
            try {
                Integer work = queue.take();
                logger.info("worker-{} starts to deal with work-{}", id, work);
                Thread.sleep(2000);
                logger.info("worker-{} finished work-{}", id, work);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public void startWork() {
        executorService.execute(this);
    }

    public void dispatch(Integer work) {
        try {
            queue.put(work);
            int size = queue.size();
            logger.info("worker-{} has {} works", id, size);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}

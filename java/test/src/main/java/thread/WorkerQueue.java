package thread;

import java.util.Map;
import java.util.concurrent.*;

public class WorkerQueue {

    int workerCount;

    Map<Integer, Worker> workers = new ConcurrentHashMap<>();

    public WorkerQueue(int workerCount) {
        this.workerCount = workerCount;
        for (int i = 0; i < workerCount; i++) {
            Worker worker = new Worker(i);
            workers.put(i, worker);
            worker.startWork();
        }
    }

    public void put(Integer work) {
        int index = work % workerCount;
        Worker worker = workers.get(index);
        if(worker != null) {
            worker.dispatch(work);
        }
    }
}

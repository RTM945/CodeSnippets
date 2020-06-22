package thread;

import java.util.concurrent.ThreadFactory;
import java.util.concurrent.atomic.AtomicInteger;

public class NamedThreadFactory implements ThreadFactory {

    private final ThreadGroup group;
    private final AtomicInteger threadNumber = new AtomicInteger(0);
    private final String namePrefix;
    private final boolean daemon;

    @Override
    public Thread newThread(Runnable r) {
        String name = namePrefix + "-thread-" + threadNumber.getAndIncrement();
        Thread t = new Thread(group, r, name, 0);
        if(daemon) {
            t.setDaemon(true);
        }else{
            if (t.isDaemon())
                t.setDaemon(false);
            if (t.getPriority() != Thread.NORM_PRIORITY)
                t.setPriority(Thread.NORM_PRIORITY);
        }
        return t;
    }

    public NamedThreadFactory(String namePrefix) {
        this(namePrefix, false);
    }

    public NamedThreadFactory(String namePrefix, boolean daemon) {
        SecurityManager s = System.getSecurityManager();
        group = (s != null) ? s.getThreadGroup() : Thread.currentThread().getThreadGroup();
        this.namePrefix = namePrefix;
        this.daemon = daemon;
    }
}

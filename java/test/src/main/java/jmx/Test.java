package jmx;

import java.util.concurrent.TimeUnit;

public class Test {

    public static void main(String[] args) throws Exception {
        MBeanManager.registerMBean(new Reload(), "test_reload");
        // for test
        TimeUnit.SECONDS.sleep(60);
        // use JConsole -> MBean -> reload
    }

}

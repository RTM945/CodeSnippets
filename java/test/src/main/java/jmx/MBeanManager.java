package jmx;

import javax.management.MBeanServer;
import javax.management.ObjectName;
import java.lang.management.ManagementFactory;

public class MBeanManager {
    private static final MBeanServer mbs = ManagementFactory.getPlatformMBeanServer();

    public static void registerMBean(Object obj, String name) throws Exception {
        ObjectName objectName = new ObjectName("mbean:name=" + name);
        mbs.registerMBean(obj, objectName);
        System.out.println(mbs.getMBeanInfo(objectName));
        System.out.println("mbean count:" + mbs.getMBeanCount());
    }
}

package simple;

//javac SyncComplier.java     
//javap -p -v SyncComplier.class
public class SyncComplier {

    Object lock = new Object();

    int a;

    //ACC_SYNCHRONIZED
    public synchronized void incr1() {
        a++;
    }

    public void incr2() {
        //monitorenter
        synchronized(lock) {
            a++;
        }
        //monitorexit
    }    
}
package nio.reactor;

import java.nio.channels.SocketChannel;

public class EventLoopGroup {

    private EventLoop[] eventLoops;
    private int index;
    private int size;

    public EventLoopGroup(int size) throws Exception {
        this.size = size;
        if(size > 0) {
            eventLoops = new EventLoop[size];
            for (int i = 0; i < eventLoops.length; i++) {
                EventLoop eventLoop = new EventLoop();
                eventLoops[i] = eventLoop;
                new Thread(eventLoop).start();
            }
        }
    }
    
    public void dispatch(SocketChannel socketChannel) throws Exception {
        if(index == size) {
            index = 0;
        }
        EventLoop eventLoop = eventLoops[index++];
        eventLoop.register(socketChannel);
    }

}
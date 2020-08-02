package nio.reactor;

import java.net.InetSocketAddress;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;

public class Reactor implements Runnable{

    private int port;
    private EventLoopGroup eventLoopGroup;

    public Reactor(int port, EventLoopGroup eventLoopGroup) {
        this.port = port;
        this.eventLoopGroup = eventLoopGroup;
        new Thread(this).start();
    }

    @Override
    public void run() {
        try {
            ServerSocketChannel serverSocketChannel = ServerSocketChannel.open();
            serverSocketChannel.bind(new InetSocketAddress(port));
            while(!Thread.interrupted()) {
                SocketChannel socketChannel = serverSocketChannel.accept();
                eventLoopGroup.dispatch(socketChannel);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
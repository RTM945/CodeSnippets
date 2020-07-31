package nio.reactor;

import java.net.InetSocketAddress;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;

public class Reactor {

    private int port;
    private EventLoopGroup eventLoopgGroup;

    public Reactor(int port, EventLoopGroup eventLoopgGroup) {
        this.port = port;
        this.eventLoopgGroup = eventLoopgGroup;
    }

    public void listen() throws Exception {
        ServerSocketChannel serverSocketChannel = ServerSocketChannel.open();
        serverSocketChannel.bind(new InetSocketAddress(port));
        while(!Thread.interrupted()) {
            SocketChannel socketChannel = serverSocketChannel.accept();
            eventLoopgGroup.dispatch(socketChannel);
        }
    }
}
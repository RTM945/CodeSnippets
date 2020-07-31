package nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;

public class NioServer implements Runnable {

    private int port;
    private ServerSocketChannel serverSocketChannel;
    private volatile boolean stop;

    public NioServer(int port) throws IOException {
        this.port = port;
        serverSocketChannel = ServerSocketChannel.open();
        serverSocketChannel.socket().bind(new InetSocketAddress(this.port));
    }

    public void start() {
        new Thread(this).start();
    }

    @Override
    public void run() {
        try {
            while (!stop) {
                SocketChannel socketChannel = serverSocketChannel.accept();
                // new thread do something with channel
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            try {
                serverSocketChannel.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

    }

    public void close() {
        stop = true;
    }
}
package nio;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.ServerSocketChannel;
import java.util.Iterator;
import java.util.Set;

public class NonBlockingServer implements Runnable {
    private int port;
    private ServerSocketChannel serverSocketChannel;
    private Selector selector;
    private volatile boolean stop;

    public NonBlockingServer(int port) throws IOException {
        this.port = port;
        serverSocketChannel = ServerSocketChannel.open();
        serverSocketChannel.configureBlocking(false);
        serverSocketChannel.socket().bind(new InetSocketAddress(this.port));
        selector = Selector.open();
        serverSocketChannel.register(selector, SelectionKey.OP_ACCEPT);
    }

    public void start() {
        new Thread(this).start();
    }

    @Override
    public void run() {
        try {
            while(!stop) {
                selector.select();
                Set<SelectionKey> selectionKeys = selector.selectedKeys();
                Iterator<SelectionKey> iterator = selectionKeys.iterator();
                while (iterator.hasNext()) {
                    SelectionKey key = iterator.next();
                    if(key.isAcceptable()) {
                        //accept
                    }
                    if (key.isReadable()) {
                        //read
                    }
                    if (key.isWritable()) {
                        //write
                    }
                    iterator.remove();
                }
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
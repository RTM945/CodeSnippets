package javaassist;

public class CTest extends Msg{

    private int type = 1;

    @Override
    public void process() throws Exception {
        throw new UnsupportedOperationException();
    }

    public int getType() {
        return type;
    }
}

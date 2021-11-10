package script;

public class Player {
    
    private int x;

    public void move(int dx) {
        x += dx;
    }

    public int getX() {
        return x;
    }

}

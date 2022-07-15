package script;

public class Player {
    
    private int x;

    public void move(int dx) {
        x += dx;
    }

    public int getX() {
        return x;
    }

    public int id;
    public String name;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

}

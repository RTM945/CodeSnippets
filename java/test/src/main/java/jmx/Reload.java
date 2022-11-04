package jmx;

public class Reload implements ReloadMBean{
    @Override
    public void reload() {
        System.out.println("reload");
    }
}

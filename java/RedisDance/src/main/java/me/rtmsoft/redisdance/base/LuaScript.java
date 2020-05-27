package me.rtmsoft.redisdance.base;

public class LuaScript {

    private final String name;

    private final String script;

    private String sha1;

    public LuaScript(String name, String script) {
        this.name = name;
        this.script = script;
    }

    public String getName() {
        return name;
    }

    public String getScript() {
        return script;
    }

    public String getSha1() {
        return sha1;
    }

    public void setSha1(String sha1) {
        this.sha1 = sha1;
    }
}

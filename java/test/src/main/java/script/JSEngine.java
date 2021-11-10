package script;

import java.util.concurrent.ConcurrentMap;

import com.google.common.base.Objects;
import com.googlecode.concurrentlinkedhashmap.ConcurrentLinkedHashMap;

import org.openjdk.nashorn.api.scripting.ScriptObjectMirror;
import javax.script.ScriptEngine;
import javax.script.ScriptEngineManager;

// before jdk 15
// import jdk.nashorn.api.scripting.ScriptObjectMirror;

public class JSEngine {

    private static final ConcurrentMap<JsCacheKey, ScriptObjectMirror> cache = new ConcurrentLinkedHashMap.Builder<JsCacheKey, ScriptObjectMirror>().maximumWeightedCapacity(20480).build();
    private static final ScriptEngine engine = new ScriptEngineManager().getEngineByName("JavaScript");

    public static void main(String[] args) {
        // 不能使用内部类
        Player player = new Player();
        eval("player.move(dx)", "player, dx", player, 5);
        // should be 5
        System.out.println(player.getX());
    }

    public static Object eval(String js, String paramNames, Object... params) {
        try {
            JsCacheKey key = new JsCacheKey(js, paramNames);
            ScriptObjectMirror mirror = cache.computeIfAbsent(key, v -> {
                try {
                    String func = "function(" + paramNames + "){return " + js + "}";
                    return (ScriptObjectMirror) engine.eval(func);
                } catch (Exception e) {
                    e.printStackTrace();
                    return null;
                }
            });
            if (mirror == null) {
                return null;
            }
            return mirror.call(null, params);
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }


    static class JsCacheKey {
        String js;
        String params;

        public JsCacheKey(String js, String params) {
            this.js = js;
            this.params = params;
        }

        @Override
        public boolean equals(Object obj) {
            return Objects.equal(js, params);
        }

        @Override
        public int hashCode() {
            return Objects.hashCode(js, params);
        }
    }

    

}

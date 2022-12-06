package script;

import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.ConcurrentMap;

import com.google.common.base.Objects;
import com.googlecode.concurrentlinkedhashmap.ConcurrentLinkedHashMap;
import org.openjdk.nashorn.api.scripting.ScriptObjectMirror;

//import org.openjdk.nashorn.api.scripting.ScriptObjectMirror;
//import org.openjdk.nashorn.api.scripting.ScriptUtils;

import javax.script.*;

// before jdk 15
//import jdk.nashorn.api.scripting.ScriptObjectMirror;

public class JSEngine {

    private static final ConcurrentMap<JsCacheKey, ScriptObjectMirror> cache = new ConcurrentLinkedHashMap.Builder<JsCacheKey, ScriptObjectMirror>().maximumWeightedCapacity(20480).build();
    private static final ScriptEngine engine = new ScriptEngineManager().getEngineByName("JavaScript");

    public static void main(String[] args) throws Exception {
        // 不能使用内部类
        Player player = new Player();
        Path scriptPath = Paths.get("./test.js");
        byte[] scriptBytes = Files.readAllBytes(scriptPath);
        Bindings bindings = engine.createBindings();
        bindings.put("player", player);
        engine.setBindings(bindings, ScriptContext.ENGINE_SCOPE);
        engine.eval(new String(scriptBytes, StandardCharsets.UTF_8));
        System.out.println(player.getX());
        Invocable invocable = (Invocable) engine;
        invocable.invokeFunction("f1", player, "xixi");
        System.out.println(player.getName());

        // eval("player.move(dx)", "player, dx", player, 5);
        // should be 5
        // System.out.println(player.getX());
        // ScriptObjectMirror mirror = (ScriptObjectMirror) eval("[1,2,3,4]", "player, dx", null, null);
        // int[] res = mirror.to(int[].class);
        // int[] res = (int[]) ScriptUtils.convert(eval("[1,2,3,4]", "player, dx", null, null), int[].class);
        // System.out.println(Arrays.toString(res));
        //ScriptObjectMirror mirror = (ScriptObjectMirror) eval("[1,2,3].filter(it => ![].include(it))", "player, dx", null, null);
        //int[] res = mirror.to(int[].class);
        //System.out.println(Arrays.toString(res));
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

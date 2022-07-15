package script;

import org.luaj.vm2.LuaFunction;
import org.luaj.vm2.LuaValue;
import org.luaj.vm2.lib.jse.CoerceJavaToLua;

import javax.script.*;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths; 

public class LuaTest {
    

    public static void main(String[] args) throws Exception {
        Player player = new Player();
        ScriptEngineManager sem = new ScriptEngineManager();
        ScriptEngine e = sem.getEngineByName("luaj");
        Path scriptPath = Paths.get("./test1.lua");
        byte[] scriptBytes = Files.readAllBytes(scriptPath);
        String lua = new String(scriptBytes, StandardCharsets.UTF_8);
        CompiledScript compile = ((Compilable) e).compile(lua);
        Bindings sb = new SimpleBindings();
        sb.put("player", player);
        compile.eval(sb);
        LuaFunction f1 = (LuaFunction) sb.get("f1");
        LuaValue coerce = CoerceJavaToLua.coerce(player);
        LuaValue[] params = {coerce};
        f1.invoke(LuaValue.varargsOf(params));
        System.out.println(player.getId() + "_" + player.getName());
    }


}

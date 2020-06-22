package proxy.jdk;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;

public abstract class MyAbstractInvocationHandler implements InvocationHandler {

    private Object target;

    public MyAbstractInvocationHandler(Class<?> clazz) throws Exception {
        this.target = clazz.newInstance();
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        before(target, args);
        Object result = method.invoke(target, args);
        after(target, args, result);
        return result;
    }

    public abstract void before(Object target, Object[] args);

    public abstract void after(Object target, Object[] args, Object result);
}

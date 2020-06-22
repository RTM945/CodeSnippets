package proxy.cglib;

import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.MethodProxy;

import java.lang.reflect.Method;

public abstract class MyAbstractMethodInterceptor implements MethodInterceptor {

    @Override
    public Object intercept(Object target, Method method, Object[] args, MethodProxy proxy) throws Throwable {
        before(target, args);
        Object result = proxy.invokeSuper(target, args);
        after(target, args, result);
        return result;
    }

    public abstract void before(Object target, Object[] args);

    public abstract void after(Object target, Object[] args, Object result);
}

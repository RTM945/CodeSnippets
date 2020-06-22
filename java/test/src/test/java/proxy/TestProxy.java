package proxy;

import net.sf.cglib.proxy.Enhancer;
import org.junit.Test;
import proxy.cglib.MyAbstractMethodInterceptor;
import proxy.jdk.MyAbstractInvocationHandler;

import java.lang.reflect.Proxy;
import java.util.Arrays;
import java.util.List;

public class TestProxy {

    @Test
    public void testStaticProxy() {
        ATestInterface test = new AStaticProxyTestImpl();
        List<Integer> result = test.aTestMethod(Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10));
        System.out.println(result);
    }

    @Test
    public void testJDKProxy() {
        List<Integer> list = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        try {
            MyAbstractInvocationHandler handler = new MyAbstractInvocationHandler(ARealTestImpl.class) {
                @Override
                public void before(Object target, Object[] args) {
                    System.out.println("before jdk proxy");
                }

                @Override
                public void after(Object target, Object[] args, Object result) {
                    List<Integer> list = (List<Integer>) result;
                    for (int i = 0; i < list.size(); i++) {
                        int n = list.get(i) - 1;
                        list.set(i, n);
                    }
                }
            };
            ClassLoader cl = this.getClass().getClassLoader();
            ATestInterface ti = (ATestInterface) Proxy.newProxyInstance(cl, new Class[]{ATestInterface.class}, handler);
            List<Integer> result = ti.aTestMethod(list);
            System.out.println(result);
        } catch (Exception e) {
            e.printStackTrace();
        }

    }

    @Test
    public void testCglib() {
        List<Integer> list = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        try{
            Enhancer enhancer = new Enhancer();
            enhancer.setSuperclass(ARealTestImpl.class);
            enhancer.setCallback(new MyAbstractMethodInterceptor() {
                @Override
                public void before(Object target, Object[] args) {
                    System.out.println("before cglib proxy");
                }

                @Override
                public void after(Object target, Object[] args, Object result) {
                    List<Integer> list = (List<Integer>) result;
                    for (int i = 0; i < list.size(); i++) {
                        int n = list.get(i) - 1;
                        list.set(i, n);
                    }
                }
            });
            ATestInterface ti = (ATestInterface) enhancer.create();
            List<Integer> result = ti.aTestMethod(list);
            System.out.println(result);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}

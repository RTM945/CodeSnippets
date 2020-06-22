package proxy;

import java.util.List;
import java.util.stream.Collectors;

//测试用接口的静态代理
public class AStaticProxyTestImpl implements ATestInterface {

    private ATestInterface real;

    public AStaticProxyTestImpl() {
        real = new ARealTestImpl();
    }

    @Override
    public List<Integer> aTestMethod(List<Integer> list) {
        System.out.println("before static proxy");
        List<Integer> result = real.aTestMethod(list);
        result = result.stream().map(i -> i = i - 1).collect(Collectors.toList());
        return result;
    }
}

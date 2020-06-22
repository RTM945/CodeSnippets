package proxy;

import java.util.List;
import java.util.stream.Collectors;

//测试用接口的实际实现
public class ARealTestImpl implements ATestInterface {
    @Override
    public List<Integer> aTestMethod(List<Integer> list) {
        return list.stream()
                .filter(i -> i % 2 == 0)
                .collect(Collectors.toList());
    }
}

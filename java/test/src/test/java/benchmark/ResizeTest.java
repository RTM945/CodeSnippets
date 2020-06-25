package benchmark;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.infra.Blackhole;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;

/**
 * https://mp.weixin.qq.com/s/1C10VD5Nl-oo8rvyE6KAOw
 */
@BenchmarkMode(Mode.AverageTime) // 测试完成时间
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@Warmup(iterations = 2, time = 1, timeUnit = TimeUnit.SECONDS) // 预热 2 轮，每次 1s
@Measurement(iterations = 5, time = 3, timeUnit = TimeUnit.SECONDS) // 测试 5 轮，每次 3s
@Fork(1) // fork 1 个线程
@State(Scope.Thread) // 每个测试线程一个实例
public class ResizeTest {

    @Benchmark
    public void noSizeMapTest(Blackhole blackhole) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < 1024; i++) {
            map.put(i, i);
        }
        blackhole.consume(map);
    }

    @Benchmark
    public void setSizeMapTest(Blackhole blackhole) {
        Map<Integer, Integer> map = new HashMap<>(2048);
        for (int i = 0; i < 1024; i++) {
            map.put(i, i);
        }
        blackhole.consume(map);
    }

    @Benchmark
    public void noSizeListTest(Blackhole blackhole) {
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < 1024; i++) {
            list.add(i);
        }
        blackhole.consume(list);
    }

    @Benchmark
    public void setSizeListTest(Blackhole blackhole) {
        List<Integer> list = new ArrayList<>(2048);
        for (int i = 0; i < 1024; i++) {
            list.add(i);
        }
        blackhole.consume(list);
    }

    public static void main(String[] args) throws RunnerException {
        Options opts = new OptionsBuilder()//
                .include(ResizeTest.class.getSimpleName())//
                .build();
        new Runner(opts).run();
    }
}

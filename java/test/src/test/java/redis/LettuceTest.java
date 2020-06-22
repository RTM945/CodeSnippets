package redis;

import org.junit.Assert;
import org.junit.Test;
import static org.hamcrest.CoreMatchers.is;

public class LettuceTest {

    @Test
    public void test() {
        RedisOps.set("a", "b");
        Assert.assertThat(RedisOps.get("a"), is("b"));
    }
}

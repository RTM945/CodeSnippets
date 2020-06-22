package redis;

import org.junit.Test;

public class VerificationCodeTest {

    @Test
    public void test() {
        String code = VerificationCode.gen("test");
        System.out.println(code);
        System.out.println(VerificationCode.verify("test", "AABBCC"));
        System.out.println(VerificationCode.verify("test", code));
    }

}

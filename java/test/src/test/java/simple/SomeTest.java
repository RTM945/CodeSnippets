package simple;

public class SomeTest {
    
    public static void main(String[] args) {
        String s1 = "a";
        String s2 = new String("a");
        String s3 = "a";
        System.out.println(s1.equals(s2));
        System.out.println(s1 == s2);
        System.out.println(s1 == s3);
    }

}
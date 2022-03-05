package algorithms.leetcode._537_complex_number_multiplication;

public class _537 {
    // 复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：

    // 实部 是一个整数，取值范围是 [-100, 100]
    // 虚部 也是一个整数，取值范围是 [-100, 100]
    // i2 == -1
    // 给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。
    // (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i 
    
    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/complex-number-multiplication
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public String complexNumberMultiply(String num1, String num2) {
            String[] strs = num1.split("\\+");
            // 实部
            int a = Integer.parseInt(strs[0]);
            // 虚部
            int b = Integer.parseInt(strs[1].substring(0, strs[1].length() - 1));

            strs = num2.split("\\+");
            // 实部
            int c = Integer.parseInt(strs[0]);
            // 虚部
            int d = Integer.parseInt(strs[1].substring(0, strs[1].length() - 1));
            // (a + b) * (c + d) = ac + bc + ad + bd
            int s = a * c - b * d;
            int i = b * c + a * d;
            return s + "+" + i + "i";
        }
    }

    class Solution1 {
        public String complexNumberMultiply(String num1, String num2) {
            String[] complex1 = num1.split("\\+|i");
            String[] complex2 = num2.split("\\+|i");
            int real1 = Integer.parseInt(complex1[0]);
            int imag1 = Integer.parseInt(complex1[1]);
            int real2 = Integer.parseInt(complex2[0]);
            int imag2 = Integer.parseInt(complex2[1]);
            return String.format("%d+%di", real1 * real2 - imag1 * imag2, real1 * imag2 + imag1 * real2);
        }
    }

    public static void main(String[] args) {
        _537 _537 = new _537();
        Solution solution = _537.new Solution();
        System.out.println(solution.complexNumberMultiply("1+1i", "1+1i"));
        System.out.println(solution.complexNumberMultiply("1+-1i", "1+-1i"));
    }
}

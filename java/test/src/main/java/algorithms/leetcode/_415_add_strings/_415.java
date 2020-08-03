package algorithms.leetcode._415_add_strings;

//https://leetcode-cn.com/problems/add-strings
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
// 提示：
// num1 和num2 的长度都小于 5100
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
public class _415 {
    class Solution {
        //将字符位数对其，然后判断每一位的值得出结果，再维护一个进位？
        public String addStrings(String num1, String num2) {
            int i = num1.length() - 1;
            int j = num2.length() - 1;
            int f = 0; //进位
            StringBuilder sb = new StringBuilder();
            while(i >= 0 || j >= 0 || f != 0) {
                int a = i >= 0 ? num1.charAt(i) - '0' : 0;
                int b = j >= 0 ? num2.charAt(j) - '0' : 0;
                int c = a + b + f;
                if (c > 9) {
                    c = c % 10;
                    f = 1;
                }else{
                    f = 0;
                }
                sb.append(c);
                i--;
                j--;
            }
            sb.reverse();
            return sb.toString();
        }
    }

    public static void main(String[] args) {
        _415 q = new _415();
        System.out.println(q.new Solution().addStrings("123", "457"));
    }
}
package algorithms.leetcode._400_nth_digit;

public class _400 {
    // 给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...]
    // 中找出并返回第 n 位数字。
    // 输入：n = 11
    // 输出：0
    // 解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，
    // 它是 10 的一部分。
    // https://leetcode-cn.com/problems/nth-digit/
    class Solution {
        // 其实是1234567891011..这个字符串的第十位，但这样写会超时
        public int findNthDigit(int n) {
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < n; i++) {
                sb.append(i + 1);
            }
            return Integer.parseInt(new String(Character.toString(sb.charAt(n - 1))));
        }
    }

    class Solution1 {
        // 找规律
        // 一位数 1 个 digit， 有 9 个（1~9），个位数共占 9 个 digit；
        // 二位数 2 个 digit， 有 90 个（10~99），十位数共占 180 个 digit；
        // 三位数 3 个 digit， 有 900 个（100~999），百位数共占 2700 个 digit；
        // i 位数 i 个 digit， 有 9*10^(i-1) 个（设此数字为k），共占 i * k 个 digit。
        public int findNthDigit(int n) {
            long k = 9, i = 1;
            while (n - k * i > 0) { 
                n -= k * i; 
                k *= 10;
                i += 1;
            }
            // 得到n为i位数，
            long num = k / 9 + (n - 1) / i; // 确定具体数字
            n -= i * (num - k / 9); // 确定最后还剩几个 digit
            return String.valueOf(num).toCharArray()[n - 1] - '0';
        }
    }

    public static void main(String[] args) {
        _400 _400 = new _400();
        System.out.println(_400.new Solution().findNthDigit(11));
        System.out.println(_400.new Solution().findNthDigit(3));
    }
}

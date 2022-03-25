package algorithms.leetcode._172_factorial_trailing_zeroes;

public class _172 {
    // 给定一个整数 n ，返回 n! 结果中尾随零的数量。
    // 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
    class Solution {
        public int trailingZeroes(int n) {
            // 有多少个 10 就有多少个 0
            // 10 又可以分解成 2 * 5
            // 等于 n! 分解出多少个 2 或 多少个 5，取最小
            // 必然是 5
            int res = 0;
            while (n >= 5) {
                res += n /= 5;
            }
            return res;
        }
    }
}

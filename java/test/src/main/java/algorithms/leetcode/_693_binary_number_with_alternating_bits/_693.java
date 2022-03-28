package algorithms.leetcode._693_binary_number_with_alternating_bits;

public class _693 {
    //给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：
    // 换句话说，就是二进制表示中相邻两位的数字永不相同。
    class Solution {
        public boolean hasAlternatingBits(int n) {
            // 如果 n 符合条件
            // n >> 1 为 001010101...
            // n ^ (n >> 1) 为 0111111...
            // 011111... + 1 = 1000000...
            // 011111... & 1000000... = 0
            int m = n ^ (n >> 1);
            return ((m + 1) & m) == 0;
        }
    }

    class Solution1 {
        public boolean hasAlternatingBits(int n) {
            int prev = 2;
            while (n != 0) {
                // cur 要么是 0 要么是 1
                int cur = n % 2;
                if (cur == prev) {
                    return false;
                }
                prev = cur;
                n /= 2;
            }
            return true;
        }
    }
}

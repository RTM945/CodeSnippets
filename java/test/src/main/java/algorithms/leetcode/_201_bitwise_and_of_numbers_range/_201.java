package algorithms.leetcode._201_bitwise_and_of_numbers_range;

/* 
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
示例 1: 
输入: [5,7]
输出: 4
示例 2:
输入: [0,1]
输出: 0
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bitwise-and-of-numbers-range
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _201 {
    class Solution {
        // 第一次做这种题目
        // 先直接与试试看
        // [0, 2147483647]超时了
        public int rangeBitwiseAnd(int m, int n) {
            int res = m;
            for (int i = m + 1; i <= n; i++) {
                res = res & i;
            }
            return res;
        }
    }

    class Solution1 {
        // 想到之前有个面试题
        // 如何判断一个数是2的幂
        // n & (n - 1) == 0
        // 未被抵消的和0相与，得到的也是0
        // 只要[m, n]包含2的幂，结果一定是0
        // 书面的说法，对所有数字执行按位与运算的结果
        // 是所有对应二进制字符串的公共前缀再用零补上后面的剩余位。
        // 一般做法是两个数字都右移到两个数字相等，并记录移了多少位shift
        // 此时他们的公共前缀就得到了，再左移shift位补上0，就能得出结果
        // 另一个方法是Brian Kernighan 算法
        // n & (n - 1)可以去掉n最右边的1
        // 判断2的幂是这个算法的一个特殊应用，因为2的幂转为二进制只有一个1
        // 一直进行n = n & (n - 1)直到n <= m
        // 这样n和m的公共前缀得以保留，并且公共前缀右边的1全部变成了0
        public int rangeBitwiseAnd(int m, int n) {
            while (n > m) {
                n = n & (n - 1);
            }
            return n;
        }
    }

    public static void main(String[] args) {
        _201 q = new _201();
        System.out.println(q.new Solution().rangeBitwiseAnd(0, 1));
        System.out.println(q.new Solution().rangeBitwiseAnd(5, 7));
    }
}
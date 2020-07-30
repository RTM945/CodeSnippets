package algorithms.leetcode._343_integer_break;

// https://leetcode-cn.com/problems/integer-break
// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。
// 返回你可以获得的最大乘积。
// 示例 1:
// 输入: 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1。
// 示例 2:
// 输入: 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
// 说明: 你可以假设 n 不小于 2 且不大于 58。
public class _343 {
    class Solution {
        //应该是dp?
        //看答案的，自己想不出来，很难受
        //n可以拆分成i和n - i，i范围是(0, n)
        //将n - i看作n，又可以拆分成i和n - i
        //递归
        public int integerBreak(int n) {
            int res = 0;
            for (int i = 1; i < n; i++) {
                res = Math.max(res, i * (n - i)); //分成两个的情况
                res = Math.max(res, i * integerBreak(n - i)); //分多个的情况
            }
            return res;
        }
    }

    class Solution1 {
        //递归好理解，但超时
        //需要记忆
        public int integerBreak(int n) {
            int[] memo = new int[n + 1];
            return dfs(n, memo);
        }

        int dfs(int n, int[] memo) {
            if(memo[n] != 0) {
                return memo[n];
            }
            int res = 0;
            for (int i = 1; i < n; i++) {
                res = Math.max(res, i * (n - i)); //分成两个的情况
                res = Math.max(res, i * dfs(n - i, memo)); //分多个的情况
            }
            memo[n] = res;
            return res;
        }
    }

    class Solution2{
        //dp来了
        //设i的拆分乘积最大值为dp[i]
        //dp[i] = max(1 * dp[i - 1], 2 * dp[i - 2], ..., j * dp[i - j])
        //dp[i] = max(dp[i], i * (i - j))
        public int integerBreak(int n) {
            int[] dp = new int[n + 1];
            for (int i = 2; i <= n; i++) {
                for (int j = 1; j < i; j++) {
                    dp[i] = Math.max(dp[i], j * (i - j)); 
                    dp[i] = Math.max(dp[i], j * dp[i - j]); 
                }
            }
            return dp[n];
        }
    }

    class Solution3{
        //数学方法
        //https://leetcode-cn.com/problems/integer-break/solution/343-zheng-shu-chai-fen-tan-xin-by-jyd/
        //n = n1 + n2 + n3 + ... + nn
        //求max(n1 * n2 * n3 * ... * nn)
        //此处有均值不等式
        //(n1 + n2 + n3 + ... + nn) / a >= (n1 * n2 * n3 * ... * nn)^(1/a)
        //当n1 = n2 = n3 = ... = nn时，等号成立，也就是期望的最大值
        //将n分成a份相同的值，n=ax，ax / a >= (n1 * n2 * n3 * ... * nn)^(1/a)
        //最大值为x^a = x^(n/x) = (x^(1/x))^n
        //n确定的，就是求y = x^(1/x)的最大值了
        //lny = (1/x)lnx
        //接下来求导 得出x应该为自然对数e
        //题目要求拆分成正整数，找最接近e的正整数2，3
        //即尽可能的把n分成多个3，余下是2就用2
        public int integerBreak(int n) {
            if(n <= 3) {
                return n - 1;
            }
            int a = n / 3;
            int b = n % 3;
            if(b == 0) {
                return (int)Math.pow(3, a);
            }
            if(b == 1) {
                return (int)Math.pow(3, a - 1) * 4;
            }
            return (int)Math.pow(3, a) * 2;
        }
    }

    public static void main(String[] args) {
        _343 q = new _343();
        System.out.println(q.new Solution().integerBreak(10));
        System.out.println(q.new Solution1().integerBreak(10));
        System.out.println(q.new Solution2().integerBreak(10));
    }
}
package algorithms.leetcode._322_coin_change;

import java.util.Arrays;

// https://leetcode-cn.com/problems/coin-change
// 给定不同面额的硬币 coins 和一个总金额 amount。
// 编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
// 示例 1:
// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3 
// 解释: 11 = 5 + 5 + 1
// 示例 2:
// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:
// 你可以认为每种硬币的数量是无限的。
public class _322 {
    class Solution {
        //中等
        //拒说是非常经典的dp题目
        //想想递归关系
        //假设硬币面值为v1,v2,v3...vn
        //个数分别为c1,c2,c3...cn
        //c1v1 + c2v2 + c3v3 + ... + cnvn = amount
        //且c1 + c2 + c3 + ... + cn是最优解
        //那么c1 + c2 + c3 + ... + c(n-1) 是 amount - cnvn的最优解
        //好像是废话
        //设dp[i]为amount = i时的最优解
        //那么dp[i - 1] = ? 思考不下去
        //答案用的方法比较巧妙..
        //虽然dp[i - 1]不知道
        //但dp[i - coins[k]] = dp[i] - 1 
        //现在只需要在k个选择中找出最小值
        //即dp[i] = min(dp[i - coins[0]], dp[i - coins[1]], ..., dp[i - coins[n]]) + 1
        public int coinChange(int[] coins, int amount) {
            int[] dp = new int[amount + 1];
            Arrays.fill(dp, amount + 1); //硬币个数不可能比amount多
            dp[0] = 0;
            for (int i = 1; i <= amount; i++) {
                for (int j = 0; j < coins.length; j++) {
                    if(i >= coins[j]) {
                        dp[i] = Math.min(dp[i], dp[i - coins[j]] + 1);
                    }
                }
            }
            return dp[amount] > amount ? -1 : dp[amount];
        }
        //对于测试用例[1, 2, 5]来说
        //正好dp[1] = dp[0] + 1 = 1
        //但可能coins[0]不为1，此时dp[i]的值会大于i
        //也就是不存在解，返回-1
    }

    public static void main(String[] args) {
        _322 q = new _322();
        System.out.println(q.new Solution().coinChange(new int[]{1, 2, 5}, 11));
        System.out.println(q.new Solution().coinChange(new int[]{3}, 2));
    }
}
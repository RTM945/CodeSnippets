package algorithms.leetcode._309_best_time_to_buy_and_sell_stock_with_cooldown;

import java.util.Map;

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//示例:
//输入: [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
public class _309 {

    class Solution {
        //如果用暴力解法，需要列举出所有操作的排列组合
        //且第一天肯定是买入，最后一天肯定是卖出，还要遵循冷冻期条件
        //又必须在再次购买前出售掉之前的股票
        //那么排列组合就很固定了
        //如果总共5天，交易状态只有[买入, 卖出, 冷冻期, 买入, 卖出]这一种
        //这样就很让人迷惑..
        //只有[买入, 卖出, 冷冻期]...的循环？从何谈起最优解
        public int maxProfit(int[] prices) {
            int sum = 0;
            for (int i = 0; i < prices.length; i++) {
                if(i % 3 == 0) {
                    sum -= prices[i];
                }else if(i % 3 == 1) {
                    sum += prices[i];
                }
                //冷冻期不需要计算
            }
            if(sum < 0) {
                return 0;
            }
            return sum;
        }
    }

    class Solution1 {
        //[1,2,4]预期结果是3，上面的方法是1
        //1是因为严格按照买入、卖出、冷冻的方法来的
        //结果3只有一种可能，第一天买入1，第二天没操作，第三天卖了
        //也就是说还有不操作的选项
        //规则变成了买入后可以不操作和卖出，卖出后被冷冻
        //可以用树来思考
        //但过于繁琐，怎样才能转换成dp/可递归问题/最优子结构
        //看答案吧
        //带状态的动态规划，即二维数组dp[i][0]表示第i天未持有该股票，dp[i][1]表示持有
        //dp[i][j]为最大收益
        //dp[i][0]可以是第i-1天持有股票，第i天卖出
        //dp[i][0] = dp[i - 1][1] + prices[i]
        //也可以是卖出后冷冻/第一天没买
        //dp[i][0] = dp[i - 1][0]
        //需要取最大值
        //dp[i][0] = max(dp[i - 1][1] + prices[i], dp[i - 1][0])
        //dp[i][1]可以是第i-1天买入股票，第i天没操作
        //dp[i][1] = dp[i - 1][1]
        //也可以是第i-2天未持有，隔了一天冷冻期，第i天买入
        //dp[i][1] = dp[i - 2][0] - prices[i]
        //仍然取最大值
        //dp[i][1] = max(dp[i - 2][0] - prices[i], dp[i - 1][0] - prices[i])
        public int maxProfit(int[] prices) {
            if(prices.length <= 1) {
                return 0;
            }
            int[][] dp = new int[prices.length][2];
            dp[0][0] = 0; //第一天未持有，总资产0
            dp[0][1] = -prices[0]; //第一天持有，总资产负
            for (int i = 1; i < prices.length; i++) {
                dp[i][0] = Math.max(dp[i - 1][1] + prices[i], dp[i - 1][0]);
                //上一个未持有股票时的资产
                int temp = 0;
                if(i > 1) {
                    temp = dp[i - 2][0]; //前天卖了，昨天冷冻了
                }
                dp[i][1] = Math.max(temp - prices[i], dp[i - 1][1]);
            }
            return Math.max(dp[prices.length - 1][0], dp[prices.length - 1][1]);
        }
    }

    class Solution2 {
        //一个优化dp空间的做法
        //第i天的总资产和前两天资产与操作有关
        //dp[i − 2][0] dp[i - 1][0] dp[i - 1][1]
        public int maxProfit(int[] prices) {
            if(prices.length <= 1) {
                return 0;
            }
            int[] dp = new int[3];
            dp[1] = -prices[0];
            for (int i = 1; i < prices.length; i++) {
                dp[1] = Math.max(dp[1], dp[2] - prices[i]);
                dp[2] = dp[0];
                dp[0] = Math.max(dp[0], dp[1] + prices[i]);
            }
            return Math.max(dp[0], dp[1]);
        }
    }

    public static void main(String[] args) {
        _309 q = new _309();
        System.out.println(q.new Solution().maxProfit(new int[]{1,2,3,0,2}));
        System.out.println(q.new Solution1().maxProfit(new int[]{1,2,3,0,2}));
    }
}

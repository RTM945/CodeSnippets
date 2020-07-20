package algorithms.leetcode._312_burst_balloons;

import java.util.Arrays;

//https://leetcode-cn.com/problems/burst-balloons/
//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//现在要求你戳破所有的气球。
//如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 
//这里的 left 和 right 代表和 i 相邻的两个气球的序号。
//注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
//求所能获得硬币的最大数量。
//说明:
//你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
//0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
//示例:
//输入: [3,1,5,8]
//输出: 167
//解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
public class _312 {
    //hard
    //觉得应该用dp，但不知道怎么用啊
    //假如暴力，应该怎么做呢
    //使用过气球不能再使用，很明显局部最优并不是全局最优
    //int len = nums.length
    //一开始有len个选择
    //从第i个开始戳戳爆了剩下有len - 1种选择，
    //每种选择又有len - 2个选择
    //以此类推下去
    //找出最大值
    //在一个格结构中找出和最大的一条路径
    //还不知道怎么办
    //看答案吧
    //首先在数组前后添加1，方便边界处理
    //每次戳爆一个气球会导致两个气球从不相邻变成相邻，不好处理
    //所以反过来看，每次添加一个气球
    //即[1,1] --> [1,8,1] --> [1,3,8,1] --> [1,3,5,8,1] --> [1,3,1,5,8,1]
    //使用solve(i, j)表示[i, j]区间中填满硬币后所得硬币的最大值
    //两端的气球为i和j
    //当i >= j - 1时，不存在气球，solve(i, j) = 0
    //当i < j - 1时，使用i，j之间的任一气球作为mid，当作是第一个加入的气球
    //那么他的得分是nums[i] * nums[mid] * nums[j]
    //再分治，使用[i,mid]或[mid,j]区间，在他们之间加入气球，比较得出最优解
    //问题就转变成了求nums[i] * nums[mid] * nums[j] + solve(i, mid) + solve(mid, j)
    //防止重复计算，记录中间结果
    class Solution {
        public int[][] rec;
        public int[] val;
        public int maxCoins(int[] nums) {
            int n = nums.length;
            val = new int[n + 2];
            val[0] = 1;
            val[n + 1] = 1;
            for (int i = 1; i <= n; i++) {
                val[i] = nums[i - 1];
            }
            rec = new int[n + 2][n + 2];
            for (int i = 0; i <= n + 1; i++) {
                Arrays.fill(rec[i], -1); //填满-1，方便比较
            }
            return solve(0, n + 1);
        }

        public int solve(int left, int right) {
            if (left >= right - 1) {
                return 0;
            }
            if (rec[left][right] != -1) {
                return rec[left][right]; //rec是已经结算好的最优解
            }
            for (int i = left + 1; i < right; i++) {
                int sum = val[left] * val[i] * val[right];
                sum += solve(left, i) + solve(i, right);
                rec[left][right] = Math.max(rec[left][right], sum);
            }
            return rec[left][right];
        }
    }

    class Solution1{
        //dp解法
        //将自顶向下的递归变为自底向上的dp
        //dp[i][j]表示[i,j]区间能得到的最大硬币
        //当i >= j - 1时，dp[i][j] = 0
        //dp[i][j] = max(dp[i] * dp[mid] * dp[j] + dp[i][mid] + dp[mid][j])
        //mid = [i, j-1)
        public int maxCoins(int[] nums) {
            int n = nums.length;
            int[][] rec = new int[n + 2][n + 2];
            int[] val = new int[n + 2];
            val[0] = val[n + 1] = 1;
            for (int i = 1; i <= n; i++) {
                val[i] = nums[i - 1];
            }
            for (int i = n - 1; i >= 0; i--) {//从后往前
                for (int j = i + 2; j <= n + 1; j++) {//从后往前
                    for (int k = i + 1; k < j; k++) {//区间内的任意mid
                        int sum = val[i] * val[k] * val[j];
                        sum += rec[i][k] + rec[k][j];
                        rec[i][j] = Math.max(rec[i][j], sum);
                    }
                }
            }
            return rec[0][n + 1];
        }
    }
}

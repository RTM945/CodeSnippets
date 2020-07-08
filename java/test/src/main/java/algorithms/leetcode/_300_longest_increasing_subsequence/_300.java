package algorithms.leetcode._300_longest_increasing_subsequence;

//https://leetcode-cn.com/problems/longest-increasing-subsequence/
//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//示例:
//输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//说明:
//可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
//你算法的时间复杂度应该为 O(n2) 。
//进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
public class _300 {
    class Solution {
        //DP解法
        //设f(i)为以bi结束的最长不下降序列
        //比如f(12)为63结束的最长不下降序列
        //f(12) = max(f(1), f(2), f(3), f(4), f(5), f(6), f(7), f(8), f(9), f(10), f(11))
        //f(11) = max(f(1), f(2), f(3), f(4), f(5), f(6), f(7), f(8), f(9), f(10))
        //以此类推，f(i)有i-1个子问题
        //递归通式
        //f(i) = Max{f(j)+1 | j < i and bj <= bi}
        //边界条件
        //f(0) = 1
        public int lengthOfLIS(int[] nums) {
            if (nums.length == 0) {
                return 0;
            }
            int[] dp = new int[nums.length];
            dp[0] = 1;
            int result = 1;//每一位都是长度为1的不下降序列
            for (int i = 1; i < nums.length; i++) {
            /*if(nums[i] > nums[i - 1]) {
                //如果后一位大于前一位，说明可以构成不下降序列
                //其长度为子不下降序列长度+1
                //但这一步做法是错误的，没有考虑周全
                //忽略了类似3, 1 ,4, 2, 5, 7这样的情况
                //只能应用于连续递增，无跳跃的情况
                dp[i] = dp[i - 1] + 1;
            }*/
                //正确做法是nums[i]需要和前面的每一位进行判断
                //一次子循环可以得出在i范围内的最长不下降序列
                //也就是f(i)，需要将其存入dp
                int max = 0;
                for (int j = 0; j < i; j++) {
                    //当后一位大于前一位时，只能形成长度为1的不下降序列
                    //但如果后一位小于前一位满足条件
                    //则将其子问题的解+1作为本次f(i)的解
                    //子问题的解也就是f(0)...f(i-1)的最大值
                    //就是max
                    if(nums[i] > nums[j]) {
                        if(max < dp[j]) {
                            max = dp[j];
                        }
                    }
                    //这一步其实很迷惑，为什么要写在判断外？
                    //会不会出现一直累加的情况？
                    //如果循环内没有不下降序列
                    //那么max为0，dp[i]=1，对应不下降序列只有nums[i]一位的情况
                    //如果循环内有不下降序列
                    //如果nums[i] > nums[j]，dp[i] = max + 1没有任何问题
                    //如果其中有nums[i] < nums[j]的情况
                    //1, 3, 2希望的dp是[1, 2, 2] i = 2 dp[0] = 1
                    //第一遍循环j = 0，dp=[1, 2, 0] max = dp[0] = 1
                    //第二遍循环j = 1，dp=[1, 2, 2] max = dp[1] = 2
                    //但这里并没有将新的dp[i]赋值给max
                    //也就是说max始终是f(j)的最优解
                    //所以此处没有错误
                    dp[i] = max + 1;
                    //与全局最优比较
                    if(result < dp[i]) {
                        result = dp[i];
                    }
                }
            }
            return result;
            //这个解法的时间复杂度是O(n^2)
            //还有O(nlogn)的解法，但不属于DP
        }
    }
}

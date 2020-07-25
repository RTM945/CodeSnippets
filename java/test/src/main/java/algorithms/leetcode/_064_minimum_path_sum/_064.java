package algorithms.leetcode._064_minimum_path_sum;

//https://leetcode-cn.com/problems/minimum-path-sum
//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。
//示例:
//输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
public class _064 {

    class Solution {
        //久违的dp
        //dp[i][j]为左上角到达(i, j)的最短路径
        //dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
        //开头多加一行一列0
        public int minPathSum(int[][] grid) {
            int m = grid.length;
            int n = grid[0].length;
            int[][] dp = new int[m + 1][n + 1];
            for (int i = 2; i < dp.length; i++) {
                dp[i][0] = Integer.MAX_VALUE;
            }
            for (int i = 2; i < dp[0].length; i++) {
                dp[0][i] = Integer.MAX_VALUE;
            }
            dp[1][1] = grid[0][0];
            for (int i = 1; i < m + 1; i++) {
                for (int j = 1; j < n + 1; j++) {
                    dp[i][j] = Math.min(dp[i - 1][j], dp[i][j - 1]) + grid[i - 1][j - 1];
                }
            }
            return dp[m][n];
        }
    }

    public static void main(String[] args) {
        _064 q = new _064();
        int[][] grid = {
                {1, 3, 1},
                {1, 5, 1},
                {4, 2, 1},
        };
//        int[][] grid = {
//                {1},
//        };
        System.out.println(q.new Solution().minPathSum(grid));
    }
}

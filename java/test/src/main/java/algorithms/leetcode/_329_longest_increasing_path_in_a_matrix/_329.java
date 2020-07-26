package algorithms.leetcode._329_longest_increasing_path_in_a_matrix;

import java.util.Arrays;

//https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
// 给定一个整数矩阵，找出最长递增路径的长度。
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 
// 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
// 示例 1:
// 输入: nums = 
// [
//   [9,9,4],
//   [6,6,8],
//   [2,1,1]
// ] 
// 输出: 4 
// 解释: 最长递增路径为 [1, 2, 6, 9]。
// 示例 2:
// 输入: nums = 
// [
//   [3,4,5],
//   [3,2,6],
//   [2,2,1]
// ] 
// 输出: 4 
// 解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
public class _329{
    //hard
    //首先想可能需要找个最大值最小值，直觉来说肯定不是这么做的
    //假设dp[i][j]为matrix[i][j]的最长递增路径的长度
    //如果matrix[i][j]大于上下左右的数字，那么dp[i][j]的数字更新为上下左右dp值较大的+1
    class Solution {
        public int longestIncreasingPath(int[][] matrix) {
            int m = matrix.length;
            int n = matrix[0].length;
            int[][] dp = new int[m][n];
            int res = 0;
            for(int i = 0; i < m; i++) {
                for(int j = 0; j < n; j++) {
                    int[] round = {0, 0, 0, 0}; //上下左右
                    int v = matrix[i][j];
                    if(i - 1 >= 0 && v < matrix[i - 1][j]) {
                        round[0] = dp[i - 1][j];
                    }
                    if(i + 1 < m && v < matrix[i + 1][j]) {
                        round[1] = dp[i + 1][j];
                    }
                    if(j - 1 >= 0 && v < matrix[i][j - 1]) {
                        round[2] = dp[i][j - 1];
                    }
                    if(j + 1 < n && v < matrix[i][j + 1]) {
                        round[3] = dp[i][j + 1];
                    }
                    Arrays.sort(round);
                    dp[i][j] = round[3] + 1;
                    res = Math.max(res, dp[i][j]);
                }
            }
            return res + 1;
        }
    }

    //上面的方法错了
    //状态转移不对，在使用上下左右dp来计算中间状态的时候没考虑到上下左右应有的状态
    //如果要使用上面的方法，得从矩阵中较小的元素开始，显然是不现实的
    //看了下题解，真有人这么做，还把matrix[i][j],i,j封装成对象
    //matrix[i][j]用来排序，ij用来设置dp，也是厉害
    class Solution1 {
        //看到使用dfs+记忆化的做法
        //dp就是用来记忆的
        public int longestIncreasingPath(int[][] matrix) {
            int m = matrix.length;
            int n = matrix[0].length;
            int[][] dp = new int[m][n];
            int res = 0;
            for(int i = 0; i < m; i++) {
                for(int j = 0; j < n; j++) {
                    int longest = dp[i][j] == 0 ? dfs(matrix, dp, i, j) : dp[i][j];
                    res = Math.max(res, longest);
                }
            }
            return res;
        }

        int dfs(int[][] matrix, int[][] dp, int i, int j) {
            int m = matrix.length;
            int n = matrix[0].length;
            int longest = 1;
            int v = matrix[i][j];
            if(i > 0 && v > matrix[i - 1][j]) {
                if(dp[i - 1][j] > 0) { //已经遍历过了
                    longest = Math.max(longest, dp[i - 1][j] + 1);
                }else{
                    //递归找路
                    longest = Math.max(longest, dfs(matrix, dp, i - 1, j) + 1);//dp初始化是0，结果需要+1
                }
            }
            //下左右同样
            if(i < m - 1 && v > matrix[i + 1][j]) {
                if(dp[i + 1][j] > 0) { //已经遍历过了
                    longest = Math.max(longest, dp[i + 1][j] + 1);
                }else{
                    //递归找路
                    longest = Math.max(longest, dfs(matrix, dp, i + 1, j) + 1);//dp初始化是0，结果需要+1
                }
            }
            if(j > 0 && v > matrix[i][j - 1]) {
                if(dp[i][j - 1] > 0) { //已经遍历过了
                    longest = Math.max(longest, dp[i][j - 1] + 1);
                }else{
                    //递归找路
                    longest = Math.max(longest, dfs(matrix, dp, i, j - 1) + 1);//dp初始化是0，结果需要+1
                }
            }
            if(j < n - 1 && v > matrix[i][j + 1]) {
                if(dp[i][j + 1] > 0) { //已经遍历过了
                    longest = Math.max(longest, dp[i][j + 1] + 1);
                }else{
                    //递归找路
                    longest = Math.max(longest, dfs(matrix, dp, i, j + 1) + 1);//dp初始化是0，结果需要+1
                }
            }
            dp[i][j] = longest;
            return longest;
        }
    }

    public static void main(String[] args) {
        _329 q = new _329();
        int[][] matrix = {
            {9, 9, 4},
            {6, 6, 8},
            {2, 1, 1}
        };
        System.out.println(q.new Solution1().longestIncreasingPath(matrix));
    }
}
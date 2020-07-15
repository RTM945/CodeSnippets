package algorithms.leetcode._120_triangle;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

//https://leetcode-cn.com/problems/triangle/
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//例如，给定三角形：
//[
//   [2],
//  [3,4],
// [6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//说明：
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
public class _120 {
    class Solution {
        //一开始想到的是完全二叉树的性质
        //即左子结点index为父结点index * 2,右子节点index=父结点index * 2 + 1
        //三角形并非二叉树
        //但也许有其性质
        //但如果用dp，也许就不用关心性质，直接暴力解决
        //设dp[i][j]为到达i层的最小路径和
        //dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle.get(i).get(j)
        //dp[1][0] = dp[0][0] + triangle.get(1).get(0)
        //dp[1][1] = dp[0][0] + triangle.get(1).get(1)
        //dp[2][0] = dp[1][0] + triangle.get(2).get(0)
        //dp[2][2] = dp[1][1] + triangle.get(2).get(2)
        //dp[2][1] = min(dp[1][0], dp[1][1]) + triangle.get(2).get(1)
        //在最后一层中找最小的
        public int minimumTotal(List<List<Integer>> triangle) {
            int height = triangle.size();
            if(height == 1) {
                return triangle.get(0).get(0);
            }
            int[][] dp = new int[height][height];
            dp[0][0] = triangle.get(0).get(0);
            int res = Integer.MAX_VALUE;
            for (int i = 1; i < height; i++) {
                for (int j = 0; j <= i; j++) {
                    int up = 0;
                    if(j == 0) {
                        up = dp[i - 1][0];
                    }else if(j == i) {
                        up = dp[i - 1][j - 1];
                    } else {
                        up = Math.min(dp[i-1][j - 1], dp[i-1][j]);
                    }
                    dp[i][j] = up + triangle.get(i).get(j);
                    if(i == height - 1) {
                        res = Math.min(res, dp[i][j]);
                    }
                }
            }
            return res;
        }
    }

    class Solution1 {
        //写上面方法的时候就想到可以用一维数组代替二维数组了
        //dp[i]表示从根到第i层的最短路径和
        //同样的局部最优不一定是全局最优
        //所以从后往前dp
        //以上面的例子
        //到第4层最短路径和为第3层的所有路径和与第4层的每个结点的和的最小值
        //这里出现了问题，一维数组无法同时表示出每一层的所有路径和
        //又要转换思路
        //看看答案
        //回顾上面方法中的状态转移方程
        //dp[i][j] = dp[i - 1][0] + triangle[i][0]  j = 0
        //dp[i][j] = dp[i - 1][i - 1] + triangle[i][i]  j = i
        //dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1][j]) + triangle[i][j] otherwise
        //所以需要知道的只有dp[i - 1][...]的各个情况
        //使用两个一维数组来交替的更新dp[i]和dp[i - 1]
        //为了准确的把每一行映射到2个数组上，可以使用取模
        public int minimumTotal(List<List<Integer>> triangle) {
            int height = triangle.size();
            if(height == 1) {
                return triangle.get(0).get(0);
            }
            int[][] dp = new int[2][height];
            dp[0][0] = triangle.get(0).get(0);
            for (int i = 1; i < height; i++) {
                int mod = i % 2;
                int prev = 1 - mod;
                dp[mod][0] = dp[prev][0] + triangle.get(i).get(0);
                for (int j = 1; j <= i; j++) {
                    dp[mod][j] = Math.min(dp[prev][j], dp[prev][j - 1]) + triangle.get(i).get(j);
                }
                dp[mod][i] = dp[prev][i - 1] + triangle.get(i).get(i);
            }
            //使用对应最后一行的最小值
            int[] a = dp[1 - height % 2];
            int res = Integer.MAX_VALUE;
            for (int i = 0; i < a.length; i++) {
                if(res > a[i]) {
                    res = a[i];
                }
            }
            return res;
        }
    }

    class Solution2{
        //还能再优化
        //在使用二维数组时就能注意到
        //第i层只会用到前i个元素
        //而第i + 1层多出来的最后的元素
        //路径值必是上一层最后一个值+这一层最后一个值
        //如果该层从后往前遍历
        //就是dp[i - 1] + triangle[i][i]
        //倒数第二个元素为min(dp[i - 1], dp[i - 2]) + triangle[i][j]
        //反应到j上就是min(dp[j], dp[j - 1]) + triangle[i][j]
        //而第一个元素必是上一层第一个元素+这一层第一个值
        //也就是dp[0] += triangle[i][0]
        //从后往前遍历的情况下，就可以将前面的dp值当作上一层的dp值
        //真有你的！
        public int minimumTotal(List<List<Integer>> triangle) {
            int height = triangle.size();
            if(height == 1) {
                return triangle.get(0).get(0);
            }
            int[] dp = new int[height];
            dp[0] = triangle.get(0).get(0);
            for (int i = 1; i < height; i++) {
                dp[i] = dp[i - 1] + triangle.get(i).get(i);
                for (int j = i - 1; j > 0; j--) {
                    dp[j] = Math.min(dp[j - 1], dp[j]) + triangle.get(i).get(j);
                }
                dp[0] += triangle.get(i).get(0);
            }
            int res = Integer.MAX_VALUE;
            for (int i = 0; i < dp.length; i++) {
                if(res > dp[i]) {
                    res = dp[i];
                }
            }
            return res;
        }
    }

    public static void main(String[] args) {
        _120 q = new _120();
        List<List<Integer>> triangle = new ArrayList<>();
        triangle.add(Arrays.asList(2));
        triangle.add(Arrays.asList(3, 4));
        triangle.add(Arrays.asList(6, 5, 7));
        triangle.add(Arrays.asList(4, 1, 8, 3));
        System.out.println(q.new Solution().minimumTotal(triangle));
        System.out.println(q.new Solution1().minimumTotal(triangle));
        System.out.println(q.new Solution2().minimumTotal(triangle));
    }
}

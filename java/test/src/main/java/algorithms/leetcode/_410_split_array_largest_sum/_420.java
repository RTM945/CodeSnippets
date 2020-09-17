package algorithms.leetcode._410_split_array_largest_sum;

import java.util.Arrays;

//https://leetcode-cn.com/problems/split-array-largest-sum
//给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。
//设计一个算法使得这 m 个子数组各自和的最大值最小。
//注意:
//数组长度 n 满足以下条件:
//1 ≤ n ≤ 1000
//1 ≤ m ≤ min(50, n)
//示例:
//输入:
//nums = [7,2,5,10,8]
//m = 2
//输出:
//18
//解释:
//一共有四种方法将nums分割为2个子数组。
//其中最好的方式是将其分为[7,2,5] 和 [10,8]，
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
public class _420 {
    class Solution {
        //Hard
        //连续子数组
        //想象成有m - 1个分隔符，比如m0, m1 ... mn, n < m
        //把数组分为[0, m0), [m0, m1), ... ,[mn, nums.length)这么多子数组
        //假设这些子数组各自的和的最大值最小
        //(m0 - 0) + (m1 - m0) + ... + (nums.length - mn) = nums.length
        //好像没啥意义...
        //换个思路
        //画了个图
        //设dp[i][j]为i到j子区间的和
        //nums = [7,2,5,10,8] m = 3
        //则有
        //  0  1  2  3  4
        //0 7  9  14
        //1    2  7  17
        //2       5  15 23
        //3          10 18
        //4             8
        //看上去是从每一行中找到最大的元素，在所有最大的元素中找到最小的元素
        public int splitArray(int[] nums, int m) {
            int len = nums.length;
            int res = Integer.MAX_VALUE;
            if(m == 1) {
                res = 0;
                for (int i = 0; i < len; i++) {
                    res += nums[i];
                }
                return res;
            }
            if(m == len) {
                res = 0;
                for (int i = 0; i < len; i++) {
                    res = Math.max(res, nums[i]);
                }
                return res;
            }
            for (int i = 0; i < len - 1; i++) {
                int x = Math.min(len, i + len - m + 1);
                int rmax = 0;
                for (int j = i; j < x; j++) {
                    rmax += nums[j];
                }
                res = Math.min(res, rmax);
            }
            return res;
        }
    }

    class Solution1 {
        //上面还是翻车了
        //[2,3,1,2,4,3]正确答案是4，[2],[3],[1,2],[4],[3]的情况
        //我的dp确实能画出所有情况，但取结果的条件不对
        //取成了3
        //看看答案
        //先dp，设dp[i][j]为前i个数分割为j段所能得到的最大连续子数组和的最小值
        //在前i个数中，假设前k个数被分成了j - 1段，那么[k + 1, i]为第j段
        //dp[i][j] = max(dp[k][j - 1] + sub(k + 1, i)) 0<=k<=i-1
        //应该i >= j才能分出j个来，对于dp中i < j的情况，将值初始化成max
        //当j = 1时，前i个数分成1段，dp[0][0] = 0. dp[k][0]不合法
        public int splitArray(int[] nums, int m) {
            int n = nums.length;
            int[][] dp = new int[n + 1][m + 1];
            for (int i = 0; i <= n; i++) {
                Arrays.fill(dp[i], Integer.MAX_VALUE);
            }
            int[] sub = new int[n + 1];
            for (int i = 0; i < n; i++) {
                sub[i + 1] = sub[i] + nums[i];
            }
            dp[0][0] = 0;
            for (int i = 1; i <= n; i++) {
                for (int j = 1; j <= Math.min(i, m); j++) {
                    for (int k = 0; k < i; k++) {
                        dp[i][j] = Math.min(dp[i][j], Math.max(dp[k][j - 1], sub[i] - sub[k]));
                    }
                }
            }
            return dp[n][m];
        }
    }

    class Solution2{
        //上面不是很明白
        //还有二分查找做法
        //结果的范围在[max(nums), sum(nums)]
        //这个max(nums)让我重新审视了“m 个子数组各自和的最大值最小”
        //假设全分开，各个数组最大值肯定是max(nums)
        //假设部分分开，那么和的最大值肯定是max(nums) + 相邻的元素
        //现在假设结果是x，知道了x的范围，使用二分查找去逼近它
        //如何验证二分查找出来的是正确结果？
        //对于二分查找mid，如果是正确答案，则每个子数组的和都小于等于mid
        //毕竟他是最大值
        //使用这个mid从头开始分割数组，令数组和<=mid，一旦大于了，就开始一个新的数组
        //这样分割出来的子数组个数再与m比较，如果比m多了，说明mid小了，去右半重新找
        //如果比m少了，说明mid太大，去左半重新找
        public int splitArray(int[] nums, int m) {
            if (nums == null || nums.length == 0 || m == 0) {
                return 0;
            }
            int left = 0;
            int right = 0;
            for (int i = 0; i < nums.length; i++) {
                left = Math.max(left, nums[i]);
                right += nums[i];
            }
            while (left <= right) {
                int mid = left + (right - left) / 2;
                int t = check(nums, mid);
                if(t > m) {
                    left = mid + 1;
                }else{
                    right = mid - 1;
                }
            }
            return left;
        }

        private int check(int[] nums, int mid) {
            int sum = 0;
            int m = 1; //子数组最少1个
            for (int i = 0; i < nums.length; i++) {
                sum += nums[i];
                if (sum > mid) {
                    sum = nums[i]; //新开一个子数组
                    m++;
                }
            }
            return m;
        }
    }

    public static void main(String[] args) {
        _420 q = new _420();
        System.out.println(q.new Solution2().splitArray(new int[]{7,2,5,10,8}, 2));
        System.out.println(q.new Solution2().splitArray(new int[]{7,2,5,10,8}, 3));
    }
}

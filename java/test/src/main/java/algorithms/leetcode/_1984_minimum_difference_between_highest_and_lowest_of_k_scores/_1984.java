package algorithms.leetcode._1984_minimum_difference_between_highest_and_lowest_of_k_scores;

import java.util.Arrays;

public class _1984 {
    // 给你一个 下标从 0 开始 的整数数组 nums ，其中 nums[i] 表示第 i 名学生的分数。
    // 另给你一个整数 k 。

    // 从数组中选出任意 k 名学生的分数，使这 k 个分数间 最高分 和 最低分 的 差值 达到 最小化 。

    // 返回可能的 最小差值 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 滑动窗口？
    class Solution {
        public int minimumDifference(int[] nums, int k) {
            int res = Integer.MAX_VALUE;
            Arrays.sort(nums);
            for (int i = 0; i + k - 1 < nums.length; i++) {
                res = Math.min(res, nums[i + k - 1] - nums[i]);
            }
            return res;
        }
    }
}

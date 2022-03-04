package algorithms.leetcode._2104_sum_of_subarray_ranges;

import org.checkerframework.checker.units.qual.min;

public class _2104 {
    // 给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。

    // 返回 nums 中 所有 子数组范围的 和 。

    // 子数组是数组中一个连续 非空 的元素序列。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/sum-of-subarray-ranges
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public long subArrayRanges(int[] nums) {
            long res = 0;
            for (int i = 0; i < nums.length; i++) {
                long min = nums[i];
                long max = nums[i];
                for (int j = i + 1; j < nums.length; j++) {
                    min = Math.min(min, nums[j]);
                    max = Math.max(max, nums[j]);
                    res += max - min;
                }
            }
            return res;
        }
    }
}

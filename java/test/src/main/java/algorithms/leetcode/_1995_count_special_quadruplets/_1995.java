package algorithms.leetcode._1995_count_special_quadruplets;

public class _1995 {
    // 给你一个 下标从 0 开始 的整数数组 nums ，
    // 返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
    // nums[a] + nums[b] + nums[c] == nums[d] ，且
    // a < b < c < d
    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/count-special-quadruplets
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 4 <= nums.length <= 50
    // 1 <= nums[i] <= 100

    // 暴力
    class Solution {
        int countQuadruplets(int[] nums) {
            int result = 0;
            for (int i = 0; i < nums.length; ++i) {
                for (int j = i + 1; j < nums.length; ++j) {
                    for (int k = j + 1; k < nums.length; ++k) {
                        for (int h = k + 1; h < nums.length; ++h) {
                            if (nums[i] + nums[j] + nums[k] == nums[h])
                                result++;
                        }
                    }
                }
            }
            return result;
        }
    }
}

package algorithms.leetcode._1995_count_special_quadruplets;

import java.util.HashMap;
import java.util.Map;

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

    // a b c 确定时 d 也就确定了 可以少一个循环
    class Solution1 {
        int countQuadruplets(int[] nums) {
            int n = nums.length;
            int ans = 0;
            Map<Integer, Integer> cnt = new HashMap<>();
            // 假设最后一位是 d 倒数第二位是 c 从倒数第二位向前遍历到假定的 b 第二位
            for (int c = n - 2; c >= 2; --c) {
                // 这段都有可能为 d 记录下来
                cnt.put(nums[c + 1], cnt.getOrDefault(nums[c + 1], 0) + 1);
                for (int a = 0; a < c; ++a) {
                    for (int b = a + 1; b < c; ++b) {
                        ans += cnt.getOrDefault(nums[a] + nums[b] + nums[c], 0);
                    }
                }
            }
            return ans;
        }
    }

    // a b 确定时 d - c 也就确定了 可以再少一个循环
    class Solution2 {
        int countQuadruplets(int[] nums) {
            int n = nums.length;
            int ans = 0;
            Map<Integer, Integer> cnt = new HashMap<>();
            // 假设最后一位是 d 倒数第三位是 b 从倒数第二位向前遍历到假定的 a 第一位
            for (int b = n - 3; b >= 1; --b) {
                for (int d = b + 2; d < n; ++d) {
                    cnt.put(nums[d] - nums[b + 1], cnt.getOrDefault(nums[d] - nums[b + 1], 0) + 1);
                }
                for (int a = 0; a < b; ++a) {
                    ans += cnt.getOrDefault(nums[a] + nums[b], 0);
                }
            }
            return ans;
        }
    }
}

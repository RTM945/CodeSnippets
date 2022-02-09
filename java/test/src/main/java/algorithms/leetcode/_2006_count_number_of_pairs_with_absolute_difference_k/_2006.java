package algorithms.leetcode._2006_count_number_of_pairs_with_absolute_difference_k;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class _2006 {
// 给你一个整数数组 nums 和一个整数 k ，请你返回数对 (i, j) 的数目，
// 满足 i < j 且 |nums[i] - nums[j]| == k 。

// |x| 的值定义为：

// 如果 x >= 0 ，那么值为 x 。
// 如果 x < 0 ，那么值为 -x 。
//  

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public int countKDifference(int[] nums, int k) {
            int ans = 0;
            for (int i = 0; i < nums.length - 1; i++) {
                for (int j = i + 1; j < nums.length; j++) {
                    if (Math.abs(nums[i] - nums[j]) == k) {
                        ans++;
                    }
                }
            }

            return ans;
        }
    }

    class Solution1 {
        public int countKDifference(int[] nums, int k) {
            int ans = 0;
            Map<Integer, Integer> cnt = new HashMap<>();
            for (int i = 0; i < nums.length; i++) {
                ans += cnt.getOrDefault(nums[i] - k, 0) + cnt.getOrDefault(nums[i] + k, 0);
                cnt.put(nums[i], cnt.getOrDefault(nums[i], 0) + 1);
            }

            return ans;
        }
    }
}

package algorithms.leetcode._219_contains_duplicate_ii;

import java.util.HashMap;
import java.util.Map;

import org.checkerframework.checker.units.qual.K;

public class _219 {
    // 给你一个整数数组 nums 和一个整数 k ，
    // 判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
    // 如果存在，返回 true ；否则，返回 false 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/contains-duplicate-ii
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public boolean containsNearbyDuplicate(int[] nums, int k) {
            Map<Integer, Integer> map = new HashMap<>();
            for (int i = 0; i < nums.length; i++) {
                if (map.containsKey(nums[i])) {
                    if (Math.abs(i - map.get(nums[i])) <= k) {
                        return true;
                    } else {
                        map.put(nums[i], i);
                    }
                } else {
                    map.put(nums[i], i);
                }
            }
            return false;
        }
    }

    public static void main(String[] args) {
        _219 _219 = new _219();
        Solution solution = _219.new Solution();
        System.out.println(solution.containsNearbyDuplicate(new int[] {1,2,3,1}, 3));
        System.out.println(solution.containsNearbyDuplicate(new int[] {1,0,1,1}, 3));
        System.out.println(solution.containsNearbyDuplicate(new int[] {1,2,3,1,2,3}, 2));
    }
}

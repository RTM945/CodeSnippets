package algorithms.leetcode._540_single_element_in_a_sorted_array;

public class _540 {
    // 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

    // 请你找出并返回只出现一次的那个数。

    // 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/single-element-in-a-sorted-array
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 二分查找？
    class Solution {
        public int singleNonDuplicate(int[] nums) {
            int left = 0;
            int right = nums.length - 1;
            while (left < right) {
                // [1, 1, 2, 2, 3]
                int mid = left + (right - left) / 2;
                // 奇数下标应该跟前一个相等 mid - 1 = mid ^ 1
                // 偶数下标应该跟后一个相等 mid + 1 = mid ^ 1
                // 相等往后找 否则往前找
                if (nums[mid] == nums[mid ^ 1]) {
                    left = mid + 1;
                } else {
                    right = mid;
                }
            }
            return nums[left];
        }
    }

    public static void main(String[] args) {
        _540 _540 = new _540();
        Solution solution = _540.new Solution();
        System.out.println(solution.singleNonDuplicate(new int[] {1,1,2,3,3,4,4,8,8}));
        System.out.println(solution.singleNonDuplicate(new int[] {3,3,7,7,10,11,11}));
    }
}

package algorithms.leetcode._334_increasing_triplet_subsequence;

public class _334 {
    // 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

    // 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，
    // 使得 nums[i] < nums[j] < nums[k] ，返回 true ；
    // 否则，返回 false 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/increasing-triplet-subsequence
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 暴力？
    class Solution {
        public boolean increasingTriplet(int[] nums) {
            for (int i = 0; i < nums.length - 2; i++) {
                for (int j = i + 1; j < nums.length - 1; j++) {
                    for (int k = j + 1; k < nums.length; k++) {
                        if (nums[i] < nums[j] && nums[j] < nums[k]) {
                            return true;
                        }
                    }
                }
            }
            return false;
        }
    }

    // 在nums[i]左边存在小于nums[i]的等价于nums[i]左边的最小值小于nums[i]
    // 在nums[i]右边存在大于nums[i]的等价于nums[i]右边的最大值大于nums[i]
    // 分别记录下每个nums[i]左边最小的数和右边最大的数
    class Solution1 {
        public boolean increasingTriplet(int[] nums) {
            int n = nums.length;
            if (n < 3) {
                return false;
            }
            int[] leftMin = new int[n];
            leftMin[0] = nums[0];
            for (int i = 1; i < n; i++) {
                leftMin[i] = Math.min(leftMin[i - 1], nums[i]);
            }
            int[] rightMax = new int[n];
            rightMax[n - 1] = nums[n - 1];
            for (int i = n - 2; i >= 0; i--) {
                rightMax[i] = Math.max(rightMax[i + 1], nums[i]);
            }
            for (int i = 1; i < n - 1; i++) {
                if (nums[i] > leftMin[i - 1] && nums[i] < rightMax[i + 1]) {
                    return true;
                }
            }
            return false;
        }
    }

    // 贪心：为了找到递增的三元子序列，first 和 second 应该尽可能地小，
    // 此时找到递增的三元子序列的可能性更大。
    class Solution2 {
        public boolean increasingTriplet(int[] nums) {
            int n = nums.length;
            if (n < 3) {
                return false;
            }
            int first = nums[0], second = Integer.MAX_VALUE;
            for (int i = 1; i < n; i++) {
                int num = nums[i];
                if (num > second) {
                    return true;
                } else if (num > first) {
                    second = num;
                } else {
                    first = num;
                }
            }
            return false;
        }
    }

    public static void main(String[] args) {
        _334 _334 = new _334();
        Solution solution = _334.new Solution();
        System.out.println(solution.increasingTriplet(new int[] {1,2,3,4,5}));
        System.out.println(solution.increasingTriplet(new int[] {5,4,3,2,1}));
        System.out.println(solution.increasingTriplet(new int[] {2,1,5,0,4,6}));
    }

}

package algorithms.leetcode._1005_maximize_sum_of_array_after_k_negations;

import java.util.Arrays;

public class _1005 {
    // 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
    // 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
    // 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
    // 以这种方式修改数组后，返回数组 可能的最大和 。
    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        // 一眼看过去，把所有负数变成正数，应该就可以得到最大和了
        // 但限制了取相反数的次数，所以要找出数组中前k小的负数，将它们变成正数
        // 当负数的个数小与k的时候，剩下的要找到最小的正数，这样取相反数对最大和的影响最小
        // 所以第一步还是排序
        public int largestSumAfterKNegations(int[] nums, int k) {
            Arrays.sort(nums);
            int sum = 0;
            for (int i = 0; i < nums.length; i++) {
                if (k > 0 && nums[i] < 0) {
                    nums[i] = -nums[i];
                    k--;
                }
                sum += nums[i];
            }
            Arrays.sort(nums);
            // 如果k没剩下返回sum
            // k有剩下，那必定全是正数
            // k剩下偶数，那么用剩下的首位加减抵消
            // k剩下奇数，需要减去2*nums[0]，因为在前面加过了
            return sum - (k % 2 == 0 ? 0 : 2 * nums[0]); 
        }
    }

    public static void main(String[] args) {
        _1005 _1005 = new _1005();
        Solution solution = _1005.new Solution();
        System.out.println(solution.largestSumAfterKNegations(new int[]{3,-1,1,2}, 4));
    }
}

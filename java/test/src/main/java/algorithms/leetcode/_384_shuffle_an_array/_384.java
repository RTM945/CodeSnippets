package algorithms.leetcode._384_shuffle_an_array;

import java.util.Arrays;
import java.util.Random;

// https://leetcode-cn.com/problems/shuffle-an-array
// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
// 实现 Solution class:
// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果

public class _384 {
    class Solution {

        int[] origin;
        Random r;

        public Solution(int[] nums) {
            origin = nums;
            r = new Random();
        }
        
        public int[] reset() {
            return origin;
        }
        
        public int[] shuffle() {
            // 从后往前，将当前元素与前面的随机元素做调换
            int[] nums = Arrays.copyOf(origin, origin.length);
            for (int i = origin.length; i > 1; i--) {
                int index = i - 1;
                int temp = nums[index];
                int swapIndex = r.nextInt(i);
                nums[index] = nums[swapIndex];
                nums[swapIndex] = temp;
            }
            return nums;
        }
    }

    public static void main(String[] args) {
        _384 _384 = new _384();
        Solution solution = _384.new Solution(new int[]{1, 2, 3});
        System.out.println(Arrays.toString(solution.shuffle()));
        System.out.println(Arrays.toString(solution.reset()));
        System.out.println(Arrays.toString(solution.shuffle()));
    }
}

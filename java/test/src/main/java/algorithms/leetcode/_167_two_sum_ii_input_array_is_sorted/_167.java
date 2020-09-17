package algorithms.leetcode._167_two_sum_ii_input_array_is_sorted;

import java.util.Arrays;

//https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//说明:
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//示例:
//输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
public class _167 {
    //easy
    //一瞬间的想法是二分查找，找到应该插入的位置
    //在该位置前找到比他小的元素再判断他们的和
    //前面最简单的两数之和问题用的hashmap，但题目中没说每个数字只出现一次，所以没法用hashmap
    class Solution {
        public int[] twoSum(int[] numbers, int target) {
            int len = numbers.length;
            int left = 0;
            int right = len - 1;
            while(left <= right) {
                int mid = left + (right - left) / 2;
                if(target > numbers[mid]) {
                    left = mid + 1;
                }else if(target < numbers[mid]) {
                    right = mid - 1;
                }else{
                    left = mid;
                    break;
                }
            }
            if(left <= 1) {
                return new int[]{1, 2};
            }
            for (int i = 0; i < left - 1; i++) {
                int sub = target - numbers[i];
                for (int j = i + 1; j < left; j++) {
                    if(numbers[j] == sub) {
                        return new int[]{i + 1, j + 1};
                    }
                }
            }
            return new int[]{};
        }
    }

    class Solution1{
        //智障了
        //可以直接找第二个数
        public int[] twoSum(int[] numbers, int target) {
            for (int i = 0; i < numbers.length; i++) {
                int left = i + 1;
                int right = numbers.length - 1;
                int sub = target - numbers[i];
                while(left <= right) {
                    int mid = left + (right - left) / 2;
                    if(sub > numbers[mid]) {
                        left = mid + 1;
                    }else if(sub < numbers[mid]) {
                        right = mid - 1;
                    }else{
                        return new int[]{i + 1, mid + 1};
                    }
                }
            }
            return new int[]{-1, -1};
        }
    }

    class Solution2{
        //双指针法
        public int[] twoSum(int[] numbers, int target) {
            int left = 0;
            int right = numbers.length - 1;
            while (left < right) {
                int sum = numbers[left] + numbers[right];
                if(sum < target) {
                    left++;
                }else  if(sum > target) {
                    right--;
                }else{
                    return new int[]{left + 1, right + 1};
                }
            }
            return new int[]{-1, -1};
        }
    }

    public static void main(String[] args) {
        _167 q = new _167();
        System.out.println(Arrays.toString(q.new Solution().twoSum(new int[]{2, 7, 11, 15}, 9)));
        System.out.println(Arrays.toString(q.new Solution().twoSum(new int[]{2, 3, 4}, 6)));
        System.out.println(Arrays.toString(q.new Solution().twoSum(new int[]{-1, 0}, -1)));
        System.out.println(Arrays.toString(q.new Solution().twoSum(new int[]{0, 0, 3, 4}, 0)));
    }
}

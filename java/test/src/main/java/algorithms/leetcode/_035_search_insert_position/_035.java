package algorithms.leetcode._035_search_insert_position;

import java.util.*;

//https://leetcode-cn.com/problems/search-insert-position
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
//如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//你可以假设数组中无重复元素。
//示例 1:
//输入: [1,3,5,6], 5
//输出: 2
//示例 2:
//输入: [1,3,5,6], 2
//输出: 1
//示例 3:
//输入: [1,3,5,6], 7
//输出: 4
//示例 4:
//输入: [1,3,5,6], 0
//输出: 0
public class _035 {

    class Solution {
        //二分查找插入啦
        public int searchInsert(int[] nums, int target) {
            int len = nums.length;
            int left = 0;
            int right = len - 1;
            while(left <= right) {
                int mid = left + (right - left) / 2;
                if(target > nums[mid]) {
                    left = mid + 1;
                } else if(target < nums[mid]) {
                    right = mid - 1;
                } else {
                    return mid;
                }
            }
            return left;
        }
    }

    class Testcase{
        int[] nums;
        int target;
        int want;

        public Testcase(int[] nums, int target, int want) {
            this.nums = nums;
            this.target = target;
            this.want = want;
        }

        @Override
        public String toString() {
            return "Testcase{" +
                    "nums=" + Arrays.toString(nums) +
                    ", target=" + target +
                    ", want=" + want +
                    '}';
        }
    }

    public static void main(String[] args) {
        _035 q = new _035();
        List<Testcase> list = new ArrayList<>();
        list.add(q.new Testcase(new int[]{1, 3, 5, 6}, 5, 2));
        list.add(q.new Testcase(new int[]{1, 3, 5, 6}, 2, 1));
        list.add(q.new Testcase(new int[]{1, 3, 5, 6}, 7, 4));
        list.add(q.new Testcase(new int[]{1, 3, 5, 6}, 0, 0));
        for(Testcase t : list) {
            int result = q.new Solution().searchInsert(t.nums, t.target);
            if(result != t.want) {
                System.err.println(t + " result = " + result);
            }
        }
    }
}

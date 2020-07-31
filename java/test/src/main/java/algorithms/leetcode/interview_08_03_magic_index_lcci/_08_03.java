package algorithms.leetcode.interview_08_03_magic_index_lcci;

// 魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。
// 给定一个有序整数数组，编写一种方法找出魔术索引，
// 若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。
// 若有多个魔术索引，返回索引值最小的一个。
// 示例1:
// 输入：nums = [0, 2, 3, 4, 5]
// 输出：0
// 说明: 0下标的元素为0
// 示例2:
// 输入：nums = [1, 1, 1]
// 输出：1
// 说明:
// nums长度在[1, 1000000]之间
// 此题为原书中的 Follow-up，即数组中可能包含重复元素的版本
public class _08_03 {
    class Solution {
        //easy
        //谜之题目
        //从前往后遍历，索引逐渐变大的，那么只要找到第一个nums[i] = i返回不就行了？
        public int findMagicIndex(int[] nums) {
            for (int i = 0; i < nums.length; i++) {
                if(i == nums[i]) {
                    return i;
                }
            }

            return -1;
        }

        //据说是配合另一道二分的题目做的，有O(logn)的解法
        //maybe next time
    }
}
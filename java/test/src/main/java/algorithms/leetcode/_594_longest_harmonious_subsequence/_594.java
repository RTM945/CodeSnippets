package algorithms.leetcode._594_longest_harmonious_subsequence;

import java.util.Map;
import java.util.Arrays;
import java.util.HashMap;

public class _594 {
    // https://leetcode-cn.com/problems/longest-harmonious-subsequence
    // 和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
    // 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
    // 数组的子序列是一个由数组派生出来的序列，
    // 它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。

    class Solution {
        // 简单的想法是找到数组中出现次数最多的元素
        // 再找到和他相差1的元素
        // [1,4,1,3,1,-14,1,-13] 错了
        public int findLHS(int[] nums) {
            Map<Integer, Integer> map = new HashMap<>();
            for (int i : nums) {
                Integer v = map.get(i);
                if (v == null) {
                    v = 0;
                }
                map.put(i, v + 1);
            }
            int maxTime = 0;
            int c = 0;
            for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
                int k = entry.getKey();
                int v = entry.getValue();
                // 但没考虑最大出现次数相等的情况
                if (v > maxTime) {
                    maxTime = v;
                    c = k;
                }
            }
            int minsOne = c - 1;
            int plusOne = c + 1;
            int res1 = 0;
            int res2 = 0;
            for (int i : nums) {
                if (i == minsOne) {
                    res1++;
                }
                if (i == plusOne) {
                    res2++;
                }
            }
            if (res1 != 0) {
                res1 += maxTime;
            }
            if (res2 != 0) {
                res2 += maxTime;
            }
            return Math.max(res1, res2);
        }
    }

    class Solution1 {
        public int findLHS(int[] nums) {
            // 排序之后滑动窗口
            Arrays.sort(nums);
            int begin = 0, res = 0;
            for (int end = 0; end < nums.length; end++) {
                while (nums[end] - nums[begin] > 1) {
                    begin++;
                }
                
                if (nums[end] - nums[begin] == 1) {
                    res = Math.max(res, end - begin + 1);
                }
            }
            return res;
        }
    }

    class Solution2 {
        public int findLHS(int[] nums) {
            HashMap<Integer, Integer> map = new HashMap<>();
            int res = 0;
            for (int num : nums) {
                map.put(num, map.getOrDefault(num, 0) + 1);
            }
            for (int key : map.keySet()) {
                if (map.containsKey(key + 1)) {
                    res = Math.max(res, map.get(key) + map.get(key + 1));
                }
            }
            return res;
        }
    }

    public static void main(String[] args) {
        _594 _594 = new _594();
        Solution solution = _594.new Solution();
        System.out.println(solution.findLHS(new int[] { 1, 4, 1, 3, 1, -14, 1, -13 }));
    }
}

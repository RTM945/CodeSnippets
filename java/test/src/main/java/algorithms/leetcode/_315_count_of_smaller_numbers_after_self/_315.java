package algorithms.leetcode._315_count_of_smaller_numbers_after_self;

import java.util.*;
import java.util.stream.Collectors;

//https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/
//给定一个整数数组 nums，按要求返回一个新数组 counts。
//数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
//示例:
//输入: [5,2,6,1]
//输出: [2,1,1,0]
//解释:
//5 的右侧有 2 个更小的元素 (2 和 1).
//2 的右侧仅有 1 个更小的元素 (1).
//6 的右侧有 1 个更小的元素 (1).
//1 的右侧有 0 个更小的元素.
public class _315 {

    //困难题目，慌的一批
    //但看上去用暴力解法好理解
    class Solution {
        public List<Integer> countSmaller(int[] nums) {
            List<Integer> list = new ArrayList<>();
            for (int i = 0; i < nums.length; i++) {
                int big = 0;
                for (int j = i + 1; j < nums.length; j++) {
                    if(nums[j] < nums[i]) {
                        big++;
                    }
                }
                list.add(big);
            }
            return list;
        }
    }

    //暴力解法超时，分析超时的原因
    //其实有很多很多的重复比较
    //那么又可以用dp了
    //二维数组，但如果nums很大，二维数组会很浪费空间
    //或许可以用一维数组解决
    //dp[i]表示第i位右边比num[i]大的数字
    //dp[nums.length - 1] = 0
    //dp[i - 1] = dp[i] + 1 (nums[i - 1] > nums[i])
    //如果nums[i - 1]不大于nums[i]，就要从nums第i位起找到一个大于的，用它的dp值+1
    //对于[2, 0, 1]的情况，这个解法错了
    //dp[2]是0，dp[1]因为0 < 1也是0，dp[1]无法得出2
    class Solution1 {
        public List<Integer> countSmaller(int[] nums) {
            if(nums == null || nums.length == 0) {
                return Collections.emptyList();
            }
            int[] dp = new int[nums.length];
            dp[nums.length - 1] = 0;
            for (int i = nums.length - 1; i >= 1; i--) {
                if(nums[i - 1] > nums[i]) {
                    dp[i - 1] = dp[i] + 1;
                }else{
                    for (int j = i + 1; j < nums.length; j++) {
                        if(nums[i - 1] > nums[j]) {
                            dp[i - 1] = dp[j] + 1;
                        }
                    }
                }
            }
            return Arrays.stream(dp).boxed().collect(Collectors.toList());
        }
     }

    class Solution2 {
        //答案中说到插入排序
        //忽略最后一个元素
        //从倒数第二位开始插入排序
        //与它后面的数进行比较，进行正序排序
        //最后移动了多少位则有多少比他小
        public List<Integer> countSmaller(int[] nums) {
            if(nums == null || nums.length == 0) {
                return Collections.emptyList();
            }
            Integer[] res = new Integer[nums.length];
            res[nums.length - 1] = 0;
            for (int i = nums.length - 2; i >= 0; i--) {
                int x = nums[i];
                int j = i + 1;
                while (j < nums.length && nums[j] >= x) {
                    nums[j - 1] = nums[j];
                    j++;
                }
                nums[j - 1] = x;
                res[i] = nums.length - j;
            }
            return Arrays.asList(res);
        }
    }

    class Solution3 {
        //插入排序的优化思路就是二分查找了
        public List<Integer> countSmaller(int[] nums) {
            if(nums == null || nums.length == 0) {
                return Collections.emptyList();
            }
            Integer[] res = new Integer[nums.length];
            res[nums.length - 1] = 0;
            for (int i = nums.length - 2; i >= 0; i--) {
                int x = nums[i];
                //二分查找
                //找到比x小的数
                int low = i + 1;
                int high = nums.length - 1;
                int mid = low + (high - low) / 2;
                while(low <= high) {
                    if(nums[mid] < x) {
                        high = mid - 1;
                    }else{
                        low = mid + 1;
                    }
                }
                //找到了比x小的，那么x的位置就在它的前面
                for (int j = i - 1; j < low - 1; j++) {
                    nums[j] = nums[j + 1];
                }
                nums[low - 1] = x;
                res[i] = nums.length - low + 1;
            }
            return Arrays.asList(res);
        }
    }

    class Solution4 {
        public List<Integer> countSmaller(int[] nums) {
            int n = nums.length;
            if (n < 1) return new ArrayList<Integer>();
            Integer[] res = new Integer[n];
            res[n - 1] = 0;
            for (int i = n - 2; i >= 0; i--) {
                int posit = binarySearch(nums, i + 1, n - 1, nums[i]) - 1;
                int pivot = nums[i];
                System.arraycopy(nums, i + 1, nums, i, posit - i);
                nums[posit] = pivot;
                res[i] = posit - i;
            }
            return Arrays.asList(res);
        }

        private int binarySearch(int[] nums, int low, int hig, int key) {
            while (low <= hig) {
                int mid = low + ((hig - low) >> 1);
                if (nums[mid] < key) low = mid + 1;
                else hig = mid - 1;
            }
            return low;
        }
    }

    public static void main(String[] args) {
        _315 q = new _315();
//        System.out.println(q.new Solution().countSmaller(new int[]{5,2,6,1}));
//        System.out.println(q.new Solution1().countSmaller(new int[]{5,2,6,1}));
//        System.out.println(q.new Solution1().countSmaller(new int[]{2,0,1}));
//        System.out.println(q.new Solution3().countSmaller(new int[]{5,2,6,1}));
        System.out.println(q.new Solution4().countSmaller(new int[]{5,2,6,1}));
    }
}

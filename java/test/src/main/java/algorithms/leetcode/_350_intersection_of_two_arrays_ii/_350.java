package algorithms.leetcode._350_intersection_of_two_arrays_ii;

import java.util.*;
import java.util.function.BiFunction;

//https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9]
//说明：
//输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//我们可以不考虑输出结果的顺序。
//进阶:
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
public class _350 {
    //输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致
    //为了达到这一目的,一般的想法是遍历到相等的元素删除
    //但为此还要转换成List,嫌烦就算了
    //于是用了个数组去存标志位,如果匹配到了标志为1,确保不重复放入结果数组
    class Solution {
        public int[] intersect(int[] nums1, int[] nums2) {
            int[] mark = new int[nums2.length];
            List<Integer> list = new ArrayList<>();
            for (int i = 0; i < nums1.length; i++) {
                for (int j = 0; j < nums2.length; j++) {
                    if(nums1[i] == nums2[j] && mark[j] == 0) {
                        list.add(nums2[j]);
                        mark[j] = 1;
                        break;
                    }
                }
            }
            int[] res = new int[list.size()];
            for (int i = 0; i < res.length; i++) {
                res[i] = list.get(i);
            }
            return res;
        }
    }

    class Solution1 {
        //hashmap的做法
        //用hashmap存nums1每一位和出现的次数
        //遍历nums2,如果匹配了hashmap的key,次数减少1,到0的时候,就不能加入结果集了
        public int[] intersect(int[] nums1, int[] nums2) {
            if (nums1.length > nums2.length) {
                return intersect(nums2, nums1);
            }
            Map<Integer, Integer> map = new HashMap<Integer, Integer>();
            for (int num : nums1) {
                int count = map.getOrDefault(num, 0) + 1;
                map.put(num, count);
            }
            int[] intersection = new int[nums1.length];
            int index = 0;
            for (int num : nums2) {
                int count = map.getOrDefault(num, 0);
                if (count > 0) {
                    intersection[index++] = num;
                    count--;
                    if (count > 0) {
                        map.put(num, count);
                    } else {
                        map.remove(num);
                    }
                }
            }
            return Arrays.copyOfRange(intersection, 0, index);
        }
    }

    class Solution2{
        //排序的方法
        //有点像归并
        public int[] intersect(int[] nums1, int[] nums2) {
            Arrays.sort(nums1);
            Arrays.sort(nums2);
            int length1 = nums1.length, length2 = nums2.length;
            int[] intersection = new int[Math.min(length1, length2)];
            int index1 = 0, index2 = 0, index = 0;
            while (index1 < length1 && index2 < length2) {
                if (nums1[index1] < nums2[index2]) {
                    index1++;
                } else if (nums1[index1] > nums2[index2]) {
                    index2++;
                } else {
                    intersection[index] = nums1[index1];
                    index1++;
                    index2++;
                    index++;
                }
            }
            return Arrays.copyOfRange(intersection, 0, index);
        }
    }

    class TestCase{
        int[] a;
        int[] b;
        int[] want;

        public TestCase(int[] a, int[] b, int[] want) {
            this.a = a;
            this.b = b;
            this.want = want;
        }
    }

    public static void main(String[] args) {
        _350 q = new _350();
        List<TestCase> testcases = new ArrayList<>();
        testcases.add(q.new TestCase(new int[]{1, 2, 2, 1}, new int[]{2, 2}, new int[]{2, 2}));
        testcases.add(q.new TestCase(new int[]{3, 1, 2}, new int[]{1, 1}, new int[]{1}));
        testcases.add(q.new TestCase(new int[]{4, 9, 5}, new int[]{9, 4, 9, 8, 4}, new int[]{4, 9}));
        for (TestCase testCase : testcases) {
            int[] res = q.new Solution1().intersect(testCase.a, testCase.b);
            Arrays.sort(res);
            if (!Arrays.equals(res, testCase.want)) {
                System.err.println(Arrays.toString(res) + " != " + Arrays.toString(testCase.want));
            }
        }
    }
}

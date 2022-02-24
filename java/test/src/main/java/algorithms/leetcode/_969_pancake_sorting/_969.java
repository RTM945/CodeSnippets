package algorithms.leetcode._969_pancake_sorting;

import java.util.ArrayList;
import java.util.List;

public class _969 {
    // 给你一个整数数组 arr ，请使用 煎饼翻转 完成对数组的排序。

    // 一次煎饼翻转的执行过程如下：

    // 选择一个整数 k ，1 <= k <= arr.length
    // 反转子数组 arr[0...k-1]（下标从 0 开始）
    // 例如，arr = [3,2,1,4] ，选择 k = 3 进行一次煎饼翻转，
    // 反转子数组 [3,2,1] ，得到 arr = [1,2,3,4] 。

    // 以数组形式返回能使 arr 有序的煎饼翻转操作所对应的 k 值序列。
    // 任何将数组排序且翻转次数在 10 * arr.length 范围内的有效答案都将被判断为正确。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/pancake-sorting
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 先找到最大的数 从他开始翻转 再数组全部翻转 让最大的数在最后
    class Solution {
        public List<Integer> pancakeSort(int[] arr) {
            List<Integer> res = new ArrayList<>();
            int len = arr.length;
            while (len > 1) {
                // 找出最大值
                int maxIndex = 0;
                int max = arr[maxIndex];
                for (int i = 1; i < len; i++) {
                    if (arr[i] > max) {
                        max = arr[i];
                        maxIndex = i;
                    }
                }
                if (maxIndex == len - 1) {
                    // 最大值已经是最后一位了
                    // 不需要翻转 找第二大的
                    len--;
                    continue;
                }
                res.add(maxIndex + 1);
                // 第一次煎饼
                pancake(arr, maxIndex + 1);
                res.add(len);
                // 第二次煎饼
                pancake(arr, len);
                len--;
            }
            return res;
        }

        private void pancake(int[] arr, int index) {
            int right = index - 1;
            int left = 0;
            while (left < right) {
                int tmp = arr[left];
                arr[left++] = arr[right];
                arr[right--] = tmp;
            }
        }
    }
}

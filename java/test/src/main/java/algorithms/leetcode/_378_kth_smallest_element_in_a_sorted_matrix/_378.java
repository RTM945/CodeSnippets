package algorithms.leetcode._378_kth_smallest_element_in_a_sorted_matrix;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

//给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
//示例：
//matrix = [
//  [ 1,  5,  9],
//  [10, 11, 13],
//  [12, 13, 15]
//],
//k = 8,
//返回 13。
//提示：
//你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。
//https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
public class _378 {
    //一开始思路是行和列分别找到第k个，再比较两个k大小
    //再一想这是n x n不是k x k，思路有问题
    //如果把矩阵展开成一维数组再排序获得第k个元素呢
    //排序的消耗可能很大
    class Solution {
        public int kthSmallest(int[][] matrix, int k) {
            List<Integer> list = new ArrayList<>();
            for (int i = 0; i < matrix.length; i++) {
                for (int j = 0; j < matrix[i].length; j++) {
                    list.add(matrix[i][j]);
                }
            }
            Collections.sort(list);
            return list.get(k - 1);
        }
    }

    class Solution2 {
        //可不可以使用已经有序的行与列呢
        //归并排序中有个做法，两个集合，当A[i] < B[j]时，取A[i]，i++，再比较A[i]与B[j]，以此类推
        //构建长度为k的数组，放入matrix[0, 0]
        //比较行的下一个matrix[0][1]与列的下一个matrix[1][0]
        //将matrix[0][1]5放入数组
        //再比较matrix[0][1]行下一位matrix[0][2]和列下一位matrix[1][1]
        //以此内推，但这样也有问题，比如matrix[0][2]11行下一位和列下一位都是13，跳过了12
        //这个思路行不通
        //看答案后发现也有归并的做法
        //但归并只用到了行有序没有用到列有序的条件
        //答案中有使用小顶堆来实现排序后输出结果，感觉没啥意思...
        public int kthSmallest(int[][] matrix, int k) {
            int[] left = matrix[0];
            for (int i = 1; i < matrix.length; i++) {
                int[] right = matrix[i];
                left = merge(left, right);
            }
            return left[k - 1];
        }

        int[] merge(int[] left, int[] right) {
            int[] res = new int[left.length + right.length];
            int i = 0;
            int j = 0;
            int t = 0;
            while (i < left.length && j < right.length) {
                if(left[i] > right[j]) {
                    res[t++] = right[j++];
                }else{
                    res[t++] = left[i++];
                }
            }
            while (i < left.length) {
                res[t++] = left[i++];
            }
            while (j < right.length) {
                res[t++] = right[j++];
            }
            return res;
        }
    }

    class Solution3 {
        //二分查找法
        //对于从左到右，从上到下都有序的矩阵来说
        //寻找第k小的元素，可以想成通过二分法得到N个比较小的元素，并且令N逼近k
        //设置mid为左上最小值+（右下最大值-左下最小值）/2
        //从左下到右上确定比mid小的元素的个数
        public int kthSmallest(int[][] matrix, int k) {
            int left = matrix[0][0]; //最小值
            int right = matrix[matrix.length - 1][matrix.length - 1]; //最大值
            while (left < right) {
                int mid = left + ((right - left) / 2);
                if (check(matrix, mid, k)) {
                    right = mid;
                } else {
                    left = mid + 1;
                }
            }
            return left;
        }

        private boolean check(int[][] matrix, int mid, int k) {
            int i = matrix.length - 1; //左下
            int j = 0;
            int num = 0; //小于mid的个数
            while (i >= 0 && j <= matrix.length - 1) {
                if(matrix[i][j] <= mid) {
                    j++;
                    num += i + 1; //上面的元素都比mid小
                } else {
                    i--; //向上一行
                }
            }
            return num >= k; //需要缩小二分范围
        }
    }
}

package algorithms.leetcode._2022_convert_1d_array_into_2d_array;

import java.util.Arrays;

public class _2022 {
    // 给你一个下标从 0 开始的一维整数数组 original 和两个整数 m 和  n 。
    // 你需要使用 original 中 所有 元素创建一个 m 行 n 列的二维数组。

    // original 中下标从 0 到 n - 1 （都 包含 ）的元素构成二维数组的第一行，
    // 下标从 n 到 2 * n - 1 （都 包含 ）的元素构成二维数组的第二行，依此类推。

    // 请你根据上述过程返回一个 m x n 的二维数组。如果无法构成这样的二维数组，
    // 请你返回一个空的二维数组。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/convert-1d-array-into-2d-array
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public int[][] construct2DArray(int[] original, int m, int n) {
            if (m * n != original.length) {
                return new int[0][0];
            }
            int[][] res = new int[m][];
            for (int i = 0; i < m; i++) {
                res[i] = new int[n];
                for (int j = 0; j < n; j++) {
                    res[i][j] = original[i * n + j];
                }
            }
            return res;
        }
    }

    public static void main(String[] args) {
        _2022 _2022 = new _2022();
        Solution solution = _2022.new Solution();
        System.out.println(Arrays.toString(solution.construct2DArray(new int[]{1,1,1,1}, 4, 0)));
    }
}

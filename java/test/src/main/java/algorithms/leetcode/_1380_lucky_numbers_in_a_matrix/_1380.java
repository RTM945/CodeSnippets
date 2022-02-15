package algorithms.leetcode._1380_lucky_numbers_in_a_matrix;

import java.util.ArrayList;
import java.util.List;

public class _1380 {
    // 给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。
    // 请你按 任意 顺序返回矩阵中的所有幸运数。

    // 幸运数是指矩阵中满足同时下列两个条件的元素：

    // 在同一行的所有元素中最小
    // 在同一列的所有元素中最大

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public List<Integer> luckyNumbers (int[][] matrix) {
            List<Integer> res = new ArrayList<>();
            int m = matrix.length;
            int n = matrix[0].length;
            
            for (int i = 0; i < m; i++) {
                int row = 0;
                int rowmin = Integer.MAX_VALUE;
                // 求每一行的最小值
                for (int j = 0; j < n; j++) {
                    if (rowmin > matrix[i][j]) {
                        rowmin = matrix[i][j];
                        row = j;
                    }
                }
                boolean find = true;
                // 这一行的最小值也是这里列的最大值吗
                for (int j = 0; j < m; j++) {
                    if (matrix[j][row] > rowmin) {
                        find = false;
                        break;
                    }   
                }
                if (find) {
                    res.add(rowmin);
                }
            }
            return res;
        }
        
    }

    public static void main(String[] args) {
        _1380 _1380 = new _1380();
        Solution solution = _1380.new Solution();
        System.out.println(solution.luckyNumbers(new int[][]{{3,7,8},{9,11,13},{15,16,17}}));
    }
}

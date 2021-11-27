package algorithms.leetcode._519_random_flip_matrix;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Random;
import java.util.Set;

public class _519 {
    // 给你一个 m x n 的二元矩阵 matrix ，且所有值被初始化为 0 。
    // 请你设计一个算法，随机选取一个满足 matrix[i][j] == 0 的下标 (i, j) ，
    // 并将它的值变为 1 。所有满足 matrix[i][j] == 0 的下标 (i, j) 被选取的概率应当均等。
    // 尽量最少调用内置的随机函数，并且优化时间和空间复杂度。
    // 实现 Solution 类：
    // Solution(int m, int n) 使用二元矩阵的大小 m 和 n 初始化该对象
    // int[] flip() 返回一个满足 matrix[i][j] == 0 的随机下标 [i, j] ，
    // 并将其对应格子中的值变为 1
    // void reset() 将矩阵中所有的值重置为 0
    //  
    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/random-flip-matrix
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 纯暴力
    class Solution {
        Set<String> set;
        int r;
        int c;
        Random rand;

        public Solution(int n_rows, int n_cols) {
            set = new HashSet<>();
            r = n_rows;
            c = n_cols;
            rand = new Random();
        }

        public int[] flip() {
            int rr = rand.nextInt(r);
            int cc = rand.nextInt(c);
            while (set.contains(rr + "," + cc)) {
                rr = rand.nextInt(r);
                cc = rand.nextInt(c);
            }
            set.add(rr + "," + cc);
            return new int[] { rr, cc };
        }

        public void reset() {
            set.clear();
        }
    }

    public static void main(String[] args) {
        _519 _519 = new _519();
        Solution solution = _519.new Solution(3, 1);
        System.out.println(Arrays.toString(solution.flip()));
        solution.reset();
        System.out.println(Arrays.toString(solution.flip()));
        solution.reset();
        System.out.println(Arrays.toString(solution.flip()));
        solution.reset();
        System.out.println(Arrays.toString(solution.flip()));
        solution.reset();
    }

}

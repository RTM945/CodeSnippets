package algorithms.leetcode._688_knight_probability_in_chessboard;

public class _688 {
    // 在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。
    // 行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。

    // 象棋骑士有8种可能的走法，如下图所示。
    // 每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
    // 每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。

    // 骑士继续移动，直到它走了 k 步或离开了棋盘。
    
    // 返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
    
    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/knight-probability-in-chessboard
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 动态规划
    class Solution {

        final int[][] dirs = {{1, 2}, {2, 1}, {-1, 2}, {2, -1}, {1, -2}, {-2, 1}, {-1, -2}, {-2, -1}};

        public double knightProbability(int n, int k, int row, int column) {
            // dp[step][i][j] 表示骑士从i j 出发，走了 step 步后还在棋盘的概率
            // 当 i j 不在棋盘上时 dp[step][i][j] = 0
            // 当 i j 在棋盘上且 step = 0 时 dp[step][i][j] = 0
            // 其他情况下 dp[step][i][j] = 1 / 8 * ∑dp[step - 1][i + di][j + dj]
            double[][][] dp = new double[k + 1][n][n];
            for (int step = 0; step <= k; step++) {
                for (int i = 0; i < n; i++) {
                    for (int j = 0; j < n; j++) {
                        if (step == 0) {
                            dp[step][i][j] = 1;
                        } else {
                            for (int[] dir : dirs) {
                                int ni = i + dir[0];
                                int nj = j + dir[1];
                                if (ni >= 0 && ni < n && nj >= 0 && nj < n) {
                                    dp[step][i][j] += dp[step - 1][ni][nj] / 8;
                                }
                            }
                        }
                    }
                }
            }
            return dp[k][row][column];
        }
    }

    // 记忆化搜索
    class Solution1 {

        final int[][] dirs = {{1, 2}, {2, 1}, {-1, 2}, {2, -1}, {1, -2}, {-2, 1}, {-1, -2}, {-2, -1}};

        public double knightProbability(int n, int k, int row, int column) {
            double[][][] memo = new double[n][n][k + 1];
            return dfs(n, k, row, column, memo);
        }

        public double dfs(int n, int k, int i, int j, double[][][] memo) {
            if (i < 0 || j < 0 || i >= n || j >= n) {
                return 0;
            }
            if (k == 0) {
                return 1;
            }
            if (memo[i][j][k] != 0) {
                return memo[i][j][k];
            }
            double ans = 0;
            for (int[] dir : dirs) {
                ans += dfs(n, k - 1, i + dir[0], j + dir[1], memo) / 8;
            }

            memo[i][j][k] = ans;

            return ans;
        }
    }
}

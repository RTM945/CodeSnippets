package algorithms.leetcode._130_surrounded_regions;

// 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 示例:
// X X X X
// X O O X
// X X O X
// X O X X
// 运行你的函数后，矩阵变为：
// X X X X
// X X X X
// X X X X
// X O X X
// 解释:
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 
// 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
// 如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/surrounded-regions
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
public class _130 {
    class Solution {
        // 讨巧做法？
        // 不管他是不是环绕，把非边界的O全部换成X可不可以？
        // 看上去不行，如果一片O围着个X围着的O，就翻车了
        // 暴力的方法是，读到一个非边界的O，如果前后左右遍历到头至少都有一个X，则能判定它被环绕
        // 但这样可能有个问题，一个一个判断并改成X的会不会对其他O的判断有影响？
        // 错了，错在有前后左右遍历到底都有X，但X不相连的情况
        // 而且也没考虑有多个环绕区的情况
        public void solve(char[][] board) {
            int m = board.length;
            if (m < 1) {
                return;
            }
            int n = board[0].length;
            if (n < 1) {
                return;
            }
            for (int i = 1; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    if (check(board, i, j, m, n)) {
                        board[i][j] = 'X';
                    }
                }
            }
        }

        boolean check(char[][] board, int i, int j, int m, int n) {
            char x = board[i][j];
            if (x == 'O') {
                boolean front, back, left, right;
                front = back = left = right = false;
                // 前
                for (int f = 0; f < i; f++) {
                    if (board[f][j] == 'X') {
                        front = true;
                        break;
                    }
                }
                // 后
                for (int b = i + 1; b < m; b++) {
                    if (board[b][j] == 'X') {
                        back = true;
                        break;
                    }
                }
                // 左
                for (int l = 0; l < j; l++) {
                    if (board[i][l] == 'X') {
                        left = true;
                        break;
                    }
                }
                // 右
                for (int r = j + 1; r < n; r++) {
                    if (board[i][r] == 'X') {
                        right = true;
                        break;
                    }
                }
                return front && back && left && right;
            }
            return false;
        }
    }

    class Solution1 {
        // 重新考虑
        // 如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
        // 不行，不会
        // 看题解
        // 很难判断哪些 O 是被包围的，哪些 O 不是被包围的
        // 但任何边界上的 O 都不会被填充为 X
        // 也就是说，所有被包围的 O 都和边界上的 O 相连
        // 找到，然后把剩余的 O 变成 X
        // 这里有个技巧
        // 可以先把所有和边界 O 相连的 O 改为 A，然后把所有 O 改为 X，再把所有 A 改为 O
        int m, n;
        public void solve(char[][] board) {
            m = board.length;
            if (m == 0) {
                return;
            }
            n = board[0].length;
            for (int i = 0; i < m; i++) {
                dfs(board, i, 0); //第一列边界
                dfs(board, i, n - 1); //最后一列边界
            }
            for (int i = 0; i < n; i++) {
                dfs(board, 0, i); //第一行边界
                dfs(board, m - 1, i); //最后一行边界
            }
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    if(board[i][j] == 'A') {
                        board[i][j] = 'O';
                    } else if(board[i][j] == 'O') {
                        board[i][j] = 'X';
                    }
                }
            }
        }

        void dfs(char[][] board, int i, int j) {
            if (i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O') {
                return;
            }
            board[i][j] = 'A';
            dfs(board, i - 1, j);//前
            dfs(board, i + 1, j);//后
            dfs(board, i, j - 1);//左
            dfs(board, i, j + 1);//右
        }
    }

}
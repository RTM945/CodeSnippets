package algorithms.leetcode._794_valid_tic_tac_toe_state;

public class _794 {
    // 给你一个字符串数组 board 表示井字游戏的棋盘。
    // 当且仅当在井字游戏过程中，棋盘有可能达到 board 所显示的状态时，才返回 true 。
    // 井字游戏的棋盘是一个 3 x 3 数组，由字符 ' '，'X' 和 'O' 组成。字符 ' ' 代表一个空位。
    // 以下是井字游戏的规则：
    // 玩家轮流将字符放入空位（' '）中。
    // 玩家 1 总是放字符 'X' ，而玩家 2 总是放字符 'O' 。
    // 'X' 和 'O' 只允许放置在空位中，不允许对已放有字符的位置进行填充。
    // 当有 3 个相同（且非空）的字符填充任何行、列或对角线时，游戏结束。
    // 当所有位置非空时，也算为游戏结束。
    // 如果游戏结束，玩家不允许再放置字符。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/valid-tic-tac-toe-state
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 看上去是判断一个字符串按3x3分割后是否有连续的X或O
    class Solution {
        public boolean validTicTacToe(String[] board) {
            int xCount = 0, oCount = 0;
            for (String row : board) {
                for (char c : row.toCharArray()) {
                    xCount = (c == 'X') ? (xCount + 1) : xCount;
                    oCount = (c == 'O') ? (oCount + 1) : oCount;
                }
            }
            if (oCount != xCount && oCount != xCount - 1) {
                return false;
            }
            if (win(board, 'X') && oCount != xCount - 1) {
                return false;
            }
            if (win(board, 'O') && oCount != xCount) {
                return false;
            }
            return true;
        }
        

        public boolean win(String[] board, char p) {
            for (int i = 0; i < 3; ++i) {
                if (p == board[0].charAt(i) && p == board[1].charAt(i) && p == board[2].charAt(i)) {
                    return true;
                }
                if (p == board[i].charAt(0) && p == board[i].charAt(1) && p == board[i].charAt(2)) {
                    return true;
                }
            }
            if (p == board[0].charAt(0) && p == board[1].charAt(1) && p == board[2].charAt(2)) {
                return true;
            }
            if (p == board[0].charAt(2) && p == board[1].charAt(1) && p == board[2].charAt(0)) {
                return true;
            }
            return false;
        }
    }

    public static void main(String[] args) {
        _794 _794 = new _794();
        Solution solution = _794.new Solution();
        System.out.println(solution.validTicTacToe(new String[]{"XOX"," X ","   "})); //false
        System.out.println(solution.validTicTacToe(new String[]{"XXX","   ","OOO"})); //false
        System.out.println(solution.validTicTacToe(new String[]{"XO ","XO ","X  "})); //true
        System.out.println(solution.validTicTacToe(new String[]{"XOX","O O","XOX"})); //true
        System.out.println(solution.validTicTacToe(new String[]{"XXX","OOX","OOX"})); //true
        System.out.println(solution.validTicTacToe(new String[]{"XXX","XOO","OO "})); //false
    }
}
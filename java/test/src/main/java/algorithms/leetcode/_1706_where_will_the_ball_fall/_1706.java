package algorithms.leetcode._1706_where_will_the_ball_fall;

public class _1706 {
    class Solution {
        public int[] findBall(int[][] grid) {
            int[] ans = new int[grid[0].length];
            // 球从某一列落下
            for (int i = 0; i < grid[0].length; i++) {
                // 球所在行 从上往下 从 1 到 grid.length
                int row = 0; 
                // 球所在列
                int col = i;
                boolean pass = true;
                while (row < grid.length) {
                    // 向左 + 1 向右 - 1
                    int nc = col + grid[row][col];
                    if (nc < 0 || nc >= grid[0].length) {
                        // 到边缘了
                        ans[i] = -1;
                        pass = false;
                        break;
                    }
                    if (grid[row][col] != grid[row][nc]) {
                        // 必须同为 1 或者同为 -1 才不会有形成 V
                        ans[i] = -1;
                        pass = false;
                        break;
                    }
                    row++; // 到下一行
                    col = nc; // 到下一列
                }
                if (pass) {
                    ans[i] = col;
                }
                
            }
            return ans;
        }
    }
}

package algorithms.leetcode._063_unique_paths_ii;

import java.util.ArrayList;
import java.util.List;

//https://leetcode-cn.com/problems/unique-paths-ii/
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
public class _063 {
    class Solution {
        //很像coursera普林斯顿大学公开课数据结构与算法第一节课的并查集问题
        //然而这个题目要求输出多少条路，而不是连通性本身
        //有一个想法，利用分支构建树，如果树的叶子结点为右下的终点，则为一条路
        //每个结点只能向下或向右
        //试试看
        //https://leetcode-cn.com/submissions/detail/85354386/testcase/ 过不去，内存炸了
        public int uniquePathsWithObstacles(int[][] obstacleGrid) {
            //一行的情况
            if (obstacleGrid.length == 1) {
                for (int i = 0; i < obstacleGrid[0].length; i++) {
                    if (obstacleGrid[0][i] == 1) {
                        return 0;
                    }
                }
                return 1;
            }

            //一列的情况
            if (obstacleGrid[0].length == 1) {
                for (int i = 0; i < obstacleGrid.length; i++) {
                    if (obstacleGrid[i][0] == 1) {
                        return 0;
                    }
                }
                return 1;
            }
            if (obstacleGrid[0][0] == 0) {
                return 0;
            }
            Node node = new Node(0, 0);
            node.down = walk(1, 0, obstacleGrid); //向下
            node.right = walk(0, 1, obstacleGrid); //向右
            //遍历树找终点
            //找出所有根到叶子节点的路径中叶子结点为终点的
            List<Node[]> allPath = new ArrayList<>();
            path(node, new ArrayList<>(), allPath, obstacleGrid.length - 1, obstacleGrid[0].length - 1);
            return allPath.size();
        }

        void path(Node node, List<Node> path, List<Node[]> allPath, int i, int j) {
            if (node == null) {
                return;
            }
            path.add(node);
            if (node.right == null && node.down == null && node.i == i && node.j == j) {
                allPath.add(path.toArray(new Node[0]));
                path.clear();
                //这是个错误的方法，path clear了就没有前面的路径了
            } else {
                path(node.down, path, allPath, i, j);
                path(node.right, path, allPath, i, j);
            }
        }

        Node walk(int i, int j, int[][] obstacleGrid) {
            if (obstacleGrid[i][j] == 1) {
                //障碍物
                return null;
            }
            Node node = new Node(i, j);
            if (i == obstacleGrid.length - 1) {
                //到纵边界了
                node.down = null;
            } else {
                node.down = walk(i + 1, j, obstacleGrid); //向下
            }
            if (j == obstacleGrid[0].length - 1) {
                //到横边界了
                node.right = null;
            } else {
                node.right = walk(i, j + 1, obstacleGrid); //向右
            }
            return node;
        }
    }

    class Node {
        int i, j;//下标
        Node right;
        Node down;

        public Node(int i, int j) {
            this.i = i;
            this.j = j;
        }
    }

    class Solution1{
        //上面的实现过于繁琐，还需要二叉树辅助
        //可以直接使用数组做dfs遍历
        //https://leetcode-cn.com/submissions/detail/85356468/testcase/
        //但还是会超时
        int count = 0;//全局变量,每走到终点就+1
        public int uniquePathsWithObstacles(int[][] arr) {
            dfs(arr,0,0);
            return count;
        }

        private void dfs(int[][] arr, int x, int y) {
            if (x >= arr.length || y >= arr[0].length || arr[x][y] == 1) {
                //越界 或者有障碍
                return;
            }
            if (x == arr.length - 1 && y == arr[0].length - 1) {
                //到达终点
                count++;
                return;
            }
            //先向右走
            dfs(arr, x, y + 1);
            //向下走
            dfs(arr, x + 1, y);
        }
    }

    class Solution2 {
        //只能用dp啦
        //一般来说dp用来做最优解问题，不知道为啥也能用在这里
        //看上去能拆分成子问题，并且需要避免子问题重复计算的都可以用dp
        //dp数组就是来存子问题结果避免重复计算的
        //动态规划的题目分为两大类，一种是求最优解类，典型问题是背包问题
        //另一种就是计数类，比如这里的统计方案数的问题，它们都存在一定的递推性质。
        //前者的递推性质还有一个名字，叫做 「最优子结构」即当前问题的最优解取决于子问题的最优解
        //后者类似，当前问题的方案数取决于子问题的方案数
        //所以在遇到求方案数的问题时，我们可以往动态规划的方向考虑。

        //grid[i][j]必是从grid[i-1][j]或grid[i][j-1]走过来的
        //grid[i][j]的路径数量则为grid[i-1][j]的路径数量+grid[i][j-1]的路径数量
        //而grid[i-1][j]或grid[i][j-1] = 1时，就等于毙掉了所有通向它的路径，它的dp值为0
        //以此构造dp
        //状态定义dp[i][j]表示走到grid[i][j]的方法数
        //状态转移 i,j无阻挡时dp[i][j] = dp[i-1][j] + dp[i][j-1]  i,j有阻挡dp[i][j] = 0
        //初始条件
        //只有一行时，只能往右走，所以dp[i,0]为1，如果有障碍，则为0
        //只有一列时，只能往下走，所以dp[0,j]为1，如果有障碍，则为0
        public int uniquePathsWithObstacles(int[][] obstacleGrid) {
            if (obstacleGrid == null || obstacleGrid.length == 0) {
                return 0;
            }
            int[][] dp = new int[obstacleGrid.length][obstacleGrid[0].length];
            for (int i = 0; i < obstacleGrid.length && obstacleGrid[i][0] == 0; i++) {
                dp[i][0] = 1;
            }
            for (int j = 0; j < obstacleGrid[0].length && obstacleGrid[0][j] == 0; j++) {
                dp[0][j] = 1;
            }
            for (int i = 1; i < obstacleGrid.length; i++) {
                for (int j = 1; j < obstacleGrid[0].length; j++) {
                    if (obstacleGrid[i][j] == 0) {
                        dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                    }
                }
            }
            return dp[obstacleGrid.length - 1][obstacleGrid[0].length - 1];
        }
    }

    public static void main(String[] args) {
        int[][] map = {
                {0, 0, 0},
                {0, 1, 0},
                {0, 0, 0}
        };
        _063 q = new _063();
        System.out.println(q.new Solution().uniquePathsWithObstacles(map));
    }
}

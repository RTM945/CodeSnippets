package algorithms.leetcode._785_is_graph_bipartite;

import java.util.LinkedList;
import java.util.Queue;

//https://leetcode-cn.com/problems/is-graph-bipartite/
//给定一个无向图graph，当这个图为二分图时返回true。
//如果我们能将一个图的节点集合分割成两个独立的子集A和B
//并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。
//graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。
//每个节点都是一个在0到graph.length-1之间的整数。
//这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。
//示例 1:
//输入: [[1,3], [0,2], [1,3], [0,2]]
//输出: true
//解释:
//无向图如下:
//0----1
//|    |
//|    |
//3----2
//我们可以将节点分成两组: {0, 2} 和 {1, 3}。
//示例 2:
//输入: [[1,2,3], [0,2], [0,1,3], [0,2]]
//输出: false
//解释:
//无向图如下:
//0----1
//| \  |
//|  \ |
//3----2
//我们不能将节点分割成两个独立的子集。
//注意:
//graph 的长度范围为 [1, 100]。
//graph[i] 中的元素的范围为 [0, graph.length - 1]。
//graph[i] 不会包含 i 或者有重复的值。
//图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。
public class _785 {

    class Solution {
        //描述好多，头有点大
        //复习了一下邻接表，就和描述中说的一样
        //graph[i]表示图中与节点i相连的所有节点
        //注意graph[i][j]的值也是下标
        //二分图是值每一条边的两个节点一个来自A集合，一个来自B集合
        //画图了解到，结果子集中的任意两个元素，都不能有边相连
        //那么有个很简单的想法是
        //列举出邻接表每一行中不存在的图中结点作为分组
        //如果能成为二分图，分组的元素要么完全相同，要么完全不同
        //实际做起来虽然简单，但是繁琐...
        //官方提到着色
        //将某一结点染成红色，从他开始遍历，深度广度皆可，与该点直接相连的结点染成绿色
        //也就是邻接表中的元素染成绿色
        //继续遍历染色，如果访问到一节点已经染色且颜色与前驱颜色不同，则说明不是二分图
        //如果全都染成了红色和绿色，则红绿分开则是分组

        private int[] color;
        private int uncolor = 0;
        private int red = 1;
        private int green = 2;

        public boolean isBipartite(int[][] graph) {
            if(graph.length == 1){
                return false;
            }
            int r = graph.length;
            color = new int[r];
            for (int i = 0; i < r; i++) {
                if(color[i] == uncolor) {
                    color[i] = red;
                    Queue<Integer> queue = new LinkedList<>();
                    queue.offer(i);
                    while (!queue.isEmpty()) {
                        int node = queue.poll();
                        int cn = color[node] == red ? green : red;
                        for (int neighbor : graph[node]) {
                            if(color[neighbor] == uncolor) {
                                queue.offer(neighbor);
                                color[neighbor] = cn;
                            }else if(color[neighbor] != cn) {
                                return false;
                            }
                        }
                    }
                }
            }
            return true;
        }
    }

    public static void main(String[] args) {
        _785 q = new _785();
        int[][] graph = {
                {1,3},
                {0,2},
                {1,3},
                {0,2}
        };
        System.out.println(q.new Solution().isBipartite(graph));
    }
}

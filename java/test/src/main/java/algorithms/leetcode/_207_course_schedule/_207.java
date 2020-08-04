package algorithms.leetcode._207_course_schedule;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
// 在选修某些课程之前需要一些先修课程。 
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
// 示例 1:
// 输入: 2, [[1,0]] 
// 输出: true
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
// 示例 2:
// 输入: 2, [[1,0],[0,1]]
// 输出: false
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；
// 并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
// 提示：
// 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
// 1 <= numCourses <= 10^5
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
public class _207 {
    class Solution {
        //听上去是拓扑排序
        //不能构成环
        //得到每个课程的入度
        //将入度为0的放入队列
        //出队时，将其邻接点入度-1，如果是0，再入队
        //成功出队，numCourses-1，如果能减到0，排课就成功了
        //如何得出入度？题目说边缘列表https://zhuanlan.zhihu.com/p/110022390
        //[1,0]表示1->0也就是1依赖0
        //每一行前面的元素指向后一个元素
        //可以用一个集合来维护指向每个课程的所有课程,它的size就是该课程的入度
        public boolean canFinish(int numCourses, int[][] prerequisites) {
            int[] in = new int[numCourses];
            Queue<Integer> queue = new LinkedList<>();
            List<List<Integer>> adjacency = new ArrayList<>();
            for (int i = 0; i < numCourses; i++) {
                adjacency.add(new ArrayList<>()); //指向它的课程列表
            }
            for(int[] cp : prerequisites) {
                in[cp[0]]++;//cp[0]的前驱个数
                adjacency.get(cp[1]).add(cp[0]);//多了一个邻接点
            }
            for (int i = 0; i < numCourses; i++) {
                if(in[i] == 0) {
                    queue.add(i);
                }
            }
            while(!queue.isEmpty()) {
                int i = queue.poll();
                numCourses--;//排了一门课
                for (int k : adjacency.get(i)) { //指向它的邻接点
                    in[k]--; //去了一个前驱
                    if(in[k] == 0) {
                        queue.add(k); //它可以成为前驱
                    }
                }
            }
            return numCourses == 0;
        }
    }
}
package algorithms.leetcode._630_course_schedule_iii;

import java.util.Arrays;
import java.util.PriorityQueue;

public class _630 {
    // 有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，
    // 其中 courses[i] = [durationi, lastDayi] 表示第 i 门课将会 持续 上 durationi 天课，
    // 并且必须在不晚于 lastDayi 的时候完成。

    // 你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。

    // 返回你最多可以修读的课程数目。

    // 输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
    // 输出：3
    // 解释：
    // 这里一共有 4 门课程，但是你最多可以修 3 门：
    // 首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
    // 第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
    // 第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
    // 第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/course-schedule-iii
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public int scheduleCourse(int[][] courses) {
            // 以结束时间排序
            // 结束时间越早越需要提前完成
            Arrays.sort(courses, (c1, c2) -> c1[1] - c2[1]);
            // 储存已选择的课程，按照持续时间排序
            PriorityQueue<int[]> heap = new PriorityQueue<>((c1, c2) -> c2[0] - c1[0]);
            int day = 0;
            for (int[] c : courses) {
                if (day + c[0] <= c[1]) {
                    // 如果当前课程时间不冲突，将该课程加入队列
                    // 这里的不冲突可以理解为，0~day+c[0]这段区间，我们还可以再插入当前一节课
                    day += c[0];
                    heap.offer(c);
                } else if (!heap.isEmpty() && heap.peek()[0] > c[0]) {
                    // 课程时间冲突，且有选过其他课，这时我们找到最长时间的课程，用当前的短课替换了，余出了更多的空区间
                    // 所以这里我们余出的时间其实就是两者的持续时间之差，课程变短了，day会前移，这样我们相当于变相给后面的课程增加了选择的区间
                    day -= heap.poll()[0] - c[0];
                    heap.offer(c);
                }
            }
            return heap.size();
        }
    }
}

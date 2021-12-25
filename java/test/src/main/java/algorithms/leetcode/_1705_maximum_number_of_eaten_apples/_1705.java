package algorithms.leetcode._1705_maximum_number_of_eaten_apples;

import java.util.PriorityQueue;

public class _1705 {
    // 有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。
    // 在第 i 天，树上会长出 apples[i] 个苹果，
    // 这些苹果将会在 days[i] 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。
    // 也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] == 0 表示。

    // 你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。

    // 给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/maximum-number-of-eaten-apples
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 将苹果按其腐烂时间排序
    // 用个数组，0位表示过期的时间，1位表示过期的数量
    class Solution {
        public int eatenApples(int[] apples, int[] days) {
            PriorityQueue<int[]> queue = new PriorityQueue<>((a, b) -> a[0] - b[0]);
            int res = 0;
            int n = apples.length; // 总天数
            int i = 0;
            while (i < n) {
                while (!queue.isEmpty() && queue.peek()[0] <= i) {
                    // 烂苹果扔了
                    queue.poll();
                }
                int rottenDay = i + days[i]; // 腐烂时间
                int count = apples[i];
                if (count > 0) {
                    queue.offer(new int[]{rottenDay, count});
                }
                if (!queue.isEmpty()) {
                    int[] arr = queue.peek();
                    arr[1]--;
                    if (arr[1] == 0) {
                        queue.poll();
                    }
                    res++;
                }
                i++;
            }
            // 存下来的苹果还能吃多少天
            while (!queue.isEmpty()) {
                while (!queue.isEmpty() && queue.peek()[0] <= i) {
                    // 烂苹果扔了
                    queue.poll();
                }
                if (queue.isEmpty()) {
                    break;
                }
                int[] arr = queue.poll();
                // 这批苹果还能吃多少天
                int curr = Math.min(arr[0] - i, arr[1]);
                res += curr;
                i += curr;
            }

            return res;
        }
    }
}

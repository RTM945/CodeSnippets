package algorithms.leetcode._786_k_th_smallest_prime_fraction;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.PriorityQueue;

public class _786 {
    // 给你一个按递增顺序排序的数组 arr 和一个整数 k 。
    // 数组 arr 由 1 和若干 素数  组成，且其中所有整数互不相同。

    // 对于每对满足 0 < i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。

    // 那么第 k 个最小的分数是多少呢?  以长度为 2 的整数数组返回你的答案,
    // 这里 answer[0] == arr[i] 且 answer[1] == arr[j] 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/k-th-smallest-prime-fraction
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {

        // 自排序集合
        // 为什么叫多路归并呢
        public int[] kthSmallestPrimeFraction(int[] arr, int k) {
            int n = arr.length;
            PriorityQueue<int[]> pq = new PriorityQueue<int[]>((x, y) -> arr[x[0]] * arr[y[1]] - arr[y[0]] * arr[x[1]]);
            // 把arr第一个元素和每个元素作分母的分数得出来，队首就是第一小的元素
            for (int i = 1; i < n; i++) {
                pq.offer(new int[] { 0, i });
            }
            // 从第一小的元素出发，他的同分母后一位元素入队
            // 自排序
            for (int i = 1; i < k; i++) {
                int[] frac = pq.poll();
                int x = frac[0], y = frac[1];
                if (x + 1 < y) {
                    pq.offer(new int[] { x + 1, y });
                }
            }
            int[] min = pq.peek();
            return new int[] { arr[min[0]], arr[min[1]] };
        }
    }

    class Solution1 {
        // 纯暴力
        public int[] kthSmallestPrimeFraction(int[] arr, int k) {
            int n = arr.length;
            List<int[]> frac = new ArrayList<int[]>();
            for (int i = 0; i < n; ++i) {
                for (int j = i + 1; j < n; ++j) {
                    frac.add(new int[]{arr[i], arr[j]});
                }
            }
            Collections.sort(frac, (x, y) -> x[0] * y[1] - y[0] * x[1]);
            return frac.get(k - 1);
        }
    }

    public static void main(String[] args) {
        _786 _786 = new _786();
        Solution solution = _786.new Solution();
        solution.kthSmallestPrimeFraction(new int[] { 1,2,3,5 }, 3);
    }
}

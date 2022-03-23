package algorithms.leetcode._440_k_th_smallest_in_lexicographical_order;

public class _440 {
    // 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
    //     输入: n = 13, k = 2
    // 输出: 10
    // 解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 十叉树的先序遍历，但全部构建会超时
    class Solution {
        public int findKthNumber(int n, int k) {
            // 从1开始
            int curr = 1;
            // 跳过根节点1
            k--;
            while (k > 0) {
                int steps = getSteps(curr, n);
                // 子节点个数不大于k的时候，说明第k小的不在以curr为根节点的子树中
                if (steps <= k) {
                    // 跳过这么多节点
                    k -= steps;
                    // 去当前节点右边的节点
                    curr++;
                } else {
                    // 先序遍历 最左边的子节点
                    curr = curr * 10;
                    // 跳过根
                    k--;
                }
            }
            return curr;
        }
    
        // 以curr为根的节点个数 
        public int getSteps(int curr, long n) {
            int steps = 0;
            long first = curr;
            long last = curr;
            // 根节点小于n的时候 一层一层往下加
            while (first <= n) {
                // 最左边为 first * 10
                // 最右边为 last * 10 + 9
                steps += Math.min(last, n) - first + 1;
                first = first * 10;
                last = last * 10 + 9;
            }
            return steps;
        }
    }
}

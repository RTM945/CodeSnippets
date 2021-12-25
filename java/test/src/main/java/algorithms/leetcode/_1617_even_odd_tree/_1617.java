package algorithms.leetcode._1617_even_odd_tree;

import java.util.LinkedList;
import java.util.Queue;

public class _1617 {
    // 如果一棵二叉树满足下述几个条件，则可以称为 奇偶树 ：

    // 二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
    // 偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
    // 奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
    // 给你二叉树的根节点，如果二叉树为 奇偶树 ，则返回 true ，否则返回 false 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/even-odd-tree
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 广度优先遍历

    class Node {
        public int level;
        public TreeNode node;

        public Node(int level, TreeNode node) {
            this.level = level;
            this.node = node;
        }
    }

    class Solution {
        public boolean isEvenOddTree(TreeNode root) {
            Queue<Node> queue = new LinkedList<>();
            int level = 0;
            queue.offer(new Node(level, root));
            int pre = 0;
            while (!queue.isEmpty()) {
                Node head = queue.poll();
                TreeNode t = head.node;
                if (head.level == level) {
                    // 同一层
                    if ((head.level & 1) == 0 && ((t.val & 1) == 0 || t.val <= pre)) {
                        // 偶数层，值必须是奇数且递增
                        return false;
                    } else if ((head.level & 1) == 1 && ((t.val & 1) == 1 || t.val >= pre)) {
                        // 奇数层，值必须是偶数且递减
                        return false;
                    }
                    pre = t.val;
                } else {
                    // 下一层
                    if ((head.level & 1) == 0 && (t.val & 1) == 0) {
                        // 偶数层
                        return false;
                    } else if ((head.level & 1) == 1 && (t.val & 1) == 1) {
                        // 奇数层
                        return false;
                    }
                    ++level;
                    pre = t.val;
                }
                if (null != t.left) {
                    queue.offer(new Node(head.level + 1, t.left));
                }
                    
                if (null != t.right) {
                    queue.offer(new Node(head.level + 1, t.right));
                }
            }
            return true;
        }
    }
}

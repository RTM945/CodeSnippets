package algorithms.leetcode._589_n_ary_tree_preorder_traversal;

import java.util.ArrayList;
import java.util.List;

public class _589 {
    // 给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。

    // n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public List<Integer> preorder(Node root) {
            List<Integer> res = new ArrayList<>();
            preorder(root, res);
            return res;
        }

        public void preorder(Node node, List<Integer> list) {
            if (node != null) {
                list.add(node.val);
                if (node.children != null) {
                    for (Node c : node.children) {
                        preorder(c, list);
                    }
                }
            }
        }
    }
}

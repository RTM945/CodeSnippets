package algorithms.leetcode._590_n_ary_tree_postorder_traversal;

import java.util.ArrayList;
import java.util.List;

public class _590 {
    //N 叉树的后序遍历
    class Solution {
        public List<Integer> postorder(Node root) {
            List<Integer> res = new ArrayList<>();
            postorder(root, res);
            return res;
        }

        public void postorder(Node node, List<Integer> list) {
            if (node != null) {
                if (node.children != null) {
                    for (Node c : node.children) {
                        postorder(c, list);
                    }
                }
                list.add(node.val);
            }
        }
    }
}

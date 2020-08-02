package algorithms.leetcode._114_flatten_binary_tree_to_linked_list;

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
// 给定一个二叉树，原地将它展开为一个单链表。
// 例如，给定二叉树
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// 将其展开为：
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
class Solution {
    //中等
    //如果不说原地还是很简单的，先序遍历回溯都OK
    //原地的话，看上去是将所有的左子树插入到右子树之前
    //递归
    public void flatten(TreeNode root) {
        if(root != null && (root.left != null || root.right != null)) {
            flatten(root.left);
            flatten(root.right);
            if(root.left != null) {
                TreeNode right = root.right;
                TreeNode left = root.left;
                root.right = left;
                root.left = null;
                //需要找到叶子结点
                while(left.right != null) {
                    left = left.right;
                }
                left.right = right;
            }
        }
    }
}
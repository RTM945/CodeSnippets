package algorithms.leetcode._094_binary_tree_inorder_traversal;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * Definition for a binary tree node.
 */
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
//二叉树中序遍历
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> vals = new ArrayList<>();
        traversal(root, vals);
        return vals;
    }

    public void traversal(TreeNode root, List<Integer> vals) {
        if (root == null) {
            return;
        }
        traversal(root.left, vals);
        vals.add(root.val);
        traversal(root.right, vals);
    }
}

// 非递归方法
class Solution2 {

    //借助其他数据结构
    //栈
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> vals = new ArrayList<>();
        LinkedList<TreeNode> stack = new LinkedList<>();
        TreeNode node = root;
        while (node != null || !stack.isEmpty()) {
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            node = stack.pop();
            vals.add(node.val);
            node = node.right;
        }
        return vals;
    }
}

//莫里斯遍历略
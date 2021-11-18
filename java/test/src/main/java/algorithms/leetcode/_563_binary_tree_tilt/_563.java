package algorithms.leetcode._563_binary_tree_tilt;


// https://leetcode-cn.com/problems/binary-tree-tilt
// 给定一个二叉树，计算 整个树 的坡度 。
// 一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。
// 如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
// 整个树 的坡度就是其所有节点的坡度之和。

public class _563 {
    // 后序遍历
    // 但需要计算坡度和左右子树val的和，所以一个后续遍历是不够的
    class Solution {
        public int findTilt(TreeNode root) {
            if(root == null) {
                return 0;
            }
            return Math.abs(sum(root.left) - sum(root.right)) + findTilt(root.left) + findTilt(root.right);
        }

        public int sum(TreeNode root){
            if(root == null) {
                return 0;
            }
            return root.val + sum(root.left) + sum(root.right);
        }
    }

    class Solution1 {
        int sum = 0;
    
        public int findTilt(TreeNode root) {
            if(root == null) {
                return 0;
            }
            sum(root);
            return sum;
        }

        private int sum(TreeNode root){
            if(root == null) {
                return 0;
            }
            int l = sum(root.left);
            int r = sum(root.right);
            sum += Math.abs(l - r);
            return l + r + root.val;
        }
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode();
        root.left = new TreeNode();
        root.left.val = 2;
        root.right = new TreeNode();
        root.right.val = 3;
        Solution solution = new _563().new Solution();
        int res = solution.findTilt(root);
        System.out.println(res);
    }
}

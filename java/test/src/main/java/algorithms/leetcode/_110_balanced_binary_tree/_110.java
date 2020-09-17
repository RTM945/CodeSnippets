package algorithms.leetcode._110_balanced_binary_tree;

/* 给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _110 {
    class Solution {
        // 看过一点平衡二叉树，觉得左旋右旋之类的还挺麻烦的
        // 这道题只判断，属于easy
        // 但真的不像..
        // 首先要知道获取深度的方法是递归1 + max(depth(root.left), depth(root.left))
        // 然后需要修改
        // 当平衡时，才能返回深度，不平衡时，返回-1，这样递归上去的和必然小于0
        public boolean isBalanced(TreeNode root) {
            return depth(root) >= 0;
        }

        int depth(TreeNode node) {
            if (node == null) {
                return 0;
            }
            int dl = depth(node.left);
            int dr = depth(node.right);
            if (dl == -1 || dr == -1 || Math.abs(dl - dr) > 1) {
                return -1;
            } else {
                return Math.max(dl, dr) + 1;
            }
        }
    }
}
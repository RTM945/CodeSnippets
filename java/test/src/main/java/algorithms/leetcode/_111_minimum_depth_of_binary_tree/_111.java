package algorithms.leetcode._111_minimum_depth_of_binary_tree;

/* 给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明: 叶子节点是指没有子节点的节点。
示例:
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _111 {
    class Solution {
        //感觉树相关的题目都能递归的样子
        //有个坑是[1, 2]要求答案是2
        //直接1 + Math.min(minDepth(root.left), minDepth(root.right))会得到1
        //这样就要求空子树的0不能参与最小值的比较
        public int minDepth(TreeNode root) {
            if(root == null) {
                return 0;
            }
            // null节点不参与比较
            if (root.left == null && root.right != null) {
                return 1 + minDepth(root.right);
            }
            if (root.right == null && root.left != null) {
                return 1 + minDepth(root.left);
            }
            return 1 + Math.min(minDepth(root.left), minDepth(root.right));
        }
    }
}
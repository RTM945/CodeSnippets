package algorithms.leetcode._108_convert_sorted_array_to_binary_search_tree;

//https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }

    //没什么花头的一道题..
    //如果只构建一个二叉搜索树，可以直接第一位做根
    //但要求有序数组，高度平衡
    //只需要从数组中间分成子数组，再持续的递归找子数组中位做父结点就可以了
    class Solution {
        public TreeNode sortedArrayToBST(int[] nums) {
            return sub(nums, 0, nums.length - 1);
        }

        public TreeNode sub(int[] nums, int start, int end) {
            if(start > end) {
                return null;
            }
            int mid = (start + end) / 2;
            TreeNode node = new TreeNode(nums[mid]);
            node.left = sub(nums, start, mid - 1);
            node.right = sub(nums, mid + 1, end);
            return node;
        }
    }
}


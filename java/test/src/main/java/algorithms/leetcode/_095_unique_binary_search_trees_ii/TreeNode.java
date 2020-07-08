package algorithms.leetcode._095_unique_binary_search_trees_ii;


import java.util.LinkedList;
import java.util.List;

// Definition for a binary tree node.
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

//给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树。
class Solution {
    //https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
    //乍一看是个数字和null的排列组合问题，但它的遍历输出一开始令人迷惑
    //比如是层序遍历，第一个例子应该是[1,null,3,null,null,2,null]
    //不过又一想，在遇到null时，不会再去找他的左右子节点，[1,null,3,2]的输出可能合理
    //再一看这是BST
    //对于BST，左子节点总是小于根结点，右子节点总是大于根结点
    //才明白它是对于不符合BST上述性质的结点，才会输出null
    //本题中，如果选定了根结点i，则左边的元素则为(0,i),右边的元素为(i,n]
    //然后根据此规则进行递归
    public List<TreeNode> generateTrees(int n) {
        List<TreeNode> list = new LinkedList<>();
        if(n == 0) {
            return list;
        }
        return generateTrees(1, n);
    }

    List<TreeNode> generateTrees(int start, int end){
        List<TreeNode> list = new LinkedList<>();
        if(start > end) {
            list.add(null);//不符合BST性质了
            return list;
        }
        for (int i = start; i <= end; i++) {
            List<TreeNode> left = generateTrees(start, i - 1);
            List<TreeNode> right = generateTrees(i + 1, end);
            //left和right所有可能的排列方式
            for(TreeNode l : left) {
                for(TreeNode r : right) {
                    //根为i
                    TreeNode parent = new TreeNode();
                    parent.val = i;
                    parent.left = l;
                    parent.right = r;
                    list.add(parent);
                }
            }
        }
        return list;
    }
}
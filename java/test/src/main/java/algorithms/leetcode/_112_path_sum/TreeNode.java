package algorithms.leetcode._112_path_sum;

import java.lang.reflect.Array;
import java.util.*;

//https://leetcode-cn.com/problems/path-sum/
//给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
//说明: 叶子节点是指没有子节点的节点。
//示例: 
//给定如下二叉树，以及目标和 sum = 22，
//5
/// \
//4   8
///   / \
//11  13  4
///  \      \
//7    2      1
//返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }

    public static void main(String[] args) {
        int[] arr = {0,1,2,3,4,5,6,7,8};
        TreeNode node = new TreeNode(arr[0]);
        node.left = node.buildTreeByLevel(arr, 1);
        node.right = node.buildTreeByLevel(arr, 2);
        List<Integer> path = new ArrayList<>();
        List<Integer[]> allPath = new ArrayList<>();
        node.dfs(node, path, allPath);
        for (Integer[] p : allPath) {
            for (int v : p) {
                System.out.print(v + " ");
            }
            System.out.println();
        }
    }

    TreeNode buildTreeByLevel(int[] arr) {
        if(arr == null || arr.length < 1) {
            return null;
        }
        return buildTreeByLevel(arr, 0);
    }

    TreeNode buildTreeByLevel(int[] arr, int index) {
        if(index > arr.length - 1) {
            return null;
        }
        TreeNode node = new TreeNode(arr[index]);
        node.left = buildTreeByLevel(arr, 2 * index + 1);
        node.right = buildTreeByLevel(arr, 2 * index + 2);
        return node;
    }

    void dfs(TreeNode node, List<Integer> path, List<Integer[]> allPath) {
        if(node == null) {
            return;
        }
        path.add(node.val);
        if(node.left == null && node.right == null) {
            allPath.add(path.toArray(new Integer[0]));
        }else{
            dfs(node.left, path, allPath);
            dfs(node.right, path, allPath);
        }
        path.remove(path.size() - 1); //往回退1
    }
}

class Solution {
    //看到简单的题很兴奋，以为能减少很多时间
    //昨天正好觉得自己知道了如何将树从根遍历到叶子并记录所有路径
    //用一个集合存路径上的结点，如果遍历到叶子结点，就保存集合副本，集合删除掉最后一个元素作为回退
    //所以思路就是遍历出所有的路径并且计算是否能得到指定结果
    //其实是很蠢的方法
    public boolean hasPathSum(TreeNode root, int sum) {
        List<Integer> path = new ArrayList<>();
        List<Integer[]> allPath = new ArrayList<>();
        dfs(root, path, allPath);
        for (Integer[] p : allPath) {
            int count = 0;
            for (int v : p) {
                count += v;
            }
            if(count == sum) {
                return true;
            }
        }
        return false;
    }

    void dfs(TreeNode node, List<Integer> path, List<Integer[]> allPath) {
        if(node == null) {
            return;
        }
        path.add(node.val);
        if(node.left == null && node.right == null) {
            allPath.add(path.toArray(new Integer[0]));
        }else{
            dfs(node.left, path, allPath);
            dfs(node.right, path, allPath);
        }
        path.remove(path.size() - 1); //往回退1
    }

    class Solution1 {
        //直接用sum从根结点依次减到子节点
        //使用递归的思路
        //sum减根结点的值应该等于根的左子树的hasPathSum结果或右子树的hasPathSum结果
        public boolean hasPathSum(TreeNode root, int sum) {
            if(root == null) {
                return false;
            }
            if(root.left == null && root.right == null) {
                return sum == root.val;
            }
            return hasPathSum(root.left, sum - root.val) || hasPathSum(root.right, sum - root.val);
        }
    }

    class Solution2{
        //另一种广度优先的做法
        //使用队列作为辅助，记录根结点到当前结点的和防止重复计算
        //即根入队，取出，找到左右结点，将和放入队列，这样循环直到匹配目标值
        //因为结点和数值不是一个类型，所以需要两个队列
        //当取出的结点没有子节点时，就要判断是否与目标匹配了
        //如果有子结点，就将新的和与子节点入队
        public boolean hasPathSum(TreeNode root, int sum) {
            if (root == null) {
                return false;
            }
            Queue<TreeNode> nodes = new LinkedList<>();
            Queue<Integer> vals = new LinkedList<>();
            nodes.offer(root);
            vals.offer(root.val);
            while (!nodes.isEmpty()) {
                TreeNode node = nodes.poll();
                int val = vals.poll();
                if(node.left == null && node.right == null) {
                    if(val == sum) {
                        return true;
                    }
                }else{
                    if(node.left != null) {
                        nodes.offer(node.left);
                        vals.offer(val + node.left.val);
                    }
                    if(node.right != null) {
                        nodes.offer(node.right);
                        vals.offer(val + node.right.val);
                    }
                }
            }
            return false;
        }
    }
}

package algorithms.leetcode._559_maximum_depth_of_n_ary_tree;

import java.util.List;

// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
// 给定一个 N 叉树，找到其最大深度。
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
// N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
public class _559 {
    
    // Definition for a Node.
    class Node {
        public int val;
        public List<Node> children;

        public Node() {}

        public Node(int _val) {
            val = _val;
        }

        public Node(int _val, List<Node> _children) {
            val = _val;
            children = _children;
        }
    }


    class Solution {
        public int maxDepth(Node root) {
            if (root == null) {
                return 0;
            }

            if (root.children != null) {
                int max = 0;
                for (Node node : root.children) {
                    max = Math.max(max, maxDepth(node));
                }
                return max + 1;
            } 
            
            return 1;
        }
    }
}

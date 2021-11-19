package algorithms.leetcode._397_integer_replacement;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// https://leetcode-cn.com/problems/integer-replacement
// 给定一个正整数 n ，你可以做如下操作：

// 如果 n 是偶数，则用 n / 2替换 n 。
// 如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
// n 变为 1 所需的最小替换次数是多少？

public class _397 {

    class TreeNode {
        int val;
        List<TreeNode> childs;

        TreeNode(int val) {
            this.val = val;
        }
    }

    class Solution {

        // 想法是构造一个树
        // 找根到叶子最短的路径
        public int integerReplacement(int n) {
            TreeNode root = new TreeNode(n);
            buildTree(root);
            return minDepth(root) - 1;
        }

        int minDepth(TreeNode node) {
            if (node == null) {
                return 0;
            }
            if (node.childs == null) {
                return 1;
            }
            int min = Integer.MAX_VALUE;
            for (TreeNode c : node.childs) {
                int depth = minDepth(c) + 1;
                min = Math.min(min, depth);
            }
            return min;
        }

        void buildTree(TreeNode node) {
            if (node.val == 1) {
                return;
            }
            node.childs = new ArrayList<>();
            if (node.val % 2 == 0) {
                TreeNode c = new TreeNode(node.val / 2);
                buildTree(c);
                node.childs.add(c);
            } else {
                TreeNode c1 = new TreeNode(node.val + 1);
                buildTree(c1);
                node.childs.add(c1);
                TreeNode c2 = new TreeNode(node.val - 1);
                buildTree(c2);
                node.childs.add(c2);
            }
        }
    }

    class Solution1 {
        public int integerReplacement(int n) {
            return (int) func((long) n);
        }

        public long func(long n) {
            if (n == 1)
                return 0;
            if (n % 2 == 0) {
                return 1 + func(n / 2);
            } else {
                return 1 + Math.min(func(n + 1), func(n - 1));
            }
        }
    }

    class Solution2 {
        int integerReplacement(int n) {
            long temp = n;
            int count = 0;
            while (temp != 1) {
                if ((temp & 3) == 3 && temp != 3) {
                    temp++;
                } else if ((temp & 1) == 1) {
                    temp--;
                } else {
                    temp = temp >> 1;
                }
                count++;
            }
            return count;
        }
    }

    class Solution3 {
        Map<Long, Long> memo = new HashMap<>();
        
        public int integerReplacement(int n) {
            return (int) func((long) n);
        }

        public long func(long n) {
            if (n == 1) {
                return 0;
            }
            if (!memo.containsKey(n)) {
                if (n % 2 == 0) {
                    memo.put(n, 1 + func(n / 2));
                } else {
                    memo.put(n, 1 + Math.min(func(n + 1), func(n - 1)));
                }
                
            } 
            return memo.get(n);
        }
    }

    public static void main(String[] args) {
        _397 _397 = new _397();
        Solution3 solution = _397.new Solution3();
        System.out.println(solution.integerReplacement(8)); // should be 3
        System.out.println(solution.integerReplacement(7)); // should be 4
        System.out.println(solution.integerReplacement(4)); // should be 2
    }
}

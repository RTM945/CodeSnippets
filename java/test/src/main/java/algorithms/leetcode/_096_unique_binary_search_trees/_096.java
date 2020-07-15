package algorithms.leetcode._096_unique_binary_search_trees;


//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//示例:
//输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//1         3     3      2      1
// \       /     /      / \      \
//  3     2     1      1   3      2
// /     /       \                 \
//2     1         2                 3
public class _096 {
    class Solution {
        //一时竟然没啥思路
        //然后又看到二叉搜索树
        //又看到做过95题
        //95比这个还难一点
        //用95的思路，对于二叉搜索树来说
        //左子节点比根小，右子节点比根大
        //如果选定了一个根，左边的取值范围和右边的取值范围都确定了
        //可以使用递归
        //再一看这题没有定义TreeNode，可能不需要使用树的结构就能做出来
        public int numTrees(int n) {
            return genTree(1, n);
        }

        int genTree(int start, int end) {
            int count = 0;
            if(start > end) {
                return 1; //子节点为空也算一种情况
            }
            for (int i = start; i <= end; i++) {
                int left = genTree(start, i - 1);
                int right = genTree(i + 1, end);
                count += left * right;
            }
            return count;
        }
    }

    class Solution1 {
        //跑到18的时候超时了
        //需要用dp吗
        //那就要找到递归中重复结算的地方将结果存储
        //n = 1时，结果1
        //n = 2时，结果2
        //n = 3时，结果5，应该包含了n = 1和2的所有排列情况
        //隐隐觉得用不到dp，而是像数列那样
        //假设n个节点存在二叉排序树的个数是G(n)，令f(i)为以i为根的二叉搜索树的个数，则
        //G(n) = f(1) + f(2) + f(3) + f(4) + ... + f(n)
        //当i为根节点时，其左子树节点个数为i - 1个，右子树节点为n - i，则
        //f(i) = G(i − 1) ∗ G(n − i)
        //G(n) = G(0) ∗ G(n − 1) + G(1) ∗ G(n − 2) + ... + G(n − 1) ∗ G(0)
        //数学上叫卡特兰数
        public int numTrees(int n) {
            int[] G = new int[n + 1];
            G[0] = 1;
            G[1] = 1;
            for (int i = 2; i <= n; i++) {
                for (int j = 1; j <= i; ++j) {
                    G[i] += G[j - 1] * G[i - j];
                }
            }
            return G[n];
        }
    }

    class Solution2 {
        //可以直接套数学公式
        //C(0) = 1 C(n + 1) = 2(2n + 1) / (n + 2) * C(n)
        public int numTrees(int n) {
            long C = 1;
            for (int i = 0; i < n; i++) {
                C = C * 2 * (2 * i + 1) / (i + 2);
            }
            return (int)C;
        }
    }

    public static void main(String[] args) {
        _096 q = new _096();
        System.out.println(q.new Solution().numTrees(3));
        System.out.println(q.new Solution().numTrees(4));
        System.out.println(q.new Solution().numTrees(18));
        System.out.println(q.new Solution1().numTrees(18));
        System.out.println(q.new Solution2().numTrees(18));
    }
}

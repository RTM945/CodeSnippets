package algorithms.leetcode._337_house_robber_iii;

import java.util.HashMap;
import java.util.Map;

// 在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
// 这个地区只有一个入口，我们称之为“根”。 
// 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。
// 一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 
// 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
// 示例 1:
// 输入: [3,2,3,null,3,null,1]
//      3
//     / \
//    2   3
//     \   \ 
//      3   1
// 输出: 7 
// 解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
// 示例 2:
// 输入: [3,4,5,1,3,null,1]
//      3
//     / \
//    4   5
//   / \   \ 
//  1   3   1
// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/house-robber-iii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
public class _337 {
    class Solution {
        // 据说很著名的打家劫舍系列题目
        // 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
        // 感觉问题可以转化成比较奇数层的和与偶数层的和的大小
        public int rob(TreeNode root) {
            if (root == null) {
                return 0;
            }
            int robRoot = root.val; // 从根开始抢
            // 下层
            TreeNode left = root.left;
            TreeNode right = root.right;
            if (left != null) {
                // 抢下下层左
                robRoot += rob(left.left) + rob(left.right);
            }
            if (right != null) {
                // 抢下下层右
                robRoot += rob(right.left) + rob(right.right);
            }

            // 不抢根
            int notRobRoot = rob(left) + rob(right);
            return Math.max(robRoot, notRobRoot);
        }
    }

    class Solution1 {
        // 上面的超时，那就记忆一下
        Map<TreeNode, Integer> map = new HashMap<>();

        public int rob(TreeNode root) {
            if (root == null) {
                return 0;
            }
            if (map.containsKey(root)) {
                return map.get(root);
            }
            int robRoot = root.val; // 从根开始抢
            // 下层
            TreeNode left = root.left;
            TreeNode right = root.right;
            if (left != null) {
                // 抢下下层左
                robRoot += rob(left.left) + rob(left.right);
            }
            if (right != null) {
                // 抢下下层右
                robRoot += rob(right.left) + rob(right.right);
            }

            // 不抢根
            int notRobRoot = rob(left) + rob(right);
            int max = Math.max(robRoot, notRobRoot);
            map.put(root, max);
            return max;
        }
    }

    class Solution2 {
        // 后序遍历
        // 取最后一层和倒数第二层的较大的和
        // 其实不是很好理解
        public int rob(TreeNode root) {
            if (root == null) {
                return 0;
            }
            postorder(root);
            return root.val;
        }

        public void postorder(TreeNode root) {
            if (root.left != null) {
                postorder(root.left);
            }
            if (root.right != null) {
                postorder(root.right);
            }
            int res1 = 0; //不抢本层，可以直接抢下层
            int res2 = root.val; //抢了本层，只能再抢下下层
            if (root.left != null) {
                res1 = res1 + root.left.val;
                if (root.left.left != null) {
                    res2 = res2 + root.left.left.val;
                }
                if (root.left.right != null) {
                    res2 = res2 + root.left.right.val;
                }
            }
            if (root.right != null) {
                res1 = res1 + root.right.val;
                if (root.right.left != null) {
                    res2 = res2 + root.right.left.val;
                }
                if (root.right.right != null) {
                    res2 = res2 + root.right.right.val;
                }
            }
            //将本层的值更新为最优解
            root.val = Math.max(res1, res2);
        }
    }
}
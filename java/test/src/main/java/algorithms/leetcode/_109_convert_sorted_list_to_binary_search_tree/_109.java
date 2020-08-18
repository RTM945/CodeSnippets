package algorithms.leetcode._109_convert_sorted_list_to_binary_search_tree;

/* 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
      0
     / \
   -3   9
   /   /
 -10  5
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _109 {
    class Solution {
        //有点眼熟，108是将一个数组转成平衡BST
        //这题是链表
        //当然可以把链表转成数组，直接复制粘贴108
        //怎么求一个链表中点
        //快慢指针法
        //就很人才，一个指针一次跑两个结点，另一个指针一次跑一个结点
        //当快的指针到终点时，慢的指针就在中点了
        //然后套用108，递归就行了
        public TreeNode sortedListToBST(ListNode head) {
            if (head == null) {
                return null;
            }
    
            if (head.next == null) {
                return new TreeNode(head.val);
            }
    
            // 快慢指针找中点
            ListNode p = head, q = head, pre = null;
            while (q != null && q.next != null) {
                pre = p;
                p = p.next;
                q = q.next.next;
            }
            pre.next = null; //断掉链表
            TreeNode root = new TreeNode(p.val); //p做根
            root.left = sortedListToBST(head); 
            root.right = sortedListToBST(p.next);
            return root;
        }

        //今天得知要进war room，心态崩了
    }
}
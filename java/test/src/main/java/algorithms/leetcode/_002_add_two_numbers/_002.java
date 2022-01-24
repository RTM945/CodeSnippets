package algorithms.leetcode._002_add_two_numbers;

public class _002 {
    // 给你两个 非空 的链表，表示两个非负的整数。
    // 它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

    // 请你将两个数相加，并以相同形式返回一个表示和的链表。

    // 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/add-two-numbers
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 暴力
    // 会溢出
    class Solution {
        public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
            long sum = getNumber(l1) + getNumber(l2);
            String str = String.valueOf(sum);
            ListNode l = new ListNode();
            l.val = str.charAt(0) - '0';
            // 头插
            for (int i = 1; i < str.length(); i++) {
                ListNode next = new ListNode(str.charAt(i) - '0');
                next.next = l;
                l = next;
            }

            return l;
        }

        long getNumber(ListNode l) {
            long a = l.val;
            int u = 1;
            while (l.next != null) {
                l = l.next;
                a += Math.pow(10, u++) * l.val;    
            }
            return a;
        }
    }

    class Solution1 {
        public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
            ListNode head = null, tail = null;
            int carry = 0;
            while (l1 != null || l2 != null) {
                int n1 = l1 != null ? l1.val : 0;
                int n2 = l2 != null ? l2.val : 0;
                int sum = n1 + n2 + carry;

                if (head == null) {
                    head = tail = new ListNode(sum % 10);
                } else {
                    tail.next = new ListNode(sum % 10);
                    tail = tail.next;
                }

                carry = sum / 10;
                if (l1 != null) {
                    l1 = l1.next;
                }
                if (l2 != null) {
                    l2 = l2.next;
                }
            }
            if (carry > 0) {
                tail.next = new ListNode(carry);
            }
            return head;
        }
    }
    
    public static void main(String[] args) {
        _002 _002 = new _002();
        Solution solution = _002.new Solution();
        ListNode l1 = new ListNode(9);
        ListNode l2 = new ListNode(1, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))))))))));
        System.out.println(solution.getNumber(l1));
        System.out.println(solution.getNumber(l2));
        System.out.println(solution.addTwoNumbers(l1, l2));
    }
}

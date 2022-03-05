package algorithms.leetcode._521_longest_uncommon_subsequence_i;

public class _521 {
    //  给你两个字符串 a 和 b，请返回 这两个字符串中 最长的特殊序列  的长度。
    // 如果不存在，则返回 -1 。

    // 「最长特殊序列」 定义如下：该序列为 某字符串独有的最长子序列（即不能是其他字符串的子序列） 。

    // 字符串 s 的子序列是在从 s 中删除任意数量的字符后可以获得的字符串。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/longest-uncommon-subsequence-i
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public int findLUSlength(String a, String b) {
            return a.equals(b) ? -1 : Math.max(a.length(), b.length());
        }
    }
}

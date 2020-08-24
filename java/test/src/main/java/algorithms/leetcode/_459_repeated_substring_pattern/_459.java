package algorithms.leetcode._459_repeated_substring_pattern;

/* 
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。
给定的字符串只含有小写英文字母，并且长度不超过10000。
示例 1:
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。
示例 2:
输入: "aba"
输出: False
示例 3:
输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/repeated-substring-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _459 {
    class Solution {
        // 看着不像easy的题目呢
        // 设s有个子串，s由这个子串重复构成的
        // s的长度子串的整数倍
        // i大于子串长度时，s[i] = s[i - 字串长度]
        // 子串长度不可能大于s的一半
        public boolean repeatedSubstringPattern(String s) {
            int n = s.length();
            for (int i = 0; i * 2 <= n; i++) {
                if (n % i == 0) {
                    boolean match = true;
                    for (int j = i; j < n; j++) {
                        if (s.charAt(j) != s.charAt(j - i)) {
                            match = false;
                            break;
                        }
                    }
                    if (match) {
                        return true;
                    }
                }
            }

            return false;
        }
    }

    class Solution1 {
        // 谜之解法
        // s如果有子串重复构成
        // 将一个子串剪下来拼到后面，还是s
        // 把两个s拼起来，去掉第一个和最后一个字符，会包含一个完整的s
        // 如果s满足要求，有这样的性质 （充分）
        // 如何证明s有这样的性质，则满足题目要求？ （必要）
        // 要用同余，下略
        public boolean repeatedSubstringPattern(String s) {
            return (s + s).substring(1, 2 * s.length() - 1).contains(s);
        }
    }

    //kmp的解法，不想花时间了，心累

    public static void main(String[] args) {
        _459 q = new _459();
        System.out.println(q.new Solution().repeatedSubstringPattern("abab"));
        System.out.println(q.new Solution().repeatedSubstringPattern("aba"));
        System.out.println(q.new Solution().repeatedSubstringPattern("abcabcabcabc"));
    }
}
package algorithms.leetcode._392_is_subsequence;

// https://leetcode-cn.com/problems/is-subsequence
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 你可以认为 s 和 t 中仅包含英文小写字母。
// 字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
//（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
// 示例 1:
// s = "abc", t = "ahbgdc"
// 返回 true.
// 示例 2:
// s = "axc", t = "ahbgdc"
// 返回 false.
// 后续挑战 :
// 如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。
// 在这种情况下，你会怎样改变代码？
// 致谢:
// 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
public class _392 {
    class Solution {
        //java的String.contains判断的是连续的子序列
        //这个题目是离散的子序列，但需要保持顺序
        //试试双指针
        public boolean isSubsequence(String s, String t) {
            if(s.length() == 0) {
                return true;
            }
            int ps = 0; //pointer of s
            for (int i = 0; i < t.length(); i++) {
                if(t.charAt(i) == s.charAt(ps)) {
                    ps++;
                    if(ps == s.length()) {
                        return true;
                    }
                }
            }
            return false;
        }
    }

    public static void main(String[] args) {
        _392 q = new _392();
        String s = "abc";
        String t = "ahbgdc";
        System.out.println(q.new Solution().isSubsequence(s, t));
    }
}
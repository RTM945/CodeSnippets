package algorithms.leetcode._044_wildcard_matching;

import java.util.ArrayList;
import java.util.List;

//https://leetcode-cn.com/problems/wildcard-matching/
//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//两个字符串完全匹配才算匹配成功。
//说明:
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
//示例 1:
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//示例 2:
//输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
//示例 3:
//输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
//示例 4:
//输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
//示例 5:
//输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false
public class _044 {
    class Solution{
        //老给我推hard是怎样..
        //看上去是实现一个简易正则
        //乍一看没啥思路
        //想起来很多程序中
        //比如传入一个文件，可以用*代替任意文件名，也不知道是怎么实现的
        //我首先理解输出可能为""，但不可能为null
        //如果p中不含*和?，则是s与p的全匹配
        //如果p中含通配符，则需要按通配符切割字符
        public boolean isMatch(String s, String p) {
            if(s == null || p == null) {
                return false;
            }
            if(!p.contains("*") && !p.contains("?")) {
                return s.equals(p);
            }
            //先切割p
            char[] chars = p.toCharArray();
            List<String> list = new ArrayList<>();
            StringBuilder tmp = new StringBuilder();
            for (int i = 0; i < chars.length; i++) {
                if(chars[i] == '*' || chars[i] == '?') {
                    //切分
                    list.add(tmp.toString());
                    tmp = new StringBuilder();
                    list.add(String.valueOf(chars[i]));
                }else{
                    tmp.append(chars[i]);
                }
            }
            //a c d c b
            //[a, *, c, ?]
            //比较
            //总觉得循环套循环很蠢，有没有更好的办法
            //到这里其实卡住了 实在想不出接下来的做法
            //
            return false;
        }
    }

    public static void main(String[] args) {
        _044 q = new _044();
        Solution s = q.new Solution();
        s.isMatch("acdcb", "a*cd?b");
    }
}

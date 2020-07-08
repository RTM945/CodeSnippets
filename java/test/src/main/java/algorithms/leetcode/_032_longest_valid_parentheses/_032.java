package algorithms.leetcode._032_longest_valid_parentheses;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

//32. 最长有效括号
//https://leetcode-cn.com/problems/longest-valid-parentheses/
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//示例 1:
//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//示例 2:
//输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
public class _032 {
    //第一次遇到hard题目，慌的一批
    //拿到题目有疑问，(())算几个？
    //群友说了两个字，栈，4
    //想了一下，可能是读到'('入栈，读到')'出栈
    //尝试一下
    class Solution {
        public int longestValidParentheses(String s) {
            LinkedList<Character> stack = new LinkedList<>();
            int max = 0;
            char[] arr = s.toCharArray();
            for (int i = 0; i < arr.length; i++) {
                char c = arr[i];
                if(c == '(') {
                    stack.push(c);
                }else{
                    if(stack.isEmpty()) {
                        continue;
                    }
                    char tmp = stack.pop();
                    if(tmp == '(') {
                        max += 2;//可以凑成一对
                    }
                }
            }
            return max;
        }
    }

    class Solution1 {
        //居然错了
        //要求"()(()"结果是2
        //也就是说要连续的()
        //那么(())只能算2个
        //我就说没这么简单
        //想法是搞个集合存储最次()的最大长度
        //当(且stack不为空时，当前的max存入集合，然后把max归零
        public int longestValidParentheses(String s) {
            LinkedList<Character> stack = new LinkedList<>();
            List<Integer> valids = new ArrayList<>();
            int max = 0;
            char[] arr = s.toCharArray();
            for (int i = 0; i < arr.length; i++) {
                char c = arr[i];
                if(c == '(') {
                    if(!stack.isEmpty()) {
                        valids.add(max);
                        max = 0;
                    }
                    stack.push(c);
                }else{
                    if(stack.isEmpty()) {
                        continue;
                    }
                    char tmp = stack.pop();
                    if(tmp == '(') {
                        max += 2;//可以凑成一对
                    }
                }
            }
            for(int m : valids) {
                if (m > max) {
                    max = m;
                }
            }
            return max;
        }
    }

    class Solution2 {
        //又错了
        //要求"()(())"输出6
        //真是奇了怪了
        //这样说(())还是4
        //将上面的程序修改，pop出来凑不出一对时，就需要重新计数了
        public int longestValidParentheses(String s) {
            LinkedList<Character> stack = new LinkedList<>();
            List<Integer> valids = new ArrayList<>();
            int max = 0;
            char[] arr = s.toCharArray();
            for (int i = 0; i < arr.length; i++) {
                char c = arr[i];
                if(c == '(') {
                    stack.push(c);
                }else{
                    if(stack.isEmpty()) {
                        continue;
                    }
                    char tmp = stack.pop();
                    if(tmp == '(') {
                        max += 2;//可以凑成一对
                    }else{
                        valids.add(max);
                        max = 0;
                    }
                }
            }
            for(int m : valids) {
                if (m > max) {
                    max = m;
                }
            }
            return max;
        }
    }

    class Solution3 {
        //仍然错了
        //要求"()(()"输出2
        //要求"()(())"输出6
        //不能简单的将max归零
        //答案的思路是存下标
        //其含义是最后一个没被匹配到'('的')'
        //'('时，放入下标
        //')'时，栈为空，放入下标，表示')'没有匹配项
        //栈不为空，当前下标减去栈顶元素+1即为结果
        //不过这样得出来"()(())"仍为4
        //答案说预先在栈内设置一个-1
        //为了满足提及的「最后一个没有被匹配的右括号的下标」
        //无法理解这一步，感觉是为了凑答案的做法
        public int longestValidParentheses(String s) {
            LinkedList<Integer> stack = new LinkedList<>();
            stack.push(-1);
            int max = 0;
            char[] arr = s.toCharArray();
            for (int i = 0; i < arr.length; i++) {
                char c = arr[i];
                if(c == '(') {
                    stack.push(i);
                }else{
//                    if (stack.isEmpty()) {
//                        stack.push(i);
//                    }else{
//                        int count = i - stack.pop() + 1; //()下标相减为1，需要凑成2
//                        if(count > max) {
//                            max = count;
//                        }
//                    }
                    stack.pop();
                    if (stack.isEmpty()) {
                        stack.push(i);
                    }else{
                        int count = i - stack.peek();//无法理解
                        if(count > max) {
                            max = count;
                        }
                    }
                }
            }
            return max;
        }
    }

    class Solution4{
        //看到其他使用栈且符合逻辑的做法，记录下来
        //对所有可以匹配的()标记0，无法匹配的标记为1
        //例如: "()(()"的mark为[0, 0, 1, 0, 0]
        //再例如: ")()((())"的mark为[1, 0, 0, 1, 0, 0, 0, 0]
        //经过这样的处理后, 此题就变成了寻找最长的连续的0的长度
        //也就是正常的入栈出栈，遗留下来的都要标记为1
        public int longestValidParentheses(String s) {
            LinkedList<Integer> stack = new LinkedList<>();
            int max = 0;
            char[] arr = s.toCharArray();
            int[] mark = new int[arr.length];
            for (int i = 0; i < arr.length; i++) {
                char c = arr[i];
                if(c == '(') {
                    stack.push(i);//如果匹配就会被pop，剩下未能匹配的左括号的下标
                }else{
                    if(stack.isEmpty()) {
                        //未能匹配的右括号
                        mark[i] = 1;
                    }else{
                        stack.pop();
                    }
                }
            }

            while (!stack.isEmpty()) {
                mark[stack.pop()] = 1;
            }

            //查找最大连续的0
            int len = 0;
            for (int i = 0; i < mark.length; i++) {
                if(mark[i] == 1) {
                    //连续断了，判断是否最大值
                    if(len > max) {
                        max = len;
                    }
                    len = 0;
                }else{
                    len++;
                    //判断是否最大值
                    if(len > max) {
                        max = len;
                    }
                }
            }
            return max;
        }
    }

    class Solution5 {
        //还有一种解法是从前往后加从后往前各扫描一次
        //目前不是很能理解，先记录下来
        public int longestValidParentheses(String s) {
            int left = 0, right = 0, max = 0;
            //从前往后右括号大于左括号，left和right都归0
            for (int i = 0; i < s.length(); i++) {
                if (s.charAt(i) == '(') {
                    left++;
                } else {
                    right++;
                }
                //left大于right先不管
                if (left == right) {
                    max = Math.max(max, 2 * right);
                } else if (right > left) {
                    //right大于left重新记
                    left = right = 0;
                }
            }
            left = right = 0;
            //从后往前左括号大于右括号，left和right都归0
            for (int i = s.length() - 1; i >= 0; i--) {
                if (s.charAt(i) == '(') {
                    left++;
                } else {
                    right++;
                }
                if (left == right) {
                    max = Math.max(max, 2 * left);
                } else if (left > right) {
                    left = right = 0;
                }
            }
            return max;
        }
    }

    public static void main(String[] args) {
        _032 q = new _032();
        String s = "()(()";
//        System.out.println(q.new Solution().longestValidParentheses(s));
//        System.out.println(q.new Solution1().longestValidParentheses(s));
        String s1 = "()(())";
//        System.out.println(q.new Solution2().longestValidParentheses(s));
//        System.out.println(q.new Solution2().longestValidParentheses(s1));
        System.out.println(q.new Solution4().longestValidParentheses(s));
        System.out.println(q.new Solution4().longestValidParentheses(s1));
    }
}

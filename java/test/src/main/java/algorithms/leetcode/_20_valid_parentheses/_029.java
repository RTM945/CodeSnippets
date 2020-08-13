package algorithms.leetcode._20_valid_parentheses;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.Map;

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 示例 1:
// 输入: "()"
// 输出: true
// 示例 2:
// 输入: "()[]{}"
// 输出: true
// 示例 3:
// 输入: "(]"
// 输出: false
// 示例 4:
// 输入: "([)]"
// 输出: false
// 示例 5:
// 输入: "{[]}"
// 输出: true
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
public class _029 {
    class Solution {
        // 简单是有原因的咯
        // 按上述规则，字符串长度一定为偶数，而且对称的咯
        // 甚至都不用判断是不是括号
        public boolean isValid(String s) {
            Map<Character, Character> map = new HashMap<>();
            map.put('(', ')');
            map.put('[', ']');
            map.put('{', '}');
            int n = s.length();
            if (n < 1 || n % 2 != 0) {
                return false;
            }
            int mid = n / 2;
            int left = mid - 1;
            int right = mid;
            // 判断左右对称
            while (left >= 0 && right <= n - 1) {
                if (map.get(s.charAt(left--)) != s.charAt(right++)) {
                    return false;
                }
            }
            return true;
        }
    }

    class Solution1 {
        // 简单翻车 sad
        // "()[]{}"没考虑
        // 那么应该也有()[{}]的情况
        // 如果有了任意一个左，在出现对应的右之前，不能出现其他右
        // 用栈吧
        public boolean isValid(String s) {
            LinkedList<Character> stack = new LinkedList<>();
            for (char c : s.toCharArray()) {
                if (c == '(' || c == '[' || c == '{') {
                    stack.push(c);
                } else {
                    if (stack.isEmpty()) {
                        return false;
                    }
                    char z = stack.pop();
                    if (z == '(' && c != ')') {
                        return false;
                    }
                    if (z == '[' && c != ']') {
                        return false;
                    }
                    if (z == '{' && c != '}') {
                        return false;
                    }
                }
            }
            return stack.isEmpty();
        }
    }

    class Solution2 {
        //更聪明的方法
        public boolean isValid(String s) {
            LinkedList<Character> stack = new LinkedList<>();
            for (char c : s.toCharArray()) {
                if (c == '(') {
                    stack.push(')');
                } else if (c == '[') {
                    stack.push(']');
                } else if (c == '{') {
                    stack.push('}');
                } else if (stack.isEmpty() || c != stack.pop()) {
                    return false;
                }
            }
            return stack.isEmpty();
        }
    }

    public static void main(String[] args) {
        _029 q = new _029();
        System.out.println(q.new Solution1().isValid("{[]}"));
        System.out.println(q.new Solution1().isValid("()"));
        System.out.println(q.new Solution1().isValid("([)]"));
        System.out.println(q.new Solution1().isValid("()[]{}"));
        System.out.println(q.new Solution1().isValid("()[{}]"));
    }
}
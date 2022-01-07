package algorithms.leetcode._1516_maximum_nesting_depth_of_the_parentheses;

public class _1516 {
    // 输入：s = "(1+(2*3)+((8)/4))+1"
    // 输出：3
    // 解释：数字 8 在嵌套的 3 层括号中。

    class Solution {
        public int maxDepth(String s) {
            int max = 0;
            int cnt = 0;
            for (char c : s.toCharArray()) {
                if (c == '(') {
                    cnt++;
                    max = Math.max(max, cnt);
                } else if (c == ')') {
                    cnt--;
                }
            }
    
            return max;
        }
    }
}

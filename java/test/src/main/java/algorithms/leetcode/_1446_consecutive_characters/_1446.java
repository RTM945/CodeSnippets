package algorithms.leetcode._1446_consecutive_characters;

public class _1446 {
    // 给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
    // 请你返回字符串的能量。    
    class Solution {
        public int maxPower(String s) {
            int max = 1;
            int count = 1;
            for (int i = 1; i < s.length(); i++) {
                if (s.charAt(i) == s.charAt(i - 1)) {
                    count++;
                    max = Math.max(max, count);
                }  else {
                    count = 1; // 重新计数
                }
                
            }
            return max;
        }
    }

    public static void main(String[] args) {
        _1446 _1446 = new _1446();
        Solution solution = _1446.new Solution();
        System.out.println(solution.maxPower("leetcode"));
    }
}

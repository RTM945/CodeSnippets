package algorithms.leetcode._917_reverse_only_letters;

public class _917 {
    // 给你一个字符串 s ，根据下述规则反转字符串：

    // 所有非英文字母保留在原有位置。
    // 所有英文字母（小写或大写）位置反转。
    // 返回反转后的 s 。
    class Solution {
        public String reverseOnlyLetters(String s) {
            char[] cs = s.toCharArray();
            int l = 0;
            int r = cs.length - 1;
            while(l < r) {
                while(l < r &&!Character.isLetter(s.charAt(l))) {
                    l++;
                }
                while(l < r &&!Character.isLetter(s.charAt(r))) {
                    r--;
                }
                if (l < r) {
                    swap(l, r, cs);
                }
                l++;
                r--;
            }
            return new String(cs);
        }

        private void swap(int i, int j, char[] cs) {
            char temp = cs[i];
            cs[i] = cs[j];
            cs[j] = temp;
        }
    }
}

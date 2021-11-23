package algorithms.leetcode._859_buddy_strings;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.checkerframework.checker.units.qual.g;

public class _859 {
    // https://leetcode-cn.com/problems/buddy-strings
    // 给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；
    // 否则返回 false 。
    // 交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，
    // 接着交换 s[i] 和 s[j] 处的字符。
    // 例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。
    class Solution {
        public boolean buddyStrings(String s, String goal) {
            if (s == null || goal == null) {
                return false;
            }
            if (s.length() != goal.length()) {
                return false;
            }
            // 暴力做法 超时
            for (int i = 0; i < s.length() - 1; i++) {
                for (int j = i + 1; j < s.length(); j++) {
                    char[] sarr = s.toCharArray();
                    char temp = sarr[i];
                    sarr[i] = sarr[j];
                    sarr[j] = temp;
                    if (new String(sarr).equals(goal)) {
                        return true;
                    }
                }
            }

            return false;
        }
    }

    class Solution1 {
        public boolean buddyStrings(String s, String goal) {
            if (s == null || goal == null) {
                return false;
            }
            if (s.length() != goal.length()) {
                return false;
            }
            if (s.equals(goal)) {
                Set<Character> set = new HashSet<>();
                for (int i = 0; i < s.length(); i++) {
                    set.add(s.charAt(i));
                }
                return set.size() < s.length();
            }

            List<Integer> idxs = new ArrayList<>();
            for (int i = 0; i < s.length(); i++) {
                if (s.charAt(i) != goal.charAt(i)) {
                    idxs.add(i);
                }
            }
            if (idxs.size() != 2) {
                return false;
            }
            // 判断交换后是否相等
            char[] arr = s.toCharArray();
            int i = idxs.get(0);
            int j = idxs.get(1);
            char tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
            return new String(arr).equals(goal);
        }
    }

    class Solution2 {
        public boolean buddyStrings(String s, String goal) {
            if (s.length() != goal.length()) {
                return false;
            }

            if (s.equals(goal)) {
                int[] count = new int[26];
                for (int i = 0; i < s.length(); i++) {
                    count[s.charAt(i) - 'a']++;
                    if (count[s.charAt(i) - 'a'] > 1) {
                        return true;
                    }
                }
                return false;
            } else {
                int first = -1, second = -1;
                for (int i = 0; i < goal.length(); i++) {
                    if (s.charAt(i) != goal.charAt(i)) {
                        if (first == -1) {
                            first = i;
                        } else if (second == -1) {
                            second = i;
                        } else {
                            return false;
                        }
                    }
                }

                return (second != -1 && s.charAt(first) == goal.charAt(second)
                        && s.charAt(second) == goal.charAt(first));
            }
        }
    }

    public static void main(String[] args) {
        _859 _859 = new _859();
        Solution solution = _859.new Solution();
        System.out.println(solution.buddyStrings("aa", "aa")); // should be true
        System.out.println(solution.buddyStrings("ab", "ba")); // should be true
        System.out.println(solution.buddyStrings("ab", "ab")); // should be false
        System.out.println(solution.buddyStrings("aaaaaaabc", "aaaaaaacb")); // should be true
        System.out.println(solution.buddyStrings("abac", "abad")); // should be true
    }
}

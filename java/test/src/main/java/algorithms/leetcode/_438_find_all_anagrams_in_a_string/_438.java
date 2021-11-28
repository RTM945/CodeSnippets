package algorithms.leetcode._438_find_all_anagrams_in_a_string;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class _438 {
    // 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，
    // 返回这些子串的起始索引。不考虑答案输出的顺序。

    // 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        // 想法是获得p的长度，s从前往后，只要匹配到p中的任一字母，就往后读几位
        // 如果和p的每个字母的数量相等，就输出该字母索引
        // 往后读几位看上去会重复计算
        public List<Integer> findAnagrams(String s, String p) {
            List<Integer> res = new ArrayList<>();
            int plen = p.length();
            int slen = s.length();
            if (slen < plen) {
                return res;
            }
            int[] pcount = new int[26];
            int[] scount = new int[26];
            for (int i = 0; i < plen; i++) {
                pcount[p.charAt(i) - 'a']++;
                scount[s.charAt(i) - 'a']++;
            }

            // 这里其实已经是从第一位开始了
            if (Arrays.equals(pcount, scount)) {
                res.add(0);
            }

            // 滑动窗口
            // 这里的比较是从s第二位开始，所以减掉第一位，往后加一位
            for (int i = 0; i < slen - plen; i++) {
                scount[s.charAt(i) - 'a']--;
                scount[s.charAt(i + plen) - 'a']++;

                if (Arrays.equals(scount, pcount)) {
                    res.add(i + 1);
                }
            }

            return res;
        }
    }

    public static void main(String[] args) {
        _438 _438 = new _438();
        Solution solution = _438.new Solution();
        System.out.println(solution.findAnagrams("cbaebabacd", "abc"));
    }
}

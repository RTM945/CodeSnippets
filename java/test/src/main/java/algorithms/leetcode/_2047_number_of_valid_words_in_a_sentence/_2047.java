package algorithms.leetcode._2047_number_of_valid_words_in_a_sentence;

public class _2047 {
    // 句子仅由小写字母（'a' 到 'z'）、数字（'0' 到 '9'）、连字符（'-'）、
    // 标点符号（'!'、'.' 和 ','）以及空格（' '）组成。
    // 每个句子可以根据空格分解成 一个或者多个 token ，这些 token 之间由一个或者多个空格 ' ' 分隔。

    // 如果一个 token 同时满足下述条件，则认为这个 token 是一个有效单词：

    // 仅由小写字母、连字符和/或标点（不含数字）。
    // 至多一个 连字符 '-' 。如果存在，连字符两侧应当都存在小写字母
    //（"a-b" 是一个有效单词，但 "-ab" 和 "ab-" 不是有效单词）。
    // 至多一个 标点符号。如果存在，标点符号应当位于 token 的 末尾 。
    // 这里给出几个有效单词的例子："a-b."、"afad"、"ba-c"、"a!" 和 "!" 。

    // 给你一个字符串 sentence ，请你找出并返回 sentence 中 有效单词的数目 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/number-of-valid-words-in-a-sentence
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public int countValidWords(String sentence) {
            String[] ss = sentence.split(" ");
            int ans = 0;
            for (String s : ss) {
                if (check(s)) {
                    ans++;
                }
            }
            return ans;
        }
        boolean check(String s) {
            int n = s.length();
            if (n == 0) {
                return false;
            }
            for (int i = 0, c1 = 0, c2 = 0; i < n; i++) {
                char c = s.charAt(i);
                if (Character.isDigit(c)) return false;
                if (c == ' ') {
                    return false;
                }
                if (c == '-' && ++c1 >= 0) {
                    if (c1 > 1 || (i == 0 || i == n - 1)) {
                        return false;
                    }
                    if (!Character.isLetter(s.charAt(i - 1)) || !Character.isLetter(s.charAt(i + 1))) {
                        return false;
                    }
                }
                if ((c == '!' || c == '.' || c == ',') && ++c2 >= 0) {
                    if (c2 > 1 || (i != n - 1)) {
                        return false;
                    }
                }
            }
            return true;
        }
    }

}

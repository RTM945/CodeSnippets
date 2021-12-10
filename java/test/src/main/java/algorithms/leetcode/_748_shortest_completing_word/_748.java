package algorithms.leetcode._748_shortest_completing_word;

import java.util.Arrays;

public class _748 {
    // 给你一个字符串 licensePlate 和一个字符串数组 words ，请你找出并返回 words 中的 最短补全词 。

    // 补全词 是一个包含 licensePlate 中所有的字母的单词。在所有补全词中，最短的那个就是 最短补全词 。

    // 在匹配 licensePlate 中的字母时：

    // 忽略 licensePlate 中的 数字和空格 。
    // 不区分大小写。
    // 如果某个字母在 licensePlate 中出现不止一次，那么该字母在补全词中的出现次数应当一致或者更多。
    // 例如：licensePlate = "aBc 12c"，那么它的补全词应当包含字母 'a'、'b' （忽略大写）和两个 'c' 。
    // 可能的 补全词 有 "abccdef"、"caaacab" 以及 "cbca" 。

    // 请你找出并返回 words 中的 最短补全词 。题目数据保证一定存在一个最短补全词。
    // 当有多个单词都符合最短补全词的匹配条件时取 words 中 最靠前的 那个。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/shortest-completing-word
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public String shortestCompletingWord(String licensePlate, String[] words) {
            int[] count = new int[26];
            String lowLicense = licensePlate.toLowerCase();
            for (int i = 0; i < lowLicense.length(); i++) {
                char c = lowLicense.charAt(i);
                if ('a'  <= c && c <= 'z') {
                    count[lowLicense.charAt(i) - 'a']++;
                }
            }
            // String res = null;
            int index = -1;
            // 尽量多的匹配 licensePlate 并尽可能短
            for (int i = 0; i < words.length; i++) {
                String word = words[i];
                int[] wcount = new int[26];
                for (int j = 0; j < word.length(); j++) {
                    wcount[word.charAt(j) - 'a']++;
                }
                boolean equal = true;
                for (int j = 0; j < count.length; j++) {
                    if (wcount[j] < count[j]) {
                        equal = false;
                        break;
                    }   
                }

                // if (equal) {
                //     res = res == null ? word : res.length() > word.length() ? word : res;
                // }
                if (equal && (index < 0 || words[index].length() > word.length())) {
                    index = i;
                }
            }

            return words[index];
        }
    }

    public static void main(String[] args) {
        _748 _748 = new _748();
        Solution solution = _748.new Solution();
        // step
        System.out.println(solution.shortestCompletingWord("1s3 PSt", new String[]{"steps", "step", "stripe", "stepple"}));
        // pest
        System.out.println(solution.shortestCompletingWord("1s3 456", new String[]{"looks", "pest", "stew", "show"}));
        // "according"
        System.out.println(solution.shortestCompletingWord("GrC8950", new String[]{"measure","other","every","base","according","level","meeting","none","marriage","rest"}));
    
    }
}

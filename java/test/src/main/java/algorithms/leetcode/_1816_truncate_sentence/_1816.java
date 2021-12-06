package algorithms.leetcode._1816_truncate_sentence;

import java.util.Arrays;
import java.util.stream.Collectors;

public class _1816 {
    // 句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。
    // 每个单词仅由大小写英文字母组成（不含标点符号）。
    // 例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。
    // 给你一个句子 s​​​​​​ 和一个整数 k​​​​​​ ，请你将 s​​ 截断 ​，​​​使截断后的句子仅含 前 k​​​​​​ 个单词。
    // 返回 截断 s​​​​​​ 后得到的句子。
    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/truncate-sentence
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public String truncateSentence(String s, int k) {
            return Arrays.asList(s.split(" ")).stream().limit(k).collect(Collectors.joining(" ", "", ""));    
        }
    }

    class Solution1 {
        public String truncateSentence(String s, int k) {
            int n = s.length();
            int end = 0, count = 0;
            for (int i = 1; i <= n; i++) {
                if (i == n || s.charAt(i) == ' ') {
                    count++;
                    if (count == k) {
                        end = i;
                        break;
                    }
                }
            }
            return s.substring(0, end);
        }
    }

    public static void main(String[] args) {
        _1816 _1816 = new _1816();
        Solution solution = _1816.new Solution();
        System.out.println(solution.truncateSentence("Hello how are you Contestant", 4));
        System.out.println(solution.truncateSentence("What is the solution to this problem", 4));
        System.out.println(solution.truncateSentence("chopper is not a tanuki", 5));
    }
}

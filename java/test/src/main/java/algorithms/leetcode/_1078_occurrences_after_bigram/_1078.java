package algorithms.leetcode._1078_occurrences_after_bigram;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class _1078 {
    // 给出第一个词 first 和第二个词 second，考虑在某些文本 text 中可能以 "first second third" 形式出现的情况，其中 second 紧随 first 出现，third 紧随 second 出现。

    // 对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/occurrences-after-bigram
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public String[] findOcurrences(String text, String first, String second) {
            String[] words = text.split(" ");
            List<String> list = new ArrayList<>();
            for (int i = 0; i < words.length - 2; i++) {
                if (first.equals(words[i]) && second.equals(words[i + 1])) {
                    list.add(words[i + 2]);
                }
            }
            return list.toArray(new String[0]);
        }
    }

    public static void main(String[] args) {
        _1078 _1078 = new _1078();
        Solution solution = _1078.new Solution();
        System.out.println(Arrays.toString(solution.findOcurrences("alice is a good girl she is a good student", "a", "good")));
    }
}

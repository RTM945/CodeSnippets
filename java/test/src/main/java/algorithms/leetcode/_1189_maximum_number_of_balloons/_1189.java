package algorithms.leetcode._1189_maximum_number_of_balloons;

import java.util.Arrays;

public class _1189 {
    // 给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。

    // 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/maximum-number-of-balloons
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public int maxNumberOfBalloons(String text) {
            // count每个字母取最小值
            int[] index = {1, 0, 'l' - 'a', 'l' - 'a', 'o' - 'a', 'o' - 'a', 'n' - 'a'};
            int[] cnt = new int[26];
            for (int i = 0; i < text.length(); i++) {
                cnt[text.charAt(i) - 'a']++;
            }
            int ans = 0;
            while(true) {
                for (int i = 0; i < index.length; i++) {
                    int c = cnt[index[i]]--;
                    if (c < 1) {
                        return ans;
                    }
                }
                ans++;
            }
        }
    }

    class Solution1 {
        public int maxNumberOfBalloons(String text) {
            int[] cnt = new int[5];
            for (int i = 0; i < text.length(); ++i) {
                char ch = text.charAt(i);
                if (ch == 'b') {
                    cnt[0]++;
                } else if (ch == 'a') {
                    cnt[1]++;
                } else if (ch == 'l') {
                    cnt[2]++;
                } else if (ch == 'o') {
                    cnt[3]++;
                } else if (ch == 'n') {
                    cnt[4]++;
                }
            }
            cnt[2] /= 2;
            cnt[3] /= 2;
            return Arrays.stream(cnt).min().getAsInt();
        }
    }

    public static void main(String[] args) {
        _1189 _1189 = new _1189();
        Solution solution = _1189.new Solution();
        System.out.println(solution.maxNumberOfBalloons("nlaebolko")); //1
        System.out.println(solution.maxNumberOfBalloons("loonbalxballpoon")); //2
        System.out.println(solution.maxNumberOfBalloons("leetcode")); //0
        System.out.println(solution.maxNumberOfBalloons("balon")); //0
    }
}

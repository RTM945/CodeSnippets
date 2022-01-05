package algorithms.leetcode._1576_replace_all_s_to_avoid_consecutive_repeating_characters;

public class _1576 {
    // 给你一个仅包含小写英文字母和 '?' 字符的字符串 s，请你将所有的 '?' 转换为若干小写字母，
    // 使最终的字符串不包含任何 连续重复 的字符。

    // 注意：你 不能 修改非 '?' 字符。

    // 题目测试用例保证 除 '?' 字符 之外，不存在连续重复的字符。

    // 在完成所有转换（可能无需转换）后返回最终的字符串。如果有多个解决方案，请返回其中任何一个。
    // 可以证明，在给定的约束条件下，答案总是存在的。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 不存在连续重复的字符
    // 第 i 个字符为 ? 时 在 'a' - 'z' 中找到与s[i-1]和s[i+1]不同的字符即可
    // 实际不需要遍历所有的小写字母，只需要遍历三个互不相同的字母，
    // 就能保证一定找到一个与前后字符均不相同的字母，在此我们可以限定三个不同的字母为 a b c
    class Solution {
        public String modifyString(String s) {
            int n = s.length();
            char[] arr = s.toCharArray();
            for (int i = 0; i < n; ++i) {
                if (arr[i] == '?') {
                    for (char ch = 'a'; ch <= 'c'; ++ch) {
                        if ((i > 0 && arr[i - 1] == ch) || (i < n - 1 && arr[i + 1] == ch)) {
                            continue;
                        }
                        arr[i] = ch;
                        break;
                    }
                }
            }
            return new String(arr);
        }
    }
    
    public static void main(String[] args) {
        _1576 _1576 = new _1576();
        Solution solution = _1576.new Solution();
        System.out.println(solution.modifyString("?zs"));
        System.out.println(solution.modifyString("ubv?w"));
        System.out.println(solution.modifyString("j?qg??b"));
        System.out.println(solution.modifyString("??yw?ipkj?"));
    }
}

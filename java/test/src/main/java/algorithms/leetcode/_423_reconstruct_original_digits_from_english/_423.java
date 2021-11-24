package algorithms.leetcode._423_reconstruct_original_digits_from_english;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.stream.Collectors;

public class _423 {
    // 给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
    // https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/
    // s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
    // 输入：s = "owoztneoer"
    // 输出："012"
    // 测试用例14错了 看上去有优先级
    class Solution {
        // 想构造字典树，但感觉好麻烦
        public String originalDigits(String s) {
            // 0 到 9 的 英文
            Map<String, Integer> words2Int = new LinkedHashMap<>();
            String[] nums = { "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };
            for (int i = 0; i < nums.length; i++) {
                words2Int.put(nums[i], i);
            }
            char[] cs = new char[] { 'e', 'g', 'f', 'i', 'h', 'o', 'n', 's', 'r', 'u', 't', 'w', 'v', 'x', 'z' };
            Map<Character, List<String>> map = new TreeMap<>();
            Map<Character, Integer> charsCount = new TreeMap<>();
            for (char c : cs) {
                List<String> list = map.computeIfAbsent(c, k -> new ArrayList<>());
                for (String num : nums) {
                    if (num.indexOf(c) > -1) {
                        list.add(num);
                    }
                }
            }
            char[] sarr = s.toCharArray();
            for (char a : sarr) {
                int count = charsCount.getOrDefault(a, 0);
                charsCount.put(a, ++count);
            }

            List<String> res = new ArrayList<>();
            for (char a : sarr) {
                List<String> words = map.get(a);
                int left = charsCount.get(a);
                if (left == 0) {
                    continue;
                }
                for (String word : words) {
                    boolean pass = true;
                    for (char c : word.toCharArray()) {
                        int count = charsCount.getOrDefault(c, 0);
                        if (count - 1 < 0) {
                            pass = false;
                            break;
                        }
                    }
                    if (!pass) {
                        continue;
                    }
                    for (char c : word.toCharArray()) {
                        charsCount.put(c, charsCount.get(c) - 1);
                    }
                    res.add(word);
                }
            }
            return res.stream().map(word -> words2Int.get(word)).sorted().map(String::valueOf)
                    .collect(Collectors.joining("", "", ""));
        }
    }

    // 检测每个字符在单词中的唯一性
    class Solution1 {
        public String originalDigits(String s) {
            Map<Character, Integer> c = new HashMap<>();
            for (int i = 0; i < s.length(); ++i) {
                char ch = s.charAt(i);
                c.put(ch, c.getOrDefault(ch, 0) + 1);
            }

            int[] cnt = new int[10];
            cnt[0] = c.getOrDefault('z', 0);
            cnt[2] = c.getOrDefault('w', 0);
            cnt[4] = c.getOrDefault('u', 0);
            cnt[6] = c.getOrDefault('x', 0);
            cnt[8] = c.getOrDefault('g', 0);

            cnt[3] = c.getOrDefault('h', 0) - cnt[8];
            cnt[5] = c.getOrDefault('f', 0) - cnt[4];
            cnt[7] = c.getOrDefault('s', 0) - cnt[6];

            cnt[1] = c.getOrDefault('o', 0) - cnt[0] - cnt[2] - cnt[4];

            cnt[9] = c.getOrDefault('i', 0) - cnt[5] - cnt[6] - cnt[8];

            StringBuffer ans = new StringBuffer();
            for (int i = 0; i < 10; ++i) {
                for (int j = 0; j < cnt[i]; ++j) {
                    ans.append((char) (i + '0'));
                }
            }
            return ans.toString();
        }
    }

    public static void main(String[] args) {
        _423 _423 = new _423();
        Solution solution = _423.new Solution();
        System.out.println(solution.originalDigits("owoztneoer")); // 012
        System.out.println(solution.originalDigits("fviefuro")); // 45
        System.out.println(solution.originalDigits("esnve")); // 7
        System.out.println(solution.originalDigits("zerozero")); // 00

    }
}

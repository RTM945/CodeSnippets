package algorithms.leetcode._318_maximum_product_of_word_lengths;

import java.util.Collections;
import java.util.HashSet;
import java.util.Set;

// https://leetcode-cn.com/problems/maximum-product-of-word-lengths
public class _318 {
    // 给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
    // 并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。
    // 如果不存在这样的两个单词，返回 0。
    class Solution {
        // 先找出不含公共字母的单词

        // 暴力
        public int maxProduct(String[] words) {
            int max = 0;
            for (int i = 0; i < words.length; i++) {
                String a = words[i];
                for (int j = i + 1; j < words.length; j++) {
                    String b = words[j];
                    char[] aarr = a.toCharArray();
                    char[] barr = b.toCharArray();
                    boolean pass = true;
                    for (int k = 0; k < aarr.length; k++) {
                        for (int k2 = 0; k2 < barr.length; k2++) {
                            if (aarr[k] == barr[k2]) {
                                pass = false;
                                break;
                            }
                        }
                        if (!pass) {
                            break;
                        }
                    }
                    if (pass) {
                        // 无重复 算乘积
                        int mul = a.length() * b.length();
                        if (mul > max) {
                            max = mul;
                        }
                    }
                }
            }
            return max;
        }
    }

    class Solution2 {
        public int maxProduct(String[] words) {
            Set<String>[] chars = new Set[words.length];
            for (int i = 0; i < words.length; i++) {
                chars[i] = new HashSet<>();
                Collections.addAll(chars[i], words[i].split(""));
            }

            int result = 0;
            for (int i = 0; i < words.length; i++) {
                for (int j = i + 1; j < words.length; j++) {
                    // 不相交，没有相同元素
                    if (Collections.disjoint(chars[i], chars[j])) {
                        result = Math.max(result, words[i].length() * words[j].length());
                    }
                }
            }
            return result;
        }
    }

    class Solution3 {
        public int maxProduct(String[] words) {
            int length = words.length;
            int[] masks = new int[length];
            for (int i = 0; i < length; i++) {
                String word = words[i];
                int wordLength = word.length();
                for (int j = 0; j < wordLength; j++) {
                    masks[i] |= 1 << (word.charAt(j) - 'a');
                }
            }
            int maxProd = 0;
            for (int i = 0; i < length; i++) {
                for (int j = i + 1; j < length; j++) {
                    if ((masks[i] & masks[j]) == 0) {
                        maxProd = Math.max(maxProd, words[i].length() * words[j].length());
                    }
                }
            }
            return maxProd;
        }
    }


    public static void main(String[] args) {
        _318.Solution solution = new _318().new Solution();
        String[] test1 = { "abcw", "baz", "foo", "bar", "xtfn", "abcdef" };
        int res = solution.maxProduct(test1);
        System.out.println(res); // should be 16
        String[] test2 = { "a", "ab", "abc", "d", "cd", "bcd", "abcd" };
        res = solution.maxProduct(test2);
        System.out.println(res); // should be 4
        String[] test3 = { "a", "aa", "aaa", "aaaa" };
        res = solution.maxProduct(test3);
        System.out.println(res); // should be 0
    }
}

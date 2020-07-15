package algorithms.leetcode.interview_17_13_re_space_lcci;

import java.util.*;

//https://leetcode-cn.com/problems/re-space-lcci/
//哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。
//像句子"I reset the computer. It still didn’t boot!"
//已经变成了"iresetthecomputeritstilldidntboot"。
//在处理标点符号和大小写之前，你得先把它断成词语。
//当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。
//假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。
//注意：本题相对原题稍作改动，只需返回未识别的字符数
//示例：
//输入：
//dictionary = ["looked","just","like","her","brother"]
//sentence = "jesslookedjustliketimherbrother"
//输出： 7
//解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
//提示：
//0 <= len(sentence) <= 1000
//dictionary中总字符数不超过 150000。
//你可以认为dictionary和sentence中只包含小写字母。
public class _17_13 {

    class Solution {
        //第一时间可以想到java的contains/indexOf API
        //觉得真要用他们性能肯定不行
        //但没有其他思路，先实现看看
        //如果匹配到了字典里的词，需要在其前后加空格
        //使用indexOf，在位置处截断字符串，添加空格后拼接
        //可以将dictionary转为hashmap，计算最后没被匹配的数字
        //her和brother中的her该如何处理?
        //题目要求返回的未识别数字越少越好
        //那么应该从长的单词开始匹配
        //那么需要dictionary重排序，长的单词在前
        //感觉太繁琐了，思路肯定不对，要想其他办法
        public int respace(String[] dictionary, String sentence) {
            Map<String, Boolean> map = new HashMap<>();
            int res = 0;
            for (String word : dictionary) {
                map.put(word, true);
                int index = sentence.indexOf(word);
                while (index > -1) {
                    String str = sentence.substring(0, index);
                    if (index == 0) {
                        str += word + " ";
                    } else {
                        if (sentence.charAt(index - 1) == ' ') {
                            str += word + " ";
                        } else {
                            str += " " + word + " ";
                        }
                    }
                    str += sentence.substring(index + word.length());
                    sentence = str;
                    index = sentence.indexOf(word, index + word.length() + 1);
                }
            }
            System.out.println(sentence);
            String[] words = sentence.split(" ");
            for (String word : words) {
                if (!map.containsKey(word)) {
                    res += word.length();
                }
            }
            return res;
        }
    }

    class Solution1 {
        //看答案吧
        //https://leetcode-cn.com/problems/re-space-lcci/solution/cong-bao-li-ru-shou-you-hua-yi-ji-triezi-dian-shu-/
        //dp，都可以dp
        //可以注意到dp关注的核心点不是匹配单词分割
        //而是最少未被匹配的字符
        //所以还是个最优解的问题
        //dp[i]表示考虑前 i 个字符最少的未识别的字符数量，从前往后计算 dp 值
        //即dp[0] = 0表示句子是空字符串时没有未识别的字符
        //对于前i个字符，即sentence[0,i)，它可能是由最前面的[0,j)子字符串加上一个字典匹配的单词得到
        //也就是dp[i]=dp[j], j<i
        //也可能没找到字典中的单词，可以用它前i-1个字符的结果加上一个没有匹配到的第i个字符，即dp[i]=dp[i-1]+1
        //即使前面存在匹配的单词，也不能保证哪一种剩下的字符最少，所以每轮都要比较一次最小值。
        public int respace(String[] dictionary, String sentence) {
            Set<String> dic = new HashSet<>(Arrays.asList(dictionary));
            int n = sentence.length();
            //dp[i]表示sentence前i个字符所得结果
            int[] dp = new int[n + 1];
            for (int i = 1; i <= n; i++) {
                dp[i] = dp[i - 1] + 1;  //先假设当前字符作为单词不在字典中
                for (int j = 0; j < i; j++) {
                    if (dic.contains(sentence.substring(j, i))) {
                        dp[i] = Math.min(dp[i], dp[j]);
                    }
                }
            }
            return dp[n];
        }
    }

    public static void main(String[] args) {
        _17_13 q = new _17_13();
        String[] dictionary = {"looked", "just", "like", "her", "brother"};
        String sentence = "jesslookedjustliketimherbrother";
        System.out.println(q.new Solution().respace(dictionary, sentence));
        System.out.println(q.new Solution1().respace(dictionary, sentence));
    }
}

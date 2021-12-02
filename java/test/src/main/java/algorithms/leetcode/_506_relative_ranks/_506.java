package algorithms.leetcode._506_relative_ranks;

import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

public class _506 {
    // 给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。
    // 所有得分都 互不相同 。

    // 运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，
    // 依此类推。运动员的名次决定了他们的获奖情况：

    // 名次第 1 的运动员获金牌 "Gold Medal" 。
    // 名次第 2 的运动员获银牌 "Silver Medal" 。
    // 名次第 3 的运动员获铜牌 "Bronze Medal" 。
    // 从名次第 4 到第 n 的运动员，只能获得他们的名次编号（即，名次第 x 的运动员获得编号 "x"）。
    // 使用长度为 n 的数组 answer 返回获奖，其中 answer[i] 是第 i 位运动员的获奖情况。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/relative-ranks
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public String[] findRelativeRanks(int[] score) {
            TreeMap<Integer, Integer> map = new TreeMap<>(Comparator.reverseOrder());
            for (int i = 0; i < score.length; i++) {
                map.put(score[i], i);
            }
            Map<Integer, String> temp = new HashMap<>();
            temp.put(1, "Gold Medal");
            temp.put(2, "Silver Medal");
            temp.put(3, "Bronze Medal");
            String[] ranks = new String[score.length];
            int i = 1;
            for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
                ranks[entry.getValue()] = temp.getOrDefault(i, String.valueOf(i));
                i++;
            }
            return ranks;
        }
    }

    public static void main(String[] args) {
        // 输入：score = [10,3,8,9,4]
        // 输出：["Gold Medal","5","Bronze Medal","Silver Medal","4"]
        // 解释：名次为 [1st, 5th, 3rd, 2nd, 4th] 。
        _506 _506 = new _506();
        System.out.println(Arrays.toString(_506.new Solution().findRelativeRanks(new int[]{10,3,8,9,4})));
    }
}

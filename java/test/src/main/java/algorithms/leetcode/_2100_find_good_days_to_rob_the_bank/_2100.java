package algorithms.leetcode._2100_find_good_days_to_rob_the_bank;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class _2100 {
    // 你和一群强盗准备打劫银行。给你一个下标从 0 开始的整数数组 security ，
    // 其中 security[i] 是第 i 天执勤警卫的数量。日子从 0 开始编号。
    // 同时给你一个整数 time 。

    // 如果第 i 天满足以下所有条件，我们称它为一个适合打劫银行的日子：

    // 第 i 天前和后都分别至少有 time 天。
    // 第 i 天前连续 time 天警卫数目都是非递增的。
    // 第 i 天后连续 time 天警卫数目都是非递减的。
    // 更正式的，第 i 天是一个合适打劫银行的日子当且仅当：
    // security[i - time] >= security[i - time + 1] >= ... 
    // >= security[i] <= ... <= security[i + time - 1] <= security[i + time].

    // 请你返回一个数组，包含 所有 适合打劫银行的日子（下标从 0 开始）。
    // 返回的日子可以 任意 顺序排列。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/find-good-days-to-rob-the-bank
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 记录
    class Solution {
        public List<Integer> goodDaysToRobBank(int[] security, int time) {
            int n = security.length;
            int[] left = new int[n];
            int[] right = new int[n];
            for (int i = 1; i < right.length; i++) {
                // 从前往后连续非递增的天数
                if (security[i - 1] >= security[i]) {
                    left[i] = left[i - 1] + 1; 
                }
                // 从后往前连续非递减的天数
                if (security[n - i - 1] <= security[n - i]) {
                    right[n - i - 1] = right[n - i] + 1; 
                }
            }
            List<Integer> res = new ArrayList<>();
            for (int i = time; i < n - time; i++) {
                if (left[i] >= time && right[i] >= time) {
                    res.add(i);    
                }
            }
            return res;
        }
    }

    public static void main(String[] args) {
        _2100 _2100 = new _2100();
        Solution solution = _2100.new Solution();
        System.out.println(solution.goodDaysToRobBank(new int[] {5,3,3,3,5,6,2}, 2));
    }
}

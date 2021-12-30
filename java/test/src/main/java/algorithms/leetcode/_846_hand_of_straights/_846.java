package algorithms.leetcode._846_hand_of_straights;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class _846 {
    // Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，
    // 使每一组的牌数都是 groupSize ，并且由 groupSize 张连续的牌组成。

    // 给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。
    // 如果她可能重新排列这些牌，返回 true ；否则，返回 false 。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/hand-of-straights
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 已经连续了 滑动窗口可以吗
    class Solution {
        public boolean isNStraightHand(int[] hand, int groupSize) {
            int n = hand.length;
            if (n % groupSize != 0) {
                return false;
            }
            Arrays.sort(hand);
            Map<Integer, Integer> cnt = new HashMap<>();
            for (int x : hand) {
                cnt.put(x, cnt.getOrDefault(x, 0) + 1);
            }
            for (int x : hand) {
                if (!cnt.containsKey(x)) {
                    continue;
                }
                for (int j = 0; j < groupSize; j++) {
                    int num = x + j;
                    if (!cnt.containsKey(num)) {
                        return false;
                    }
                    cnt.put(num, cnt.get(num) - 1);
                    if (cnt.get(num) == 0) {
                        cnt.remove(num);
                    }
                }
            }
            return true;
        }
    }
    
    public static void main(String[] args) {
        _846 _846 = new _846();
        Solution solution = _846.new Solution();
        System.out.println(solution.isNStraightHand(new int[]{1,2,3,6,2,3,4,7,8}, 3));
    }
}

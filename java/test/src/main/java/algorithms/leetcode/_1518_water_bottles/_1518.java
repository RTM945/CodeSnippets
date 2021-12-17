package algorithms.leetcode._1518_water_bottles;

import java.util.LinkedList;

public class _1518 {
    // 小区便利店正在促销，用 numExchange 个空酒瓶可以兑换一瓶新酒。你购入了 numBottles 瓶酒。

    // 如果喝掉了酒瓶中的酒，那么酒瓶就会变成空的。

    // 请你计算 最多 能喝到多少瓶酒。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/water-bottles
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 用队列咯
    class Solution {
        public int numWaterBottles(int numBottles, int numExchange) {
            LinkedList<Integer> queue = new LinkedList<>();
            for (int i = 0; i < numBottles; i++) {
                queue.add(i);
            }
            int result = numBottles;
            while(!queue.isEmpty()) {
                for (int i = 0; i < numExchange; i++) {
                    try {
                        queue.removeFirst();
                    } catch (Exception e) {
                        return result;
                    }
                    
                }
                result++;
                queue.add(1);
            }

            return result;
        }
    }

    class Solution1 {
        public int numWaterBottles(int numBottles, int numExchange) {
            int bottle = numBottles, ans = numBottles;
            while (bottle >= numExchange) {
                bottle -= numExchange;
                ++ans;
                ++bottle;
            }
            return ans;
        }
    }

    public static void main(String[] args) {
        _1518 _1518 = new _1518();
        Solution solution = _1518.new Solution();
        System.out.println(solution.numWaterBottles(9, 3)); //13
    }
}

package algorithms.leetcode._319_bulb_switcher;

import java.util.Arrays;

// https://leetcode-cn.com/problems/bulb-switcher/
// 初始时有 n 个灯泡处于关闭状态。第一轮，你将会打开所有灯泡。接下来的第二轮，你将会每两个灯泡关闭一个。
// 第三轮，你每三个灯泡就切换一个灯泡的开关（即，打开变关闭，关闭变打开）。第 i 轮，你每 i 个灯泡就切换一个灯泡的开关。直到第 n 轮，你只需要切换最后一个灯泡的开关。
// 找出并返回 n 轮后有多少个亮着的灯泡。
// 示例1
// 输入：n = 3
// 输出：1 
// 解释：
// 初始时, 灯泡状态 [关闭, 关闭, 关闭].
// 第一轮后, 灯泡状态 [开启, 开启, 开启].
// 第二轮后, 灯泡状态 [开启, 关闭, 开启].
// 第三轮后, 灯泡状态 [开启, 关闭, 关闭]. 
// 你应该返回 1，因为只有一个灯泡还亮着。
public class _319 {
    // 中等
    // 纯按逻辑写会超时
    // 最后变成了数学问题
    class Solution {
    
        public int bulbSwitch(int n) {
            // 全部开启
            byte[] bulbs = new byte[n];
            Arrays.fill(bulbs, (byte) 1);
            // 第二轮开始
            for(int i = 2; i <= n; i++) {
                // 第 i 轮 每 i 个关闭
                for(int j = i; j <= n; j = j + i){
                    bulbs[j - 1] ^= 1;
                }
            }
            int count = 0;
            for (int i = 0; i < n; i++) {
                if (bulbs[i] == 1) {
                    count++;
                }
            }
            return count;
            //return (int) Math.sqrt(n);
        }
    
    }

    public static void main(String[] args) {
        _319 q = new _319();
        System.out.println(q.new Solution().bulbSwitch(3));
    }
}


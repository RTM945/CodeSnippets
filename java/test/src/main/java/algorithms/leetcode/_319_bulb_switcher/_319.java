package algorithms.leetcode._319_bulb_switcher;

import java.util.Arrays;

public class _319 {
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


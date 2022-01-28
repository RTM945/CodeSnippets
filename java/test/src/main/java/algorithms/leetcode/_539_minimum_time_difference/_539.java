package algorithms.leetcode._539_minimum_time_difference;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import com.google.common.collect.Lists;

public class _539 {
    // 给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，
    // 找出列表中任意两个时间的最小时间差并以分钟数表示。
    // 输入：timePoints = ["23:59","00:00"]
    // 输出：1

    // 作者：LeetCode-Solution
    // 链接：https://leetcode-cn.com/problems/minimum-time-difference/solution/zui-xiao-shi-jian-chai-by-leetcode-solut-xolj/
    // 来源：力扣（LeetCode）
    // 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
    class Solution {
        public int findMinDifference(List<String> timePoints) {
            int n = timePoints.size();
            if (n > 1440) {
                return 0;
            }
            List<Integer> times = new ArrayList<>();
            for (String time : timePoints) {
                String[] strs = time.split(":");
                int min = Integer.parseInt(strs[0]) * 60 + Integer.parseInt(strs[1]);
                times.add(min);
                times.add(min + 1440);
            }
            Collections.sort(times);
            
            int min = 1440;
            for (int i = 0; i < n - 1; i++) {
                min = Math.min(Math.abs(times.get(i + 1) - times.get(i)), min);
            }
            return min;
        }
    }

    class Solution1 {
        public int findMinDifference(List<String> timePoints) {
            int n = timePoints.size();
            if (n > 1440) {
                return 0;
            }
            Collections.sort(timePoints);
            int ans = Integer.MAX_VALUE;
            int t0Minutes = getMinutes(timePoints.get(0));
            int preMinutes = t0Minutes;
            for (int i = 1; i < n; ++i) {
                int minutes = getMinutes(timePoints.get(i));
                ans = Math.min(ans, minutes - preMinutes); // 相邻时间的时间差
                preMinutes = minutes;
            }
            ans = Math.min(ans, t0Minutes + 1440 - preMinutes); // 首尾时间的时间差
            return ans;
        }
    
        public int getMinutes(String t) {
            return ((t.charAt(0) - '0') * 10 + (t.charAt(1) - '0')) * 60 + (t.charAt(3) - '0') * 10 + (t.charAt(4) - '0');
        }
    }
    

    public static void main(String[] args) {
        _539 _539 = new _539();
        Solution solution = _539.new Solution();
        // System.out.println(solution.findMinDifference(Lists.newArrayList("23:59","00:00")));
        // System.out.println(solution.findMinDifference(Lists.newArrayList("00:00","23:59","00:00")));
        // System.out.println(solution.findMinDifference(Lists.newArrayList("00:00","23:59")));
        //System.out.println(solution.findMinDifference(Lists.newArrayList("00:00","04:00","22:00")));
        System.out.println(solution.findMinDifference(Lists.newArrayList("05:31","22:08","00:35")));
    }
}

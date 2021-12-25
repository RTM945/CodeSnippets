package algorithms.leetcode._1514_day_of_the_year;

public class _1514 {
    // 给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。
    // 请你计算并返回该日期是当年的第几天。

    // 通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。
    // 每个月的天数与现行公元纪年法（格里高利历）一致。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/day-of-the-year
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public int dayOfYear(String date) {
            String[] strs = date.split("-");
            int year = Integer.parseInt(strs[0]);
            int month = Integer.parseInt(strs[1]);
            int day = Integer.parseInt(strs[2]);
            int[] monthDays = new int[]{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
            // 闰年
            boolean flag = year % 100 == 0 ? year % 400 == 0 : year % 4 == 0;
            if (flag) {
                monthDays[1] = 29;
            }
            int res = 0;
            for (int i = 0; i < month - 1; i++) {
                res += monthDays[i];
            }
            return res + day;
        }
    }

    public static void main(String[] args) {
        _1514 _1514 = new _1514();
        Solution solution = _1514.new Solution();
        System.out.println(solution.dayOfYear("2019-01-09"));  //9
        System.out.println(solution.dayOfYear("2019-02-10"));  //41
        System.out.println(solution.dayOfYear("2003-03-01"));  //60
        System.out.println(solution.dayOfYear("2004-03-01"));  //61
    }
}

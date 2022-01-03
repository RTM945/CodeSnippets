package algorithms.leetcode._1185_day_of_the_week;

public class _1185 {
    // 给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。

    // 输入为三个整数：day、month 和 year，分别表示日、月、年。

    // 您返回的结果必须是这几个值中的一个 
    // {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
    // "Saturday"}。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/day-of-the-week
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public String dayOfTheWeek(int day, int month, int year) {
            String[] week = { "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" };
            // 蔡勒公式
            int[] ints = new int[] { 0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4 };
            year -= month < 3 ? 1 : 0;
            return week[(year + year / 4 - year / 100 + year / 400 + ints[month - 1] + day) % 7];
        }
    }

    public static void main(String[] args) {
        _1185 _1185 = new _1185();
        Solution solution = _1185.new Solution();
        System.out.println(solution.dayOfTheWeek(3, 1, 2022));
    }
}

package algorithms.leetcode._682_baseball_game;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class _682 {
    // 你现在是一场采用特殊赛制棒球比赛的记录员。这场比赛由若干回合组成，
    // 过去几回合的得分可能会影响以后几回合的得分。

    // 比赛开始时，记录是空白的。你会得到一个记录操作的字符串列表 ops，其中 ops[i] 
    // 是你需要记录的第 i 项操作，ops 遵循下述规则：

    // 整数 x - 表示本回合新获得分数 x
    // "+" - 表示本回合新获得的得分是前两次得分的总和。
    // 题目数据保证记录此操作时前面总是存在两个有效的分数。
    // "D" - 表示本回合新获得的得分是前一次得分的两倍。
    // 题目数据保证记录此操作时前面总是存在一个有效的分数。
    // "C" - 表示前一次得分无效，将其从记录中移除。
    // 题目数据保证记录此操作时前面总是存在一个有效的分数。
    // 请你返回记录中所有得分的总和。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/baseball-game
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    class Solution {
        public int calPoints(String[] ops) {
            LinkedList<Integer> list = new LinkedList<>();
            int idx = 0;
            for (String s : ops) {
                if (s.equals("+")) {
                    list.add(list.get(idx - 1) + list.get(idx - 2));
                    idx++;
                } else if (s.equals("D")) {
                    list.add(list.get(idx - 1) * 2);
                    idx++;
                } else if (s.equals("C")) {
                    list.removeLast();
                    idx--;
                } else {
                    list.add(Integer.parseInt(s));
                    idx++;
                }
            }
            int res = 0;
            for (int i : list) {
                res += i;
            }
            return res;
        }
    }

    class Solution1 {
        public int calPoints(String[] ops) {
            int ret = 0;
            List<Integer> points = new ArrayList<Integer>();
            for (String op : ops) {
                int n = points.size();
                switch (op.charAt(0)) {
                    case '+':
                        ret += points.get(n - 1) + points.get(n - 2);
                        points.add(points.get(n - 1) + points.get(n - 2));
                        break;
                    case 'D':
                        ret += 2 * points.get(n - 1);
                        points.add(2 * points.get(n - 1));
                        break;
                    case 'C':
                        ret -= points.get(n - 1);
                        points.remove(n - 1);
                        break;
                    default:
                        ret += Integer.parseInt(op);
                        points.add(Integer.parseInt(op));
                        break;
                }
            }
            return ret;
        }
    }

    public static void main(String[] args) {
        _682 _682 = new _682();
        Solution solution = _682.new Solution();
        System.out.println(solution.calPoints(new String[]{"5","2","C","D","+"}));
    }
}

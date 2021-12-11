package algorithms.leetcode._911_online_election;

import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;
import java.util.Map.Entry;

public class _911 {
    // 给你两个整数数组 persons 和 times 。
    // 在选举中，第 i 张票是在时刻为 times[i] 时投给候选人 persons[i] 的。

    // 对于发生在时刻 t 的每个查询，需要找出在 t 时刻在选举中领先的候选人的编号。

    // 在 t 时刻投出的选票也将被计入我们的查询之中。
    // 在平局的情况下，最近获得投票的候选人将会获胜。

    // 实现 TopVotedCandidate 类：

    // TopVotedCandidate(int[] persons, int[] times) 使用 persons 和 times 数组初始化对象。
    // int q(int t) 根据前面描述的规则，返回在时刻 t 在选举中领先的候选人的编号。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/online-election
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // [0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]
    // 第0张票在time[0]时投给person[0]
    class TopVotedCandidate {

        // 人和票数
        private Map<Integer, Integer> person = new HashMap<>();
        // 每个time得票最多的
        private TreeMap<Integer, Integer> time2map = new TreeMap<>();

        public TopVotedCandidate(int[] persons, int[] times) {
            int maxPersion = -1;
            for (int i = 0; i < times.length; i++) {
                int count = person.getOrDefault(persons[i], 0) + 1;
                person.put(persons[i], count);
                int max = person.getOrDefault(maxPersion, 0);
                if (count >= max) {
                    maxPersion = persons[i];
                }
                time2map.put(times[i], maxPersion);
            }
        }
        
        public int q(int t) {
            return time2map.floorEntry(t).getValue();
        }
    }

    public static void main(String[] args) {
        _911 _911 = new _911();
        TopVotedCandidate tvc = _911.new TopVotedCandidate(new int[]{0, 1, 1, 0, 0, 1, 0}, new int[]{0, 5, 10, 15, 20, 25, 30});
        System.out.println(tvc.q(3)); //0
        System.out.println(tvc.q(12)); //1
    }
    
    /**
     * Your TopVotedCandidate object will be instantiated and called as such:
     * TopVotedCandidate obj = new TopVotedCandidate(persons, times);
     * int param_1 = obj.q(t);
     */
}

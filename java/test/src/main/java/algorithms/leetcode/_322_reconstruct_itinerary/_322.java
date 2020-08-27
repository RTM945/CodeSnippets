package algorithms.leetcode._322_reconstruct_itinerary;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;

/* 给定一个机票的字符串二维数组 [from, to]，
子数组中的两个成员分别表示飞机出发和降落的机场地点，
对该行程进行重新规划排序。
所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
说明:
如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
所有的机场都用三个大写字母表示（机场代码）。
假定所有机票至少存在一种合理的行程。
示例 1:
输入: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出: ["JFK", "MUC", "LHR", "SFO", "SJC"]
示例 2:
输入: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
解释: 另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reconstruct-itinerary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _322 {
    class Solution {
        // 看上去是拓扑排序
        // 先找起点，就是JFK开始且终点最小的那个
        public List<String> findItinerary(List<List<String>> tickets) {
            String start = "JFK";
            List<String> res = new ArrayList<>();
            res.add(start);
            int remove = -1;
            while (!tickets.isEmpty()) {
                String end = "ZZZ";
                for (int i = 0; i < tickets.size(); i++) {
                    if (tickets.get(i).get(0).equals(start)) {
                        if (less(tickets.get(i).get(1), end)) {
                            remove = i;
                            end = tickets.get(i).get(1);
                        }
                    }
                }
                if (end.equals("zzz")) {

                }
                res.add(end);
                start = end;
                tickets.remove(remove);
            }
            return res;
        }

        boolean less(String a, String b) {
            if (a.equals(b)) {
                return false;
            }
            if (a.length() < b.length()) {
                return true;
            }
            for (int i = 0; i < a.length(); i++) {
                if (a.charAt(i) == b.charAt(i)) {
                    continue;
                } else {
                    return a.charAt(i) < b.charAt(i);
                }
            }
            return false;
        }

    }

    class Solution1 {
        // [["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]翻车
        // 做邻接表。再dfs
        // Hierholzer 算法
        Map<String, PriorityQueue<String>> map = new HashMap<String, PriorityQueue<String>>();
        List<String> itinerary = new LinkedList<String>();

        public List<String> findItinerary(List<List<String>> tickets) {
            for (List<String> ticket : tickets) {
                String src = ticket.get(0), dst = ticket.get(1);
                if (!map.containsKey(src)) {
                    map.put(src, new PriorityQueue<String>());
                }
                map.get(src).offer(dst);
            }
            dfs("JFK");
            Collections.reverse(itinerary);
            return itinerary;
        }

        public void dfs(String curr) {
            while (map.containsKey(curr) && map.get(curr).size() > 0) {
                String tmp = map.get(curr).poll();
                dfs(tmp);
            }
            itinerary.add(curr);
        }
    }

    public static void main(String[] args) {
        _322 q = new _322();
        List<List<String>> tickets = new ArrayList<>();
        // [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
        tickets.add(Arrays.asList("JFK", "SFO"));
        tickets.add(Arrays.asList("JFK", "ATL"));
        tickets.add(Arrays.asList("SFO", "ATL"));
        tickets.add(Arrays.asList("ATL", "JFK"));
        tickets.add(Arrays.asList("ATL", "SFO"));
        System.out.println(q.new Solution().findItinerary(tickets));
    }
}
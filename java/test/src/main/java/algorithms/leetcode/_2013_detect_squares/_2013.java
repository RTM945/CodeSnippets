package algorithms.leetcode._2013_detect_squares;

import java.util.HashMap;
import java.util.Map;

public class _2013 {
    // 给你一个在 X-Y 平面上的点构成的数据流。设计一个满足下述要求的算法：

    // 添加 一个在数据流中的新点到某个数据结构中。可以添加 重复 的点，并会视作不同的点进行处理。
    // 给你一个查询点，请你从数据结构中选出三个点，使这三个点和查询点一同构成一个 面积为正 的 轴对齐正方形 ，
    // 统计 满足该要求的方案数目。
    // 轴对齐正方形 是一个正方形，除四条边长度相同外，还满足每条边都与 x-轴 或 y-轴 平行或垂直。

    // 实现 DetectSquares 类：

    // DetectSquares() 使用空数据结构初始化对象
    // void add(int[] point) 向数据结构添加一个新的点 point = [x, y]
    // int count(int[] point) 统计按上述方式与点 point = [x, y] 共同构造 轴对齐正方形 的方案数。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/detect-squares
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class DetectSquares {
        Map<Integer, Map<Integer, Integer>> cnt;
    
        public DetectSquares() {
            cnt = new HashMap<>();
        }
    
        public void add(int[] point) {
            int x = point[0], y = point[1];
            Map<Integer, Integer> yCnt = cnt.computeIfAbsent(y, k -> new HashMap<>());
            yCnt.put(x, yCnt.getOrDefault(x, 0) + 1);
        }
    
        public int count(int[] point) {
            int res = 0;
            int x = point[0], y = point[1];
            if (!cnt.containsKey(y)) {
                return 0;
            }
            Map<Integer, Integer> yCnt = cnt.get(y);
            
            for (Map.Entry<Integer, Map<Integer, Integer>> entry : cnt.entrySet()) {
                // 另一个y
                int ny = entry.getKey();
                Map<Integer, Integer> nyCnt = entry.getValue();
                if (ny != y) {
                    // 用ny 和 (x, y) 构建正方形 边长为d
                    int d = ny - y;
                    // (x, y + d) * (x + d, y) * (x + d, y + d)
                    res += nyCnt.getOrDefault(x, 0) * yCnt.getOrDefault(x + d, 0) * nyCnt.getOrDefault(x + d, 0);
                    // (x, y + d) * (x - d, y) * (x - d, y + d)
                    res += nyCnt.getOrDefault(x, 0) * yCnt.getOrDefault(x - d, 0) * nyCnt.getOrDefault(x - d, 0);
                }
            }
            return res;
        }
    }
    
    
    /**
     * Your DetectSquares object will be instantiated and called as such:
     * DetectSquares obj = new DetectSquares();
     * obj.add(point);
     * int param_2 = obj.count(point);
     */
}

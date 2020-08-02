package algorithms.leetcode._632_smallest_range_covering_elements_from_k_lists;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.PriorityQueue;

// https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists
// 你有 k 个升序排列的整数数组。
// 找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
// 我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。
// 示例 1:
// 输入:[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
// 输出: [20,24]
// 解释: 
// 列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
// 列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
// 列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
// 注意:
// 给定的列表可能包含重复元素，所以在这里升序表示 >= 。
// 1 <= k <= 3500
// -105 <= 元素的值 <= 105
// 对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。
public class _632 {
    class Solution {
        //hard 害怕
        //根据大小的定义倒是很容易写出比较大小的逻辑
        //最小区间使得 k 个列表中的每个列表至少有一个数包含在其中
        //最小区间的边界应该是每个列表中都有的一个值
        //打算先找出来 然后从小到大排序
        //这样第一个元素和最后一个元素构成区间
        //就没有比较的事情了
        //好奇怪哦 估计想错了
        //一开始想使用hashmap保存每个元素的个数, 这样个数=list数时说明每个集合都有
        //但题目说数字可重复
        //又或者使用第一个集合做set, 遍历其他集合,使用map标记出其他集合也有的元素
        //性能估计爆炸
        
        //答案首先使用堆
        //将题目转化为, 从k个集合中各取一个数, 使k个数的最大值和最小值之差最小
        //无疑是对的, 但我就想不到这个转化...
        //假设k中的最小值是第i个集合中的x
        //那么对于另一个集合j,被选入k的元素y, 应该是j中大于x的最小的数
        //证明: 假设j中的另一个元素z > y, 则z - x > y - x, 结果还是取了y
        //构造一个对象, 当前集合的下标, 当前遍历到的值
        //使用小顶堆维护这个对象, 并维护堆中的最大值
        //遍历每个堆寻找边界
        //意思就是, 因为每个集合都是递增的
        //所有集合中第一个元素的最小值同时也是所有元素的最小值
        //他和所有集合中第一个元素的最大值构成的区间
        //绝对是符合"每个集合至少有个一个在这个区间中"的条件
        //新的左边界就应该是
        class Pointer {
            int i;
            int val;
        }

        public int[] smallestRange(List<List<Integer>> nums) {
            int left = 0;
            int right = Integer.MAX_VALUE;
            int max = Integer.MIN_VALUE;
            int size = nums.size();
            int[] next = new int[size];
            int minIndex = 0; //list中第一个元素是最小的
            PriorityQueue<Pointer> heap = new PriorityQueue<>((p1, p2) -> p1.val - p2.val);
            for (int i = 0; i < size; i++) {
                Pointer ptr = new Pointer();
                ptr.i = i;
                ptr.val = nums.get(i).get(minIndex);
                heap.offer(ptr);
                max = Math.max(max, ptr.val);
            }
            while(!heap.isEmpty()) {
                Pointer ptr = heap.poll();
                if(max - ptr.val < right - left) { 
                    right = max;
                    left = ptr.val;
                }
                next[ptr.i]++; //该集合中已经判断过的数量,也可用于下一个需要取的元素下标
                if (next[ptr.i] == nums.get(ptr.i).size()) {
                    break;
                }
                ptr.val = nums.get(ptr.i).get(next[ptr.i]);
                heap.offer(ptr);
                max = Math.max(max, ptr.val);
            }

            return new int[]{left, right};
        }

    }

    public static void main(String[] args) {
        _632 q = new _632();
        List<List<Integer>> list = new ArrayList<>();
        list.add(Arrays.asList(4, 10 ,15, 24, 26));
        list.add(Arrays.asList(0, 9 ,12, 20));
        list.add(Arrays.asList(5, 18, 22, 30));
        q.new Solution().smallestRange(list);
    }
}
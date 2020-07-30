package algorithms.leetcode._046_permutations;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
//https://leetcode-cn.com/problems/permutations
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
// 示例:
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]
public class _046 {
    class Solution {
        //据说是回溯法的经典题目
        //用nums构建决策树，以[1,2,3]为例
        //      root
        //  1     2     3
        // 2 3   1 3   1 2
        // 3 2   3 1   2 1
        //结果就是从root开始dfs到叶子
        //因为不知道有多少个叶子，只能从回退到叶子结点的父结点找其他分支
        //这就是回溯
        List<List<Integer>> res = new ArrayList<>();
        public List<List<Integer>> permute(int[] nums) {
            LinkedList<Integer> list = new LinkedList<>();
            permute(nums, list);
            return res;
        }

        void permute(int[] nums, LinkedList<Integer> list) {
            if(list.size() == nums.length){
                //已经到了叶子结点
                res.add(new LinkedList<>(list));
                return;
            }
            for (int i = 0; i < nums.length; i++) {
                if(list.contains(nums[i])) {
                    continue; //选择nums[i],先确保list中不存在nums[i]
                }
                list.add(nums[i]);
                permute(nums, list); //进入下一层决策树
                list.removeLast(); //回溯，最其他选择
            }
        }
    }

    class Solution1 {
        //显式的dfs回溯
        public List<List<Integer>> permute(int[] nums) {
            int len = nums.length;
            List<List<Integer>> res = new ArrayList<>();
            if(len == 0) {
                return res;
            }
            LinkedList<Integer> path = new LinkedList<>();
            //其实也可以用path.contains(nums[i])来判断是否使用过
            //空间换时间
            boolean[] used = new boolean[len]; 
            dfs(nums, path, res, used, 0);
            return res;
        }

        private void dfs(int[] nums, LinkedList<Integer> path, List<List<Integer>> res, boolean[] used, int depth) {
            if(nums.length == path.size()){
                res.add(new LinkedList<>(path));
                return;
            }
            for (int i = 0; i < nums.length; i++) {
                if(!used[i]) {
                    path.add(nums[i]);
                    used[i] = true;
                    dfs(nums, path, res, used, depth++);
                    //回溯
                    used[i] = false;
                    path.removeLast();
                }
            }
        }
    }

    public static void main(String[] args) {
        _046 q = new _046();
        System.out.println(q.new Solution().permute(new int[]{1, 2, 3}));
        System.out.println(q.new Solution1().permute(new int[]{1, 2, 3}));
    }
}
package algorithms.leetcode.interview_16_11_diving_board_lcci;

import java.util.*;
import java.util.stream.Collectors;

//你正在使用一堆木板建造跳水板。
//有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。
//你必须正好使用k块木板。
//编写一个方法，生成跳水板所有可能的长度。
//返回的长度需要从小到大排列。
//示例：
//输入：
//shorter = 1
//longer = 2
//k = 3
//输出： {3,4,5,6}
//提示：
//0 < shorter <= longer
//0 <= k <= 100000
public class _16_11 {
    //乍一看不知道啥意思，结合输入输出想明白了
    //求shorter和longer在总数为k的情况下每种组合的和
    //shorter和longer可以重复
    //一瞬间的想法是构造一个k + 1层的满二叉树，每个结点的左子结点是shorter，右子节点是longer
    //遍历出所有路径存入set去重得到结果
    //最近玩树玩的有点智障
    //应该有更简单的做法
    //那么设x个shorter，longer的个数就是k-x
    //那么可能的长度就是x*shorter + (k-x)*longer
    //结果排序一下
    //那么结果可不可以不用排序呢？
    //如果从0到k遍历，使用shorter的数量从k开始
    //后一个结果一定大于前一个结果的
    //方法要求返回int[]在java里面比较麻烦
    //因为不知道结果有多少个
    //只能将结果转int[]了
    //但又要过滤重复
    //只能用TreeSet了
    class Solution {
        public int[] divingBoard(int shorter, int longer, int k) {
            if(k == 0) {
                return new int[]{};
            }
            TreeSet<Integer> set = new TreeSet<>();
            for (int i = 0; i <= k; i++) {
                int sum = shorter * i + (k - i) * longer;
                set.add(sum);
            }
            int[] res = new int[set.size()];
            for (int i = 0; i < res.length; i++) {
                res[i] = set.pollFirst();
            }
            return res;
        }
    }

    class Solution1 {
        //上面的方法74 ms, 在所有 Java 提交中击败了5.01%的用户
        //到底有没有重复的情况
        //存不存在
        //as + (k-a)l = bs + (k-b)l a<>b
        //解不出来
        //其实不存在结果重复的情况
        //所以没必要上set
        //并且总数也知道，就是k+1
        //为什么呢，用数学语言就能一眼看明白
        //总数={(a, k-a)| 0 <= a <= k}
        //现在只需要注意顺序
        public int[] divingBoard(int shorter, int longer, int k) {
            if(k == 0) {
                return new int[]{};
            }
            if(shorter == longer) {
                return new int[]{shorter * k};
            }
            int[] res = new int[k + 1];
            for (int i = 0; i <= k; i++) {
                res[i] = shorter * (k - i) + longer * i;
            }
            return res;
        }
    }

    public static void main(String[] args) {
        System.out.println(Arrays.toString(new _16_11().new Solution().divingBoard(1, 2, 3)));
        System.out.println(Arrays.toString(new _16_11().new Solution1().divingBoard(1, 2, 3)));
    }
}

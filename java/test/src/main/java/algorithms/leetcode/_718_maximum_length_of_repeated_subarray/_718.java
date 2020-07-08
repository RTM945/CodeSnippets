package algorithms.leetcode._718_maximum_length_of_repeated_subarray;

public class _718 {
    //https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
    //给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
    //示例 1:
    //输入:
    //A: [1,2,3,2,1]
    //B: [3,2,1,4,7]
    //输出: 3
    //解释:
    //长度最长的公共子数组是 [3, 2, 1]。
    //说明:
    //1 <= len(A), len(B) <= 1000
    //0 <= A[i], B[i] < 100

    //仔细看题目得知不需要排序，需要数值和顺序相同
    //一开始思路是枚举出所有子数组，存入Map判断重复
    //性能一定很差
    //如果不借用Map更暴力一点，就是循环套循环判断起始数字是否相同，再判断接下去的元素是否相同
    class Solution {
        public int findLength(int[] A, int[] B) {
            int maxSubLen = 0;
            for (int i = 0; i < A.length; i++) {
                //一步可以优化的操作
                //如果后面的元素个数小于等于maxSubLen，就可以不用再遍历了
                if(A.length - i <= maxSubLen) {
                    break;
                }
                for (int j = 0; j < B.length; j++) {
                    if(B.length - j <= maxSubLen) {
                        break;
                    }
                    if (A[i] == B[j]) {
                        //至少一个相同
                        int subLen = 1;
                        //从此开始向后比较
                        int ai = i + 1;
                        int bi = j + 1;
                        while (ai < A.length &&
                                bi < B.length && //越界判断
                                A[ai] == B[bi]) { //后面的元素是否相同
                            subLen++;
                            ai++;
                            bi++;
                        }
                        //有将subLen作为i和j增量的做法，可以不需要ai和bi变量
                        /*while (i + subLen < A.length &&
                                j + subLen < B.length &&
                                A[i + subLen] == B[j + subLen]) {
                            subLen++;
                        }*/
                        //比较subLen
                        if (subLen > maxSubLen) {
                            maxSubLen = subLen;
                        }
                    }
                }
            }
            return maxSubLen;
        }
    }

    class Solution2 {
        //于是就有了动态规划做法
        //https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zhe-yao-jie-shi-ken-ding-jiu-dong-liao-by-hyj8/
        //思路：
        //如果有最长子序列，则该序列为相同的前缀序列+相同的最后一位
        //相同的前缀序列=相同的前几位+相同的最后一位
        //递归到相同的第一位
        //从二维数组角度来看，将相同的第一位dp[i][j]设置为1
        //如果后面的序列仍然相同，dp[i+1][j+1] = dp[i][j] + 1
        //出于方便编码，起始行和列各增加1作为辅助
        //dp[i][j] = dp[i - 1][j - 1] + 1
        public int findLength(int[] A, int[] B) {
            //按套路建二维数组
            int[][] dp = new int[A.length + 1][B.length + 1];
            int max = 0;
            //找到相同的那一位开始
            //边界条件i和j是作为二维数组下标时范围是[1, A.length]
            //作为AB下标时范围是[0, A.length)即上面的范围-1
            //这样二维数组可以忽略起始用于辅助的一行一列
            //AB的遍历也不会越界
            for (int i = 1; i <= A.length; i++) {
                for (int j = 1; j <= B.length; j++) {
                    if(A[i - 1] == B[j - 1]) {
                        dp[i][j] = dp[i - 1][j - 1] + 1;
                    }
                    //比较子序列长度
                    if(max < dp[i][j]) {
                        max = dp[i][j];
                    }
                }
            }
            //也有int[][] dp = new int[A.length][B.length];的方法
            //需要在循环中进行判断
            //但不能让i-1和j-1小于0
            //如果i或j=0，且A[i] == B[j]，可以直接将dp[i][j] = 1
            /*int[][] dp = new int[A.length][B.length];
            for (int i = 0; i < A.length; i++) {
                for (int j = 0; j < B.length; j++) {
                    if(A[i] == B[j]) {
                        if (i > 0 && j > 0) {
                            dp[i][j] = dp[i-1][j-1]+1;
                        } else {
                            dp[i][j] = 1;
                        }
                    }
                    //比较子序列长度
                    if(max < dp[i][j]) {
                        max = dp[i][j];
                    }
                }
            }*/
            return max;
        }
    }

    class Solution3{
        //滑动窗口
        //https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zhe-yao-jie-shi-ken-ding-jiu-dong-liao-by-hyj8/
        //想象A在上，B在下，B的最后一位与A的第一位对齐
        //A位置不变，一位一位的向后滑动B数组，可以滑到A的最后一位与B的第一位对齐
        //上下对齐的数字即为窗口，比较窗口中的上下数字，得出连续相等的最大长度
        //可以分为两步：
        //A固定，B向后滑
        //B固定，A向前滑
        //https://segmentfault.com/a/1190000023069410
        public int findLength(int[] A, int[] B) {
            int max = 0;
            //B第一位对齐A第一位，B向后滑动
            for (int i = 0; i < A.length; i++) {
                //窗口的大小是B的长度与A.length - i的最小值
                int window = A.length - i;
                if (B.length < window) {
                    window = B.length;
                }
                int thisMax = maxLength(A, B, i, 0, window);
                if (thisMax > max) {
                    max = thisMax;
                }
            }
            //A第一位对齐B第一位，A向后滑动
            for (int i = 0; i < B.length; i++) {
                int window = B.length - i;
                if(A.length < window) {
                    window = A.length;
                }
                int thisMax = maxLength(A, B, 0, i, window);
                if (thisMax > max) {
                    max = thisMax;
                }
            }
            return max;
        }

        public int maxLength(int[] A, int[] B, int aStart, int bStart, int window) {
            int max = 0, len = 0;
            for (int i = 0; i < window; i++) {
                if (A[aStart + i] == B[bStart + i]) {
                    len++; //上下相同
                } else {
                    len = 0; //上下不同，长度归零
                }
                if(len > max) {
                    max = len;
                }
            }
            return max;
        }
    }

    public static void main(String[] args) {
        int[] A = {1,2,3,2,1};
        int[] B = {3,2,1,4,7};
        _718 tmp = new _718();
        System.out.println(tmp.new Solution().findLength(A, B));
        System.out.println(tmp.new Solution2().findLength(A, B));
        System.out.println(tmp.new Solution3().findLength(A, B));
    }
}

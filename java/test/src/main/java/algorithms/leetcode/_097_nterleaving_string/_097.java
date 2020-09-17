package algorithms.leetcode._097_nterleaving_string;

//https://leetcode-cn.com/problems/interleaving-string/
//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
//示例 1:
//输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出: true
//示例 2:
//输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出: false
public class _097 {
    //开始每天期待leetcode给我推到底是简单还是困难还是中等的题了
    //说到交错，首先要定义交错
    //示例2中，虽然s3也是由s1和s2组成的
    //但不算交错
    //在示例1中，s3 = s1[0:2] + s2[0:2] + s1[2:4] + s2[2:4] + s2[4] + s1[4]
    //也就是说每次交替使用的s1和s2位数相等
    //而且顺序还可以乱
    //现在有了s1 s2 s3，怎么验证s3是交错呢
    //暴力解法应该是强行判断s3头部有几位s1的元素，再把数字套到s2上判断，以此类推
    //写起来好麻烦
    //转换思路
    //继续研究s2为什么是false，猜想重要的不是位数，而是顺序
    //于是想把s3从头开始挨个儿remove掉s1 s2能符合的最长字串，都无法remove时，返回false
    //还是好麻烦
    //算了看答案
    //和https://leetcode-cn.com/problems/unique-paths/ 相同的思路 真的天才...
    //将s1和s2矩阵排列，假设竖s1横s2，开头多一行一列辅助
    //s3则是从(0, 0)开始，向右取s2或向下取s1组成的
    //定义bool[i][j]表示s1前i个和s2前j个是否能构成通路
    //dp[0][0] = true
    //i=0时 dp[0][j] = s3[0:j - 1] == s2[0:j]
    //j=0时 dp[i][0] = s3[0:i] == s1[0:i]
    //dp[i][j] = (dp[i-1][j] && s3[i+j-1] == s1[i-1]) || (dp[i][j-1]) && s3[i+j-1] == s2[j-1])
    class Solution {
        public boolean isInterleave(String s1, String s2, String s3) {
            int len1 = s1.length();
            int len2 = s2.length();
            if (s3.length() != len1 + len2) {
                return false;
            }
            boolean[][] dp = new boolean[len1 + 1][len2 + 1];
            dp[0][0] = true;
            for (int i = 1; i <= len1 && s1.charAt(i - 1) == s3.charAt(i - 1); i++) {
                dp[i][0] = true;
            }
            for (int j = 1; j <= len2 && s2.charAt(j - 1) == s3.charAt(j - 1); j++) {
                dp[0][j] = true;
            }
            for (int i = 1; i <= len1; i++) {
                for (int j = 1; j <= len2; j++) {
                    dp[i][j] = (dp[i-1][j] && s3.charAt(i+j-1) == s1.charAt(i-1)
                            || (dp[i][j-1]) && s3.charAt(i+j-1) == s2.charAt(j-1));
                }
            }
            return dp[len1][len2];
        }
    }

    class TestCase{
        String s1;
        String s2;
        String s3;
        boolean want;

        public TestCase(String s1, String s2, String s3, boolean want) {
            this.s1 = s1;
            this.s2 = s2;
            this.s3 = s3;
            this.want = want;
        }
    }

    public static void main(String[] args) {
        _097 q = new _097();
        TestCase t = q.new TestCase("aabcc", "dbbca", "aadbbcbcac", true);
        System.out.println(q.new Solution().isInterleave(t.s1, t.s2, t.s3));
    }
}

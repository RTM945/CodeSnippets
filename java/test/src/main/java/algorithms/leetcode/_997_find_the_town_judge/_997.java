package algorithms.leetcode._997_find_the_town_judge;

public class _997 {
    // 在一个小镇里，按从 1 到 n 为 n 个人进行编号。传言称，这些人中有一个是小镇上的秘密法官。

    // 如果小镇的法官真的存在，那么：

    // 1.小镇的法官不相信任何人。
    // 2.每个人（除了小镇法官外）都信任小镇的法官。
    // 3.只有一个人同时满足条件 1 和条件 2 。
    // 给定数组 trust，该数组由信任对 trust[i] = [a, b] 组成，
    // 表示编号为 a 的人信任编号为 b 的人。

    // 如果小镇存在秘密法官并且可以确定他的身份，请返回该法官的编号。否则，返回 -1。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/find-the-town-judge
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 有向图 法官只有出度没有入度 所以入度是n - 1 出度是0
    class Solution {
        public int findJudge(int n, int[][] trust) {
            int[] in = new int[n + 1];
            int[] out = new int[n + 1];
            for(int[] arr : trust) {
                int a = arr[0];
                int b = arr[1];
                in[b]++;
                out[a]++;
            }
            for (int i = 1; i <= n; i++) {
                if (in[i] == n - 1 && out[i] == 0) {
                    return i;
                }
            }
            return -1;
        }
    }
}

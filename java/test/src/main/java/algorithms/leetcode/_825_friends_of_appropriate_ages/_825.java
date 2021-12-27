package algorithms.leetcode._825_friends_of_appropriate_ages;

import java.util.Arrays;

public class _825 {
    // 在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。

    // 如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：

    // age[y] <= 0.5 * age[x] + 7
    // age[y] > age[x]
    // age[y] > 100 && age[x] < 100
    // 否则，x 将会向 y 发送一条好友请求。

    // 注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。

    // 返回在该社交媒体网站上产生的好友请求总数。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/friends-of-appropriate-ages
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 不带脑子暴力一个
    // 超时了
    class Solution {
        public int numFriendRequests(int[] ages) {
            int res = 0;
            for (int x = 0; x < ages.length; x++) {
                for (int y = 0; y < ages.length; y++) {
                    if (x == y || ages[y] <= 0.5 * ages[x] + 7 || ages[y] > ages[x] || (ages[y] > 100 && ages[x] < 100)) {
                        continue;
                    }
                    res++;
                }
            }
            return res;
        }
    }

    // 带上脑子
    // age[y] <= 0.5 * age[x] + 7
    // age[y] > age[x]
    // age[y] > 100 && age[x] < 100 -> age[y] > 100 > age[x] => age[y] > age[x]
    // 得到
    // age[x] < age[y] <= 0.5 * age[x] + 7 符合这个条件的要排除
    // 问题变成了在 age 从小到大区间中找每个age对应的(0.5 * age[x] + 7, age[x]]范围
    // 还能算出来age > 14
    // 双指针
    class Solution1 {
        public int numFriendRequests(int[] ages) {
            Arrays.sort(ages);
            int left = 0;
            int right = 0;
            int res = 0;
            for (int i = 0; i < ages.length; i++) {
                // 小于15岁不配有朋友
                if (ages[i] < 15) {
                    continue;
                }
                // 
                while (ages[left] <= 0.5 * ages[i] + 7) {
                    left++;
                }
                while (right + 1 < ages.length && ages[right + 1] <= ages[i]) {
                    right++;
                }
                res += right - left;
            }
            return res;
        }
    }

    // 有时候提示也能得到题解
    // 提示：
    // n == ages.length
    // 1 <= n <= 2 * 104
    // 1 <= ages[i] <= 120
    // 用一个数组记录每个年龄的数量
    // 前缀和
    class Solution2 {
        public int numFriendRequests(int[] ages) {
            int[] cnt = new int[121];
            for (int age : ages) {
                ++cnt[age];
            }
            int[] pre = new int[121];
            for (int i = 1; i <= 120; ++i) {
                // 本年龄与小于本年龄的和
                pre[i] = pre[i - 1] + cnt[i];
            }
            int ans = 0;
            for (int i = 15; i <= 120; ++i) {
                if (cnt[i] > 0) {
                    int bound = (int) (i * 0.5 + 8);
                    ans += cnt[i] * (pre[i] - pre[bound - 1] - 1);
                }
            }
            return ans;
        }
    }
}

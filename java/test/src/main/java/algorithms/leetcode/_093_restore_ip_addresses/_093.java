package algorithms.leetcode._093_restore_ip_addresses;

import java.util.ArrayList;
import java.util.List;

// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
// 有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。
// 示例:
// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/restore-ip-addresses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
public class _093 {
    class Solution {
        // 先暴力
        public List<String> restoreIpAddresses(String s) {
            List<String> list = new ArrayList<>();
            if (s.length() > 12) {
                return list;
            }
            StringBuilder ip = new StringBuilder();
            for (int a = 1; a < 4; a++) {
                for (int b = 1; b < 4; b++) {
                    for (int c = 1; c < 4; c++) {
                        for (int d = 1; d < 4; d++) {
                            if (a + b + c + d == s.length()) {
                                int n1 = Integer.parseInt(s.substring(0, a));
                                int n2 = Integer.parseInt(s.substring(a, a + b));
                                int n3 = Integer.parseInt(s.substring(a + b, a + b + c));
                                int n4 = Integer.parseInt(s.substring(a + b + c));
                                if (n1 <= 255 && n2 <= 255 && n3 <= 255 && n4 <= 255) {
                                    ip.append(n1).append('.').append(n2).append('.').append(n3).append('.').append(n4);
                                    if (ip.length() == s.length() + 3) {
                                        list.add(ip.toString());
                                    }
                                    ip.delete(0, ip.length());
                                }
                            }
                        }
                    }
                }
            }
            return list;
        }
    }

    //还有dfs回溯方法，懒狗，跳过了
}
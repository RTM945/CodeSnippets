package algorithms.leetcode._507_perfect_number;

import java.util.ArrayList;
import java.util.List;

public class _507 {
    // 对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。

    // 给定一个 整数 n， 如果是完美数，返回 true，否则返回 false

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/perfect-number
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 求这个数的所有约数再加起来
    class Solution {
        public boolean checkPerfectNumber(int num) {
            if (num == 1) {
                return false;
            }
    
            int sum = 1;
            for (int d = 2; d * d <= num; ++d) {
                if (num % d == 0) {
                    sum += d;
                    if (d * d < num) {
                        sum += num / d;
                    }
                }
            }
            return sum == num;
        }
    }

    // 根据欧几里得-欧拉定理
    // 题目范围 [1,10^8] 内的完全数 6,28,496,8128,33550336

    public static void main(String[] args) {
        _507 _507 = new _507();
        Solution solution = _507.new Solution();
        System.out.println(solution.checkPerfectNumber(28));
        System.out.println(solution.checkPerfectNumber(496));
        System.out.println(solution.checkPerfectNumber(8128));
        System.out.println(solution.checkPerfectNumber(2));
        System.out.println(solution.checkPerfectNumber(6));
    }
}

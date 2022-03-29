package algorithms.leetcode._2024_maximize_the_confusion_of_an_exam;

public class _2024 {
    // 一位老师正在出一场由 n 道判断题构成的考试，
    // 每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。
    // 老师想增加学生对自己做出答案的不确定性，方法是 最大化 有 连续相同 结果的题数。
    // （也就是连续出现 true 或者连续出现 false）。

    // 给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。
    // 除此以外，还给你一个整数 k ，表示你能进行以下操作的最多次数：

    // 每次操作中，将问题的正确答案改为 'T' 或者 'F' 
    // （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。
    // 请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 双指针？
    // 失败了
    class Solution {
        public int maxConsecutiveAnswers(String answerKey, int k) {
            int n = answerKey.length();
            int l = 0;
            int r = 1;
            int temp = k;
            int res = 0;
            for (int i = 0; i < n - 1; i++) {
                char c1 = answerKey.charAt(i);
                for (int j = i + 1; j < n; j++) {
                    char c2 = answerKey.charAt(j);
                    if (c1 == c2) {
                        r++;
                    } else {
                        if (temp > 0) {
                            temp--;
                            r++;
                        } else {
                            temp = k;
                            break;
                        }
                    }
                }
                res = Math.max(res, r - l);
                l++;
                r = l + 1;
            }
            return res;
        }
    }

    // 滑动窗口
    // 只记录窗口的大小用于答案，并非将窗口应用到字符串上
    class Solution1 {
        public int maxConsecutiveAnswers(String answerKey, int k) {
            return Math.max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'));
        }
    
        public int maxConsecutiveChar(String answerKey, int k, char ch) {
            int n = answerKey.length();
            int ans = 0;
            // sum为不相等的字符数量
            // right - left + 1 为窗口大小
            // 当sum < k时，可以消除掉，窗口不减小
            // 反之，找到不相等的字符，缩小窗口
            for (int left = 0, right = 0, sum = 0; right < n; right++) {
                sum += answerKey.charAt(right) != ch ? 1 : 0;
                while (sum > k) {
                    sum -= answerKey.charAt(left++) != ch ? 1 : 0;
                }
                ans = Math.max(ans, right - left + 1);
            }

            return ans;
        }
    }

    public static void main(String[] args) {
        _2024 _2024 = new _2024();
        Solution1 solution = _2024.new Solution1();
        System.out.println(solution.maxConsecutiveAnswers("TTFF", 2)); //4
        System.out.println(solution.maxConsecutiveAnswers("TFFT", 1)); //3
        System.out.println(solution.maxConsecutiveAnswers("TTFTTFTT", 1)); //5
    }
}

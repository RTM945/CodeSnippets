package algorithms.leetcode._043_multiply_strings;

import java.util.LinkedList;

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积
// 它们的乘积也表示为字符串形式。
// 示例 1:
// 输入: num1 = "2", num2 = "3"
// 输出: "6"
// 示例 2:
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"
// 说明：
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/multiply-strings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
public class _043 {
    class Solution {
        //可能又有数学
        //克制住看答案的冲动，想想吧
        //写了个非直接转换的string to long，但测试用例有超过long范围的
        //只能从字符串本身下手了
        public String multiply(String num1, String num2) {
            return s2l(num1) * s2l(num2) + "";
        }

        public long s2l(String str) {
            char[] cs = str.toCharArray();
            int n = str.length() - 1;
            long r = 0;
            for (int i = n; i >= 0; i--) {
                r += (cs[n - i] - '0') * Math.pow(10, i);
            }
            return r;
        }
    }

    class Solution1 {
        //想想手算乘法的过程，上面一个数字，下面用每一位去乘
        public String multiply(String num1, String num2) {
            if("0".equals(num1) || "0".equals(num2)) {
                return "0";
            }
            int a = num1.length() - 1;
            int b = num2.length() - 1;
            int s = 0; //进位
            StringBuilder sb = new StringBuilder('0');
            for (int i = 0; i <= a; i++) {
                int x = num1.charAt(a - i) - '0';
                LinkedList<Character> l = new LinkedList<>();
                for (int j = 0; j <= b; j++) {
                    int y = num2.charAt(b - j) - '0';
                    int z = x * y + s;
                    if(z > 9) {
                        s = z / 10;
                        z = z % 10;
                    }else{
                        s = 0;
                    }
                    //但是用数值存最终结果还是会溢出
                    l.addFirst((char)(z + '0'));
                }
                if(s != 0) {
                    l.addFirst((char)(s + '0'));
                    s = 0;
                }
                for (int j = 0; j < i; j++) {
                    l.addLast('0');
                }
                StringBuilder tmp = new StringBuilder();
                for(char c : l) {
                    tmp.append(c);
                }
                sb = addStrings(sb, tmp);
            }
            return sb.toString();
        }

        //leetcode 415
        public StringBuilder addStrings(StringBuilder num1, StringBuilder num2) {
            int i = num1.length() - 1;
            int j = num2.length() - 1;
            int f = 0; //进位
            StringBuilder sb = new StringBuilder();
            while(i >= 0 || j >= 0 || f != 0) {
                int a = i >= 0 ? num1.charAt(i) - '0' : 0;
                int b = j >= 0 ? num2.charAt(j) - '0' : 0;
                int c = a + b + f;
                if (c > 9) {
                    c = c % 10;
                    f = 1;
                }else{
                    f = 0;
                }
                sb.append(c);
                i--;
                j--;
            }
            sb.reverse();
            return sb;
        }
    }

    class Solution2 {
        //就知道有一堆数学的做法
        //或者说另一个巧方法吧
        //需要证明m * n的结果字符串长度最多为m + n
        /**
        num1的第i位(高位从0开始)和num2的第j位相乘的结果在乘积中的位置是[i+j, i+j+1]
        例: 123 * 45,  123的第1位 2 和45的第0位 4 乘积 08 存放在结果的第[1, 2]位中
          index:    0 1 2 3 4  
              
                        1 2 3
                    *     4 5
                    ---------
                          1 5
                        1 0
                      0 5
                    ---------
                      0 6 1 5
                        1 2
                      0 8
                    0 4
                    ---------
                    0 5 5 3 5
        这样我们就可以单独都对每一位进行相乘计算把结果存入相应的index中        
        **/
        public String multiply(String num1, String num2) {
            if (num1.equals("0") || num2.equals("0")) {
                return "0";
            }
            int m = num1.length(), n = num2.length();
            int[] ansArr = new int[m + n];
            for (int i = m - 1; i >= 0; i--) {
                int x = num1.charAt(i) - '0';
                for (int j = n - 1; j >= 0; j--) {
                    int y = num2.charAt(j) - '0';
                    ansArr[i + j + 1] += x * y;
                }
            }
            for (int i = m + n - 1; i > 0; i--) {
                ansArr[i - 1] += ansArr[i] / 10;
                ansArr[i] %= 10;
            }
            int index = ansArr[0] == 0 ? 1 : 0;
            StringBuffer ans = new StringBuffer();
            while (index < m + n) {
                ans.append(ansArr[index]);
                index++;
            }
            return ans.toString();
        }
    }

    public static void main(String[] args) {
        _043 q = new _043();
        System.out.println(q.new Solution().s2l("123"));
        System.out.println(q.new Solution1().multiply("123", "456")); //"56088"
    }
}
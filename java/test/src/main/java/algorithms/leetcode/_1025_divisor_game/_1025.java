package algorithms.leetcode._1025_divisor_game;

import java.util.HashMap;
import java.util.Map;

//https://leetcode-cn.com/problems/divisor-game/
//爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
//最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
//选出任一 x，满足 0 < x < N 且 N % x == 0 。
//用 N - x 替换黑板上的数字 N 。
//如果玩家无法执行这些操作，就会输掉游戏。
//只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。
//示例 1：
//输入：2
//输出：true
//解释：爱丽丝选择 1，鲍勃无法进行操作。
//示例 2：
//输入：3
//输出：false
//解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。
//提示：
//1 <= N <= 1000
public class _1025 {
    //名字好炫酷以为很难
    //没想到是个easy题
    //但细想也不是那么简单
    //只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。
    //意思是所有情况下爱丽丝都必须获胜，还是N个数中只需要爱丽丝获胜一次就行？
    //假设前者
    class Solution {
        int time;
        public boolean divisorGame(int N) {
            for (int x = 1; x < N; x++) {
                if(N % x == 0) {
                    time++;
                    if (!divisorGame(N - x)) {
                        if(time % 2 != 0) {
                            return true;
                        }
                    }
                }
            }
            return false;
        }
    }

    class Solution1 {
        //以上超时，但思路没错
        //那就要存储中间结果了
        int time;
        Map<Integer, Boolean> map = new HashMap<>();
        public boolean divisorGame(int N) {
            if(map.containsKey(N)){
                return map.get(N);
            }
            for (int x = 1; x < N; x++) {
                if(N % x == 0) {
                    time++;
                    if (!divisorGame(N - x)) {
                        if(time % 2 != 0) {
                            map.put(N, true);
                            return true;
                        }
                    }
                }
            }
            return false;
        }
    }

    class Solution2{
        //答案解释了为何次题目简单
        //选出任一 x，满足 0 < x < N 且 N % x == 0 。
        //N = 1的时候，区间(0, 1)中没有整数是n的因数，所以此时Alice败。
        //N = 2的时候，Alice只能拿1，N变成1，Bob无法继续操作，故Alice胜。
        //N = 3的时候，Alice只能拿1，N变成2，根据N = 2的结论，我们知道此时Bob会获胜，Alice败。
        //N = 4的时候，Alice能拿1或2，如果Alice拿1，根据N = 3的结论，Bob会失败，Alice会获胜。
        //N = 5的时候，Alice只能拿1，根据N = 4的结论，Alice会失败。
        //于是猜想N为奇数的时候Alice（先手）必败，N为偶数的时候Alice必胜。
        //证明
        //N = 1和N = 2时结论成立。
        //N > 2时，假设N ≤ k时该结论成立，则N = k + 1时：
        //如果k为偶数，则k + 1为奇数，x是k + 1的因数，只可能是奇数，而奇数减去奇数等于偶数，且k + 1 − x≤ k，故轮到Bob的时候都是偶数。
        //而根据我们的猜想假设N ≤ k的时候偶数的时候先手必胜，故此时无论Alice拿走什么，Bob都会处于必胜态，所以Alice处于必败态。
        //如果k为奇数，则k + 1为偶数，x可以是奇数也可以是偶数，若Alice减去一个奇数，那么k + 1 − x是一个小于等于k的奇数，此时Bob占有它，处于必败态，则 Alice 处于必胜态。
        //综上所述，这个猜想是正确的。
        public boolean divisorGame(int N) {
            return N % 2 == 0;
        }
    }

    public static void main(String[] args) {
        _1025 q = new _1025();
        System.out.println(q.new Solution1().divisorGame(2)); //true
        System.out.println(q.new Solution1().divisorGame(3)); //false
        System.out.println(q.new Solution1().divisorGame(5)); //false
        System.out.println(q.new Solution1().divisorGame(10)); //true
        System.out.println(q.new Solution1().divisorGame(196)); //false
    }
}

package algorithms.leetcode._1996_the_number_of_weak_characters_in_the_game;

import java.util.Arrays;


public class _1996 {
    // 你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。
    // 给你一个二维整数数组 properties ，其中 properties[i] = [attacki, defensei] 
    // 表示游戏中第 i 个角色的属性。

    // 如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。
    // 更正式地，如果认为角色 i 弱于 存在的另一个角色 j ，
    // 那么 attackj > attacki 且 defensej > defensei 。

    // 返回 弱角色 的数量。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    // 暴力
    class Solution {
        public int numberOfWeakCharacters(int[][] properties) {
            Arrays.sort(properties, (a, b) -> {
                return a[0] == b[0] ? (a[1] - b[1]) : (b[0] - a[0]);
            });
            int res = 0;
            int maxDef = 0;
            for (int[] prop : properties) {
                if (prop[1] < maxDef) {
                    res++;
                } else {
                    maxDef = prop[1];
                }
            }
            return res;
        }
    }

    class Prop {
        int attack;
        int defense;

        Prop(int attack, int defense) {
            this.attack = attack;
            this.defense = defense;
        }
    }

    public static void main(String[] args) {
        _1996 _1996 = new _1996();
        Solution solution = _1996.new Solution();
        System.out.println(solution.numberOfWeakCharacters(new int[][]{new int[] {2, 2}, new int[] {3, 3}}));
    }
}

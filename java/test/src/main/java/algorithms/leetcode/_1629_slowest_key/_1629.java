package algorithms.leetcode._1629_slowest_key;

public class _1629 {
    class Solution {
        public char slowestKey(int[] releaseTimes, String keysPressed) {
            int max = releaseTimes[0];
            char c = keysPressed.charAt(0);
            for (int i = 1; i < releaseTimes.length; i++) {
                int time = releaseTimes[i] - releaseTimes[i - 1];
                if (time > max || (time == max && c < keysPressed.charAt(i))) {
                    max = time;
                    c = keysPressed.charAt(i);
                }
            }
            return c;
        }
    }

    public static void main(String[] args) {
        _1629 _1629 = new _1629();
        Solution solution = _1629.new Solution();
        System.out.println(solution.slowestKey(new int[] {23,34,43,59,62,80,83,92,97}, "qgkzzihfc"));
    }
}

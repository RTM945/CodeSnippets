package algorithms.leetcode._042_trapping_rain_water;


public class _042 {

  class Solution {
    public int trap(int[] height) {
      int sum = 0;
      // 两端不会有水
      for (int i = 1; i < height.length - 1; i++) {
        // 左边最高的板
        int max_left = 0;
        for (int j = i - 1; j >= 0; j--) {
          if (height[j] > max_left) {
            max_left = height[j];
          }
        }
        // 右边最高的板
        int max_right = 0;
        for (int j = i + 1; j < height.length; j++) {
          if (height[j] > max_right) {
            max_right = height[j];
          }
        }
        // 短板
        int min = Math.min(max_left, max_right);
        if (min > height[i]) {
          sum += min - height[i];
        }
      }
      return sum;
    }
  }

  class Solution1 {
    public int trap(int[] height) {
      int sum = 0;
      int[] max_left = new int[height.length];
      int[] max_right = new int[height.length];

      for (int i = 1; i < height.length - 1; i++) {
        max_left[i] = Math.max(max_left[i - 1], height[i - 1]);
      }
      for (int i = height.length - 2; i >= 0; i--) {
        max_right[i] = Math.max(max_right[i + 1], height[i + 1]);
      }
      for (int i = 1; i < height.length - 1; i++) {
        int min = Math.min(max_left[i], max_right[i]);
        sum += min - height[i];
      }
      return sum;
    }
  }

  public static void main(String[] args) {
        _042 q = new _042();
        Solution s = q.new Solution();
        System.out.println(s.trap(new int[] {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}));
    }
}

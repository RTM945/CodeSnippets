package algorithms.leetcode._733_flood_fill;

import java.util.Arrays;

/* 
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。
给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，
让你重新上色这幅图像。
为了完成上色工作，从初始坐标开始，
记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，
接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，
……，重复该过程。
将所有有记录的像素点的颜色值改为新的颜色值。
最后返回经过上色渲染后的图像。
示例 1:
输入: 
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析: 
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
注意:
image 和 image[0] 的长度在范围 [1, 50] 内。
给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flood-fill
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
public class _733 {
    class Solution {
        // easy图像渲染
        // 听上去很让人感兴趣
        // 画了个图了解，看来递归标记就能做
        // 标记后，在遍历，将标记的数值替换为新的数值
        // 又想了想，先标记再替换会引入额外步骤，可不可以遍历时就做了
        // 知道sr sc，需要替换的color就知道了，所以遍历到的该color就能直接替换了
        // 而且从方法带返回参数来看，需要不改动原image
        // 那么需要复制
        // 有个问题是如果替换的和原来的颜色相同，dfs会死循环
        // 直接在入口干掉
        int color = -1;
        public int[][] floodFill(int[][] image, int sr, int sc, int newColor) {
            color = image[sr][sc];
            if(newColor == color) {
                return image;
            }
            int[][] res = new int[image.length][image[0].length];
            for (int i = 0; i < res.length; i++) {
                res[i] = Arrays.copyOf(image[i], image[i].length);
            }
            dfs(res, sr, sc, newColor);
            return res;
        }

        void dfs(int[][] image, int sr, int sc, int newColor) {
            if (sr < 0 || sr >= image.length || sc < 0 || sc >= image[0].length || image[sr][sc] != color) {
                return;
            }
            image[sr][sc] = newColor;
            dfs(image, sr - 1, sc, newColor); // 上
            dfs(image, sr + 1, sc, newColor); // 下
            dfs(image, sr, sc - 1, newColor); // 左
            dfs(image, sr, sc + 1, newColor); // 右
        }
    }

    public static void main(String[] args) {
        _733 q = new _733();
        int[][] image = { { 1, 1, 1 }, { 1, 1, 0 }, { 1, 0, 1 } };
        q.new Solution().floodFill(image, 1, 1, 2);
    }
}
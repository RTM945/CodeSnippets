package algorithms.leetcode._661_image_smoother;

public class _661 {
    // 图像平滑器 是大小为 3 x 3 的过滤器，用于对图像的每个单元格平滑处理，
    // 平滑处理后单元格的值为该单元格的平均灰度。

    // 每个单元格的  平均灰度 定义为：该单元格自身及其周围的 8 个单元格的平均值，结果需向下取整。
    // （即，需要计算蓝色平滑器中 9 个单元格的平均值）。

    // 如果一个单元格周围存在单元格缺失的情况，则计算平均灰度时不考虑缺失的单元格
    // （即，需要计算红色平滑器中 4 个单元格的平均值）。

    // 来源：力扣（LeetCode）
    // 链接：https://leetcode-cn.com/problems/image-smoother
    // 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    class Solution {
        public int[][] imageSmoother(int[][] img) {
            int r = img.length;
            int c = img[0].length;

            int[][] nimg = new int[r][c];
            for (int i = 0; i < r; i++) {
                for (int j = 0; j < c; j++) {
                    int sum = 0;
                    int num = 0;
                    for (int x = i - 1; x <= i + 1; x++) {
                        for (int y = j - 1; y <= j + 1; y++) {
                            if (x >= 0 && x < r && y >= 0 && y < c) {
                                num++;
                                sum += img[x][y];
                            }
                        }
                    }
                    nimg[i][j] = sum / num;
                }
            }
            return nimg;
        }
    }
}

package algorithms.dp;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

//最短通路问题
//问题: 求城市间最短通路. 设有如图所示的N座城市, 相邻城市之间有若干条通路,
//线上的数字表示通路的距离. 试求出从A到D的最短距离.
public class DP_1 {
    //使用邻接矩阵表示图
    static int[][] graph = {
            //A B1 B2 C1 C2 C3 D
            {0, 5, 2, 0, 0, 0, 0}, //A
            {0, 0, 0, 3, 2, 0, 0}, //B1
            {0, 0, 0, 0, 7, 4, 0}, //B2
            {0, 0, 0, 0, 0, 0, 4}, //C1
            {0, 0, 0, 0, 0, 0, 3}, //C2
            {0, 0, 0, 0, 0, 0, 5}, //C3
            {0, 0, 0, 0, 0, 0, 0}, //D
    };

    //寻路
    //使用dp
    //局部最优不一定能保证结果最优
    //动态规划步骤
    //1. 描述一个最优解的结构；
    //2. 递归地定义最优解的值；
    //3. 以自底向上的方式计算最优解的值；
    //4. 从已计算的信息中构建出最优解的路径。
    //本题的解为
    //Min(A到B1的距离+B1到D的距离的最优解, A到B2的距离+B2到D的距离的最优解)
    //而B1到D的距离的最优解又为Min(B1到C1的距离+C1到D的距离的最优解, B1到C2的距离+C2到D的距离的最优解)
    //类推下去，问题的最优解被分成了若干个子问题的最优解
    //这一性质被称为最优子结构
    //最优子结构是可用动态规划来解决问题的标志之一
    //但如果满足递归寻找最优解，则没有达到动态规划的目的
    //在得出最优子结构过程中，可以得到多个需要被重复用到的路径，比如C2到D，一般是递归底层的运算
    //越是底层，被用到的次数就越多，递归法每次都会重复计算，这些重复计算需要被消除
    //必须做到二件事:
    //1. 自底向上计算子问题最优解. 越是靠近底层的子问题, 越是先计算出来.
    //2. 用表格存储计算出的子问题的最优解的值, 以便在计算它的上层子问题时能够直接引用, 而不是去重新计算.
    static int minPathLen() {
        //用来存储每个结点间的最短路径
        int[] minPaths = new int[7];
        //最后一个结点自己和自己的距离肯定是0
        minPaths[6] = 0;
        for (int i = 5; i >= 0 ; i--) {
            minPaths[i] = Integer.MAX_VALUE;//设置最大值方便比较
            for (int j = i + 1; j <= 6; j++) {
                if(graph[i][j] != 0) { //如果有通路
                    //这条通路与下一个最短路径的和
                    int len = minPaths[j] +  graph[i][j];
                    if(len < minPaths[i]) {
                        minPaths[i] = len;
                    }
                }
            }
        }
        return minPaths[0]; //得出A到D最短路径
    }

    //但大多数时候，得出A到D的行进顺序是更有价值的问题
    //要得到最优解的路径(最优解的形成过程), 需要在计算的过程中记录额外的过程信息.
    static int[] minPath() {
        //用来存储每个结点间的最短路径
        int[] minPaths = new int[7];
        //使用另一个数组来记录路程中的结点
        int[] nodes = new int[7];
        //最后一个结点自己和自己的距离肯定是0
        minPaths[6] = 0;
        for (int i = 5; i >= 0 ; i--) {
            minPaths[i] = Integer.MAX_VALUE;//设置最大值方便比较
            nodes[i] = 0;//下一个结点默认值0
            for (int j = i + 1; j <= 6; j++) {
                if(graph[i][j] != 0) { //如果有通路
                    //这条通路与下一个最短路径的和
                    int len = minPaths[j] +  graph[i][j];
                    if(len < minPaths[i]) {
                        minPaths[i] = len;
                        //下一个结点是j
                        nodes[i] = j;
                    }
                }
            }
        }
        return nodes; //得出A到D最短路径
    }

    public static void main(String[] args) {
        System.out.println(minPathLen());
        String[] seq = {"A", "B1", "B2", "C1", "C2", "C3", "D"};
        int[] minPath = minPath();
        int p = minPath[0];
        System.out.print(seq[0]);
        while(p != 0) {
            System.out.print("->" + seq[p]);
            p = minPath[p];
        }
    }
}

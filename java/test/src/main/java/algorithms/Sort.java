package algorithms;

import java.util.Arrays;

/**
 * 一些排序算法实现
 */
public class Sort {

    static void insertionSort(int[] array) {
        for (int i = 1; i < array.length; i++) {
            int x = array[i];
            int j = i - 1;
            // 如果前面的数比后面的数字大，前面的数往后移一位
            while (j >= 0 && array[j] > x) {
                array[j + 1] = array[j];
                j = j - 1;
            }
            // 将当前值插入合适的位置
            array[j + 1] = x;
        }
    }

    static void binInsertionSort(int[] array) {
        for (int i = 1; i < array.length; i++) {
            int x = array[i];
            // 使用二分查找
            int low = 0;
            int high = i;
            while (low < high) {
                int mid = (low + high) / 2;
                if (x >= array[mid]) {
                    low = mid + 1;
                } else {
                    high = mid;
                }
            }
            // 将当前值插入合适的位置
            // 将该位置后开始到i的数据全部往后移一位再将x插入该位置
//            for (int j = i - 1; j > low - 1; j--) {
//                array[j + 1] = array[j];
//            }
//            array[low] = x;
            // 或者从后往前，把x一个一个的交换到指定位置
            for (int j = i; j > low; j--) {
                int tmp = array[j - 1];
                array[j - 1] = array[j];
                array[j] = tmp;
            }
        }
    }

    static void shellSort(int[] array) {
        for (int gap = array.length / 2; gap > 0; gap /= 2) {
            for (int i = gap; i < array.length; i++) {
                int x = array[i];
                int j = i;
                while (j >= gap && array[j - gap] > x) {
                    array[j] = array[j - gap];
                    j = j - gap;
                }
                array[j] = x;
            }
        }
    }

    static void bubbleSort(int[] array) {
        for (int i = 0; i < array.length - 1; i++) {
            for (int j = 0; j < array.length - 1 - i; j++) {
                if (array[j] > array[j + 1]) {
                    int tmp = array[j + 1];
                    array[j + 1] = array[j];
                    array[j] = tmp;
                }
            }
        }
    }

    // 递归模板
    static void quickSort(int[] array, int low, int high) {
        if (low < high) {
            int j = partition(array, low, high);
            quickSort(array, low, j - 1);
            quickSort(array, j + 1, high);
        }
    }

    //用第一个元素做pivot
    static int partition(int[] array, int low, int high) {
        int pivot = array[low];
        int i = low;
        int j = high;
        while (i < j) {
            //从后往前找到一个小于pivot的
            while (i < j && array[j] > pivot) {
                j--; //跳过比pivot大的
            }
            if (i < j) {
                //交换
                array[i] = array[j]; //array[i]此时是pivot
                //其实这一步的一个便于理解的做法是
                //int temp = array[i];
                //array[i] = array[j];
                //array[j] = temp;
                //但因为已经记录了pivot，所以省略了
                //此时可以将array[j]看作pivot
                i++; // 从下一位开始找比pivot大的
            }
            //从前往后找到一个大于pivot的
            while (i < j && array[i] < pivot) {
                i++; //跳过比pivot小的
            }
            if (i < j) {
                //交换
                array[j] = array[i];//这里把array[j]看作pivot
                //和上面相同，省略了显式的交换步骤
                //因为换来换去都是pivot
                j--;
            }
        }
        //i=j时跳出循环
        array[i] = pivot; //填上pivot
        return i;
    }

    static void selectionSort(int[] array) {
        for (int i = 0; i < array.length - 1; i++) {
            int index = i;
            for (int j = i + 1; j < array.length; j++) {
                if (array[j] < array[index]) {
                    //选择小的
                    index = j;
                }
            }
            if (index != i) {
                // 交换
                int temp = array[i];
                array[i] = array[index];
                array[index] = temp;
            }
        }
    }

    static void heapSort(int[] array) {
        // 因为完全二叉树编号为i的右子树结点为2i+1
        // 所以从数组构造完全二叉树，根结点在数组中下标的最大值为array.length / 2 - 1
        int root = array.length / 2 - 1;
        int n = array.length - 1;
        for (int i = root; i >= 0; i--) {
            maxHeapify(array, i, n);
        }

        // 构建堆后，令数组有序
        // 每次移出根结点（与最后一个元素交换，再令堆的大小-1）
        for (int i = n; i > 0; i--) {
            int temp = array[0];
            array[0] = array[i];
            array[i] = temp;
            maxHeapify(array, 0, i - 1);
        }
    }

    //小顶堆化
    static void maxHeapify(int[] array, int i, int n) {
        int li = i * 2 + 1; // 左子节点的下标
        if (li > n) {
            return; //没有子节点了
        }
        int ri = li + 1; // 右子节点的下标
        int j = li;
        //左右子节点谁小
        if (ri <= n && array[ri] > array[li]) {
            j = ri;
        }
        //子节点比父结点小就交换位置
        if (array[j] > array[i]) {
            int temp = array[i];
            array[i] = array[j];
            array[j] = temp;
            //交换了就要重新检查该数据与新的子节点之间的关系
            maxHeapify(array, j, n);
        }
    }

//    static void mergeSort(int[] array, int[] reg, int low, int high) {
//        if (low < high) {
//            int mid = low + (high - low) / 2;
//            mergeSort(array, reg, low, mid);
//            mergeSort(array, reg, mid + 1, high);
//            merge(array, reg, low, mid, high);
//        }
//    }
//
//    static void merge(int[] array, int[] reg, int low, int mid, int high) {
//        int i = low; //第一个子集的初始下标
//        int j = mid + 1; //第二个子集的初始下标
//        int t = 0; //辅助集合的下标
//        while(i <= mid && j <= high) {
//            if(array[i] <= array[j]) {
//                //将小的元素插入辅助集合
//                reg[t++] = array[i++]; //同时对应子集下标向后移一位
//            }else{
//                reg[t++] = array[j++];
//            }
//        }
//        //如果子集还有剩下的元素，要放入reg中
//        while (i <= mid) {
//            reg[t++] = array[i++];
//        }
//        while (j <= high) {
//            reg[t++] = array[j++];
//        }
//        //循环结束后，reg中就是基本有序且归并好的集合了
//        t = 0;
//        //把reg中的元素复制到源集合中，用有序的数据覆盖掉原来无序的相同数据
//        while (low <= high) {
//            array[low++] = reg[t++];
//        }
//    }


    static void mergeSort(int[] array, int[] reg) {
        int len = array.length;
        //令集合长度为2, 4, 8, 16...
        for (int width = 1; width < len; width *= 2) {
            //得到每个集合的起始下标
            for (int i = 0; i < len - 1; i += 2 * width) {
                int mid = i + width - 1;
                int high = i + 2 * width - 1;
                if(mid > len - 1) {
                    mid = len - 1;
                }
                if(high > len - 1) {
                    high = len - 1;
                }
                //合并array[i:i+width-1]和array[i+width:i+2*width-1]
                //当i + width > len - 1时，其实只剩array[i:len-1]了
                merge(array, reg, i, mid, high);
            }
        }
    }

    static void merge(int[] array, int[] reg, int low, int mid, int high) {
        int i = low; //一个子集的起始下标
        int j = mid + 1; //另一个子集的起始下标
        int t = 0; //辅助集合下标
        while(i <= mid && j <= high) {
            if(array[i] <= array[j]) {
                //将小的元素插入辅助集合
                reg[t++] = array[i++]; //同时对应子集下标向后移一位
            }else{
                reg[t++] = array[j++];
            }
        }
        //如果子集还有剩下的元素，要放入reg中
        while (i <= mid) {
            reg[t++] = array[i++];
        }
        while (j <= high) {
            reg[t++] = array[j++];
        }
        //循环结束后，reg中就是基本有序且归并好的集合了
        t = 0;
        //把reg中的元素复制到源集合中，用有序的数据覆盖掉原来无序的相同数据
        while (low <= high) {
            array[low++] = reg[t++];
        }
    }

    public static void main(String[] args) {
        int[] a = {1, 3, 4, 8, 2, 5, 0, 7, 6, 9};
//        binInsertionSort(a);
//        shellSort(a);
//        bubbleSort(a);
//        quickSort(a, 0, a.length - 1);
//        Arrays.sort(a);
//        selectionSort(a);
//        heapSort(a);
//        mergeSort(a, new int[a.length], 0, a.length - 1);
        mergeSort(a, new int[a.length]);
        System.out.println(Arrays.toString(a));
    }

}

package algorithms.leetcode.offer09;

//比较有效率的解法是使用数组自己实现栈操作，操作下标指针，比java的Stack性能好
//Stack已不被java推荐使用，可以用LinkedList替代
//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/solution/mian-shi-ti-09-yong-liang-ge-zhan-shi-xian-dui-l-3/467970
public class CQueue02 {

    final int[] enqueueStack;
    final int[] dequeueStack;
    int enqueueIndex = -1;
    int dequeueIndex = -1;

    public CQueue02() {
        this.enqueueStack = new int[100];
        this.dequeueStack = new int[100];
    }

    public void appendTail(int value) {
        enqueueStack[++enqueueIndex] = value;
    }

    public int deleteHead() {
        if(dequeueIndex == -1) {
            while(enqueueIndex != -1) {
                //
                dequeueStack[++dequeueIndex] = enqueueStack[enqueueIndex--];
            }
        }
        return dequeueIndex == -1 ? -1 : dequeueStack[dequeueIndex--];
    }

    public static void main(String[] args) {
        CQueue02 obj = new CQueue02();
//        for (int i = 0; i < 10; i++) {
//            obj.appendTail(i);
//        }
//        for (int i = 0; i < 10; i++) {
//            System.out.println(obj.deleteHead());
//        }
        obj.appendTail(1);
        obj.appendTail(2);
        obj.appendTail(3);
        System.out.println(obj.deleteHead());
        System.out.println(obj.deleteHead());
        obj.appendTail(4);
        obj.appendTail(5);
        System.out.println(obj.deleteHead());
        System.out.println(obj.deleteHead());
        System.out.println(obj.deleteHead());
    }
}

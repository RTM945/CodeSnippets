package algorithms.leetcode.offer009;

import java.util.Stack;

//用两个栈实现一个队列。队列的声明如下
//请实现它的两个函数 appendTail 和 deleteHead
//分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof

//一般路过解法是初始化两个栈，出队时将有数据的栈倒入空栈，这样就能输出先进入的元素
//入队时操作相同，将有数据的栈倒入空栈，能保证元素的顺序
public class CQueue01 {

    Stack<Integer> stack1;
    Stack<Integer> stack2;

    public CQueue01() {
        stack1 = new Stack<>();
        stack2 = new Stack<>();
    }

    public void appendTail(int value) {
        while (!stack1.isEmpty()) {
            stack2.push(stack1.pop());
        }
        stack2.push(value);
    }

    public int deleteHead() {
        while (!stack2.isEmpty()) {
            stack1.push(stack2.pop());
        }
        if(stack1.isEmpty()) {
            return -1;
        }
        return stack1.pop();
    }

    public static void main(String[] args) {
        CQueue01 obj = new CQueue01();
        for (int i = 0; i < 10; i++) {
            obj.appendTail(i);
        }
        for (int i = 0; i < 10; i++) {
            System.out.println(obj.deleteHead());
        }
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */

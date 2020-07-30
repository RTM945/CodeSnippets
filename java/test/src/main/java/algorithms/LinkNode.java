package algorithms;

public class LinkNode {

    int val;
    LinkNode next;

    public LinkNode(int val) {
        this.val = val;
    }

    public LinkNode reverse() {
        LinkNode node = this;
        LinkNode next = null;
        LinkNode prev = null;
        LinkNode rare = null;
        while(node != null) {
            next = node.next;
            if(next == null) {
                rare = node;
            }
            node.next = prev; //当前结点与next断开
            prev = node; //当前结点变为前驱
            node = next; //指针移动到下一个结点
        }
        return rare;
    }

    //递归法
    public LinkNode reverser() {
        return reverser(this);
    }

    private LinkNode reverser(LinkNode node) {
        if(node == null || node.next == null) {
            return node;
        }
        LinkNode next = reverser(node.next); //一直向后到达最后一个结点
        node.next.next = node; //将前驱变为后继以转换
        node.next = null; //断掉前驱与当前节点的前后关系
        return next;
    }

    LinkNode appendNext(LinkNode parent, int val) {
        if(parent == null) {
            return null;
        }
        parent.next = new LinkNode(val);
        return parent.next;
    }

    public String toString() {
        LinkNode node = this;
        StringBuilder sb = new StringBuilder("LinkNode[");
        while(node != null) {
            sb.append(node.val).append(',');
            node = node.next;
        }
        sb.deleteCharAt(sb.length() - 1);
        sb.append(']');
        return sb.toString();
    }
    
    public static void main(String[] args) {
        LinkNode head = new LinkNode(0);
        LinkNode lNode = head;
        lNode = lNode.appendNext(lNode, 1);
        lNode = lNode.appendNext(lNode, 2);
        lNode = lNode.appendNext(lNode, 3);
        lNode = lNode.appendNext(lNode, 4);
        lNode = lNode.appendNext(lNode, 5);
        System.out.println(head);
        head = head.reverse();
        System.out.println(head);
        head = head.reverser();
        System.out.println(head);
    }
    
}
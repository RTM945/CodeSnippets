package algorithms;

import java.util.LinkedList;
import java.util.Queue;
import java.util.Stack;

/**
 * Binary Tree实现
 * @param <T>
 */
public class BinaryTree<T> {

    private Node<T> root;

    private static class Node<T> {
        T data;
        Node<T> left;
        Node<T> right;

        public Node(T data) {
            this.data = data;
        }
    }

    public static <T> BinaryTree<T> fromLevelOrder(T[] arrays) {
        BinaryTree<T> tree = new BinaryTree<>();
        tree.root = fromLevelOrder(arrays, 0);
        return tree;
    }

    private static <T> Node<T> fromLevelOrder(T[] arrays, int index) {
        if (index < arrays.length) {
            Node<T> node = new Node<>(arrays[index]);
            node.left = fromLevelOrder(arrays, index * 2 + 1);
            node.right = fromLevelOrder(arrays, index * 2 + 2);
            return node;
        }
        return null;
    }

    public static <T> BinaryTree<T> fromLevelOrderByQueue(T[] arrays) {
        BinaryTree<T> tree = new BinaryTree<>();
        for (T t : arrays) {
            tree.insert(t);
        }
        return tree;
    }

    public void insert(T data) {
        if (root == null) {
            root = new Node<>(data);
            return;
        }
        Queue<Node<T>> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            Node<T> tmp = queue.poll();
            if (tmp.left == null) {
                tmp.left = new Node<>(data);
                break;
            } else {
                queue.add(tmp.left);
                if (tmp.right == null) {
                    tmp.right = new Node<>(data);
                    break;
                } else {
                    queue.add(tmp.right);
                }
            }
        }
    }

    public interface TreeTraversal<T> {
        void apply(T data);
    }

    public void preOrderTraversal(TreeTraversal<T> traversal) {
        preOrderTraversal(root, traversal);
    }

    //dfs is the same
    private void preOrderTraversal(Node<T> node, TreeTraversal<T> traversal) {
        if (node == null) {
            return;
        }
        traversal.apply(node.data);
        preOrderTraversal(node.left, traversal);
        preOrderTraversal(node.right, traversal);
    }


    public void inOrderTraversal(TreeTraversal<T> traversal) {
        inOrderTraversal(root, traversal);
    }

    private void inOrderTraversal(Node<T> node, TreeTraversal<T> traversal) {
        if (node == null) {
            return;
        }
        inOrderTraversal(node.left, traversal);
        traversal.apply(node.data);
        inOrderTraversal(node.right, traversal);
    }

    private void inOrderTraversalUnRecursively(TreeTraversal<T> traversal) {
        Stack<Node<T>> stack = new Stack<>();
        Node<T> node = root;
        while(node != null || !stack.isEmpty()){
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            node = stack.pop();
            traversal.apply(node.data);
            node = node.right;
        }
    }

    public void postOrderTraversal(TreeTraversal<T> traversal) {
        postOrderTraversal(root, traversal);
    }

    private void postOrderTraversal(Node<T> node, TreeTraversal<T> traversal) {
        if (node == null) {
            return;
        }
        postOrderTraversal(node.left, traversal);
        postOrderTraversal(node.right, traversal);
        traversal.apply(node.data);
    }

    private void postOrderTraversalUnRecursively(TreeTraversal<T> traversal) {
        Stack<Node<T>> stack1 = new Stack<>();
        Stack<Node<T>> stack2 = new Stack<>();
        stack1.push(root);
        while (!stack1.isEmpty()) {
            Node<T> node = stack1.pop();
            stack2.push(node);
            if(node.left != null) {
                stack1.push(node.left);
            }
            if(node.right != null) {
                stack1.push(node.right);
            }
        }
        while (!stack2.isEmpty()) {
            Node<T> node = stack2.pop();
            traversal.apply(node.data);
        }
    }

    public void bfs(TreeTraversal<T> traversal) {
        Queue<Node<T>> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            Node<T> node = queue.poll();
            traversal.apply(node.data);
            if (node.left != null) {
                queue.add(node.left);
            }
            if (node.right != null) {
                queue.add(node.right);
            }
        }
    }

    public void bfsRecursively(TreeTraversal<T> traversal) {
        int level = 1;
        while (bfsRecursively(root, level, traversal)) {
            level++;
        }
    }

    private boolean bfsRecursively(Node<T> node, int level, TreeTraversal<T> traversal) {
        if (node == null) {
            return false;
        }
        if (level == 1) {
            traversal.apply(node.data);
            return true;
        }
        boolean left = bfsRecursively(node.left, level - 1, traversal);
        boolean right = bfsRecursively(node.right, level - 1, traversal);
        return left || right;
    }

    //no recursive
    public void dfs(TreeTraversal<T> traversal) {
        Stack<Node<T>> stack = new Stack<>();
        stack.push(root);
        while (!stack.isEmpty()) {
            Node<T> node = stack.pop();
            traversal.apply(node.data);
            if(node.right != null) {
                stack.push(node.right);
            }
            if(node.left != null) {
                stack.push(node.left);
            }
        }
    }

    public static void main(String[] args) {
        Integer[] arr = {0, 1, 2, 3, 4, 5, 6};
        BinaryTree<Integer> tree = BinaryTree.fromLevelOrderByQueue(arr);
        System.out.println("preOrderTraversal");
        tree.preOrderTraversal(System.out::print);
        System.out.println();
        System.out.println("dfs");
        tree.dfs(System.out::print);
        System.out.println();
        System.out.println("inOrderTraversal");
        tree.inOrderTraversal(System.out::print);
        System.out.println();
        System.out.println("inOrderTraversalUnRecursively");
        tree.inOrderTraversalUnRecursively(System.out::print);
        System.out.println();
        System.out.println("postOrderTraversal");
        tree.postOrderTraversal(System.out::print);
        System.out.println();
        System.out.println("postOrderTraversalUnRecursively");
        tree.postOrderTraversalUnRecursively(System.out::print);
        System.out.println();
        System.out.println("bfs");
        tree.bfs(System.out::print);
        System.out.println();
        System.out.println("bfsRecursive");
        tree.bfsRecursively(System.out::print);
    }

}

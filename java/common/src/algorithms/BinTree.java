package algorithms;

import java.util.LinkedList;
import java.util.Queue;
import java.util.Stack;

public class BinTree<T> {

    private TreeNode<T> root;

    private static class TreeNode<T> {
        T data;
        TreeNode<T> left;
        TreeNode<T> right;

        public TreeNode(T data) {
            this.data = data;
        }
    }

    public static <T> BinTree<T> fromLevelOrder(T[] arrays) {
        BinTree<T> tree = new BinTree<>();
        tree.root = fromLevelOrder(arrays, 0);
        return tree;
    }

    private static <T> TreeNode<T> fromLevelOrder(T[] arrays, int index) {
        if (index < arrays.length) {
            TreeNode<T> node = new TreeNode<>(arrays[index]);
            node.left = fromLevelOrder(arrays, index * 2 + 1);
            node.right = fromLevelOrder(arrays, index * 2 + 2);
            return node;
        }
        return null;
    }

    public static <T> BinTree<T> fromLevelOrderByQueue(T[] arrays) {
        BinTree<T> tree = new BinTree<>();
        for (T t : arrays) {
            tree.insert(t);
        }
        return tree;
    }

    public void insert(T data) {
        if (root == null) {
            root = new TreeNode<>(data);
            return;
        }
        Queue<TreeNode<T>> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            TreeNode<T> tmp = queue.poll();
            if (tmp.left == null) {
                tmp.left = new TreeNode<>(data);
                break;
            } else {
                queue.add(tmp.left);
                if (tmp.right == null) {
                    tmp.right = new TreeNode<>(data);
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
    private void preOrderTraversal(TreeNode<T> node, TreeTraversal<T> traversal) {
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

    private void inOrderTraversal(TreeNode<T> node, TreeTraversal<T> traversal) {
        if (node == null) {
            return;
        }
        inOrderTraversal(node.left, traversal);
        traversal.apply(node.data);
        inOrderTraversal(node.right, traversal);
    }

    public void postOrderTraversal(TreeTraversal<T> traversal) {
        postOrderTraversal(root, traversal);
    }

    private void postOrderTraversal(TreeNode<T> node, TreeTraversal<T> traversal) {
        if (node == null) {
            return;
        }
        postOrderTraversal(node.left, traversal);
        postOrderTraversal(node.right, traversal);
        traversal.apply(node.data);
    }

    public void bfs(TreeTraversal<T> traversal) {
        Queue<TreeNode<T>> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            TreeNode<T> node = queue.poll();
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

    private boolean bfsRecursively(TreeNode<T> node, int level, TreeTraversal<T> traversal) {
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
        Stack<TreeNode<T>> stack = new Stack<>();
        stack.push(root);
        while (!stack.isEmpty()) {
            TreeNode<T> node = stack.pop();
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
        BinTree<Integer> tree = BinTree.fromLevelOrderByQueue(arr);
        System.out.println("preOrderTraversal");
        tree.preOrderTraversal(System.out::print);
        System.out.println();
        System.out.println("inOrderTraversal");
        tree.inOrderTraversal(System.out::print);
        System.out.println();
        System.out.println("postOrderTraversal");
        tree.postOrderTraversal(System.out::print);
        System.out.println();
        System.out.println("bfsRecursive");
        tree.bfsRecursively(System.out::print);
        System.out.println();
        System.out.println("bfs");
        tree.bfs(System.out::print);
        System.out.println();
        System.out.println("dfs");
        tree.dfs(System.out::print);
    }

}

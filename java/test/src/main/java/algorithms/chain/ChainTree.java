package algorithms.chain;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

public class ChainTree {

    
    static int mapw = 5;
    static int maph = 5;
    
    // 下 右 为正方向
    static int[][] map = new int[maph][mapw];

    // 周围一圈8格
    static int[][] range1 = {
        {0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {-1, -1}, {1, -1},
    };

    // 十字形4格
    static int[][] range2 = {
        {0, 1}, {0, -1}, {-1, 0}, {1, 0}, 
    };

    static class Unit {
        int id;
        int row, col;

        Unit(int id, int row, int col) {
            this.id = id;
            // 地图数组的行列
            this.row = row;
            this.col = col;
        }

        @Override
        public String toString() {
            return "Unit [id = " + id + ", row = " + row + ", col = " + col + "]";
        }
    }
    static Map<Integer, Unit> units = new HashMap<>();
    
    // 0 0 1 2 3
    // 0 0 4 0 0
    // 0 5 6 0 0 
    // 0 0 7 0 0
    // 0 0 8 0 0
    static {
        enterMap(new Unit(1, 0, 2));
        enterMap(new Unit(2, 0, 3));
        enterMap(new Unit(3, 0, 4));
        enterMap(new Unit(4, 1, 2));
        enterMap(new Unit(5, 2, 1));
        enterMap(new Unit(6, 2, 2));
        enterMap(new Unit(7, 3, 2));
        enterMap(new Unit(8, 4, 2));
    }

    static void enterMap(Unit unit) {
        units.put(unit.id, unit);
        map[unit.row][unit.col] = unit.id;
    }

    static List<Unit> getTarget(Unit center, int[][] range) {
        List<Unit> res = new ArrayList<>();
        for (int[] vec : range) {
            int nrow = center.row + vec[0];
            if (nrow >= maph || nrow < 0) {
                continue;
            }
            int ncol = center.col + vec[1];
            if (ncol >= mapw || ncol < 0) {
                continue;
            }
            int id = map[nrow][ncol];
            if (id > 0) {
                res.add(units.get(id));
            }
        }
        return res;
    }

    static class Node {
        int id;
        List<Node> childNodes = new ArrayList<>();
    }


    // 记忆
    static Map<Integer, Node> nodes = new HashMap<>();

    static Node toNode(Unit n, Set<Integer> visited) {
        Node node = new Node();
        node.id = n.id;
        visited.add(n.id);
        List<Unit> targets = getTarget(n, range2);
        for (Unit target : targets) {
            if (visited.contains(target.id)) {
                continue;
            }
            if (nodes.containsKey(target.id)) {
                node.childNodes.add(nodes.get(target.id));
            } else {
                Node child = toNode(target, visited);
                node.childNodes.add(child);
                nodes.put(target.id, child);
            }
            
        }
        visited.remove(n.id);

        return node;
    }

    static List<List<Integer>> res = new ArrayList<>();

    static void dfs(Node node, List<Integer> path) {
        path.add(node.id);
        List<Node> list = node.childNodes;
        if (list.isEmpty()) {
            List<Integer> r = new ArrayList<>(path);
            res.add(r);
        } else {
            for (Node n : list) {
                dfs(n, path);
            }
        }
        path.remove(path.size() - 1);
    }
    
    public static void main(String[] args) {
        // 从1出发
        Unit u = units.get(1);
        // 构造树
        Set<Integer> visited = new HashSet<>();
        Node root = toNode(u, visited);
        // dfs找最长路径
        List<Integer> path = new ArrayList<>();
        dfs(root, path);
        for (List<Integer> p : res) {
            System.out.println(p.stream().map(String::valueOf).collect(Collectors.joining("->")));
        }
    }

}

package algorithms.chain;

import java.util.*;

public class Chain {

    static Map<Integer, int[]> graph = new HashMap<>();
    static List<List<Integer>> paths = new ArrayList<>();

    // 有环
    static {
        //            3
        //           /
        //  1---4---2
        //    \ | \ |
        //      5---0
        graph.put(0, new int[]{2, 4, 5});
        graph.put(1, new int[]{4, 5});
        graph.put(2, new int[]{3, 4, 0});
        graph.put(3, new int[]{2});
        graph.put(4, new int[]{0, 1, 5, 2});
        graph.put(5, new int[]{0, 1, 4});
    }

    public static void main(String[] args) {
        // 从0出发
        for (int p : graph.get(0)) {
            LinkedList<Integer> path = new LinkedList<>();
            path.add(0);
            Set<Integer> visit = new HashSet<>();
            visit.add(0);
            dfs(p, path, visit);
            System.out.println(path);
        }
        for(List<Integer> path : paths) {
            System.out.println(path);
        }
    }

    private static void dfs(int start, LinkedList<Integer> path, Set<Integer> visit) {
        path.addLast(start);
        visit.add(start);
        int[] ints = graph.get(start);
        List<Integer> targets = new ArrayList<>();
        for (int p : ints) {
            if (!visit.contains(p)) {
                targets.add(p);
            }
        }
        if (targets.isEmpty()) {
            paths.add(new ArrayList<>(path));
        } else {
            for (int p : targets) {
                dfs(p, path, visit);
            }
        }
        // 回溯
        visit.remove(start);
        path.removeLast();
    }

}


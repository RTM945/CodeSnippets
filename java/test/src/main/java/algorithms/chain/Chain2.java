package algorithms.chain;

import java.util.*;

import org.luaj.vm2.lib.PackageLib.require;

public class Chain2 {

    // 周围一圈8格
    static int[][] range = {
        {0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {-1, -1}, {1, -1},
    };

    static Map<Integer, int[]> graph = new HashMap<>();

    static int total = 0;

    static void gengraph(int row, int col) {
        // 凑在一起的单位
        // 每个单位连接周围八向
        int[][] arr = new int[row][];
        for (int i = 0; i < row; i++) {
            arr[i] = new int[col];
            for (int j = 0; j < col; j++) {
                arr[i][j] = col * i + j + 1;
            }
        }
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                System.out.print(arr[i][j] + " ");
            }
            System.out.print("\n");
        }

        // 做邻接表
        for (int i = 1; i < row - 1; i++) {
            for (int j = 1; j < col - 1; j++) {
                int[] edge = new int[range.length];
                int v = arr[i][j];
                for (int k = 0; k < range.length; k++) {
                    int[] dir = range[k];
                    edge[k] = arr[i + dir[0]][j + dir[1]];
                }
                graph.put(v, edge);
            }   
        }

        // 外围一圈
        // 4个角
        int leftup = arr[0][0];
        graph.put(leftup, new int[]{leftup + 1, leftup + col, leftup + col + 1});
        int rightup = arr[0][col - 1];
        graph.put(rightup, new int[]{rightup - 1, rightup + col, rightup + col - 1});
        int leftdown = arr[row - 1][0];
        graph.put(leftdown, new int[]{leftdown - col, leftdown - col + 1, leftdown + 1});
        int leftright = arr[row - 1][col - 1];
        graph.put(leftright, new int[]{leftright - 1, leftright - col, leftright - col - 1});

        for (int i = 1; i < col - 1; i++) {
            // 上边
            int v = arr[0][i];
            graph.put(v, new int[]{v - 1, v + col - 1, v + col, v + col + 1, v + 1});
            // 下边
            v = arr[row - 1][i];
            graph.put(v, new int[]{v - 1, v - col - 1, v - col, v - col + 1, v + 1});
        }
        for (int i = 1; i < row - 1; i++) {
            // 左边
            int v = arr[i][0];
            graph.put(v, new int[]{v - col, v - col + 1, v + 1, v + col + 1, v + col});
            // 右边
            v = arr[i][col - 1];
            graph.put(v, new int[]{v - col, v - col - 1, v - 1, v + col - 1, v + col});
        }

        for (Map.Entry<Integer, int[]> entry : graph.entrySet()) {
            System.out.print("顶点" + entry.getKey() + "的边 : ");
            for (int v : entry.getValue()) {
                System.out.print(v + " ");
            }
            System.out.print("\n");
        }
        total = col * row;
    }

    static List<List<Integer>> paths = new ArrayList<>();
    static List<Integer> maxPath = new ArrayList<>();

    public static void main(String[] args) {
        gengraph(5, 4);
        long now = System.currentTimeMillis();
        // 需要改为bfs!
        bfs(1);

        // 作死dfs
        //dfs(1);
        System.out.println("耗时: " + (System.currentTimeMillis() - now));
        System.out.println(maxPath);
    }

    static void bfs(int start) {
        Set<Integer> visit = new HashSet<>();
        visit.add(start);
        List<Integer> path = new ArrayList<>();
        path.add(start);
        LinkedList<Integer> q = new LinkedList<>();
        q.push(start);
        while (!q.isEmpty()) {
            int u = q.pop();
            System.out.print(u + "->");
            int[] ins = graph.get(u);
            for (int v : ins) {
                if (!visit.contains(v)) {
                    q.push(v);
                    visit.add(v);
                }
            }
        }
    }

    static void dfs(int start) {
        for (int p : graph.get(start)) {
            List<Integer> path = new ArrayList<>();
            path.add(start);
            Set<Integer> visit = new HashSet<>();
            visit.add(start);
            boolean goon = dfs(p, path, visit);
            if (!goon) {
                break;
            }
        }
    }

    

    private static boolean dfs(int start, List<Integer> path, Set<Integer> visit) {
        path.add(start);
        visit.add(start);
        int[] ints = graph.get(start);
        List<Integer> targets = new ArrayList<>();
        for (int p : ints) {
            if (!visit.contains(p)) {
                targets.add(p);
            }
        }
        if (targets.isEmpty()) {
            if (path.size() > maxPath.size()) {
                maxPath = new ArrayList<>(path);
                // 剪枝!
                if (path.size() == total) {
                    return false;
                }
            }
        } else {
            for (int p : targets) {
                boolean goon = dfs(p, path, visit);
                if (!goon) {
                    return false;
                }
            }
        }
        // 回溯
        visit.remove(start);
        path.remove(path.size() - 1);
        return true;
    }
    
}

package algorithms;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Objects;
import java.util.Random;
import java.util.Set;

import algorithms.AStar.AStarNode;
import algorithms.AStar.OpenList;

public class AStar {

    public final static int DIRECT_G = 10; // 横竖移动代价
	public final static int OBLIQUE_G = 14; // 斜移动代价

    protected final AStarNode[][] map;
	protected final int maxX;
	protected final int maxY;
	protected int id = 0; // 不维护closelist,提升性能,具体看调用

    protected AStarNode start;
	protected AStarNode end;
	protected boolean is4Dir = false;
	protected boolean goToEnd = false;

    public boolean isIs4Dir() {
		return is4Dir;
	}

	public void setIs4Dir(boolean is4Dir) {
		this.is4Dir = is4Dir;
	}

	public boolean isGoToEnd() {
		return goToEnd;
	}

	public void setGoToEnd(boolean goToEnd) {
		this.goToEnd = goToEnd;
	}

	protected final OpenList openList = new OpenList();

    public AStar(AStarNode[][] map) {
		this.map = map;
		this.maxX = map.length;
		this.maxY = map[0].length;
		for (int x = 0; x < maxX; x++) {
			AStarNode[] line = map[x];
			for (int y = 0; y < maxY; y++) {
				AStarNode node = line[y];
				node.x = x;
				node.y = y;
			}
		}
	}

    protected Object o; // 寻路参数

    public List<Node> findPath(Node _start, Node _end) {
		return findPath(_start, _end, null);
	}

    public List<Node> findPath(Node _start, Node _end, Object o) {
        init(_start, _end, o);

		open(start);
		while (!openList.isEmpty()) {
			AStarNode poll = openList.poll();
			printMap(poll, true);
			try {
				Thread.sleep(1);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
			if (poll == end) {
				List<Node> paths = new LinkedList<>();
				AStarNode end = this.end;
				while (end != null) {
					paths.add(new Node(end.x, end.y));
					end = end.tempNext;
				}
				return paths;
			}
			addAround(poll ,o);
		}
		return null;
    }

    public void printMap() {
		printMap(end, false);
	}

	public void printMap(AStarNode end, boolean printOpen) {
		Set<Node> paths = new HashSet<>();
		while (end != null) {
			paths.add(new Node(end.x, end.y));
			end = end.tempNext;
		}
		StringBuilder sb = new StringBuilder("===================\r\n");
		for (int i = 0; i < maxX; i++) {
			for (int j = 0; j < maxY; j++) {
				AStarNode node = map[i][j];
				if (!node.canIn(null, null)) {
					sb.append("■");
				} else {
					if (paths.contains(node)) {
						sb.append("●");
					} else {
						if (printOpen && isOpen(node)) {
							sb.append("o");
						} else {
							sb.append(".");
						}
					}
				}
				sb.append(" ");
			}
			sb.append("\r\n");
		}
		System.out.println(sb.toString());
	}

	protected void init(Node _start, Node _end, Object o) {
		this.o = o;

		this.start = map[_start.x][_start.y];
		this.start.tempNext = null;
		this.start.tempG = 0;
		this.start.tempF = 0;

		this.end = map[_end.x][_end.y];
		this.end.tempNext = null;

		openList.clear();
		id++;
	}

	protected void addAround(AStarNode poll, Object o) {
		tryOpen(poll.x, poll.y + 1, poll, DIRECT_G);
		tryOpen(poll.x, poll.y - 1, poll, DIRECT_G);
		tryOpen(poll.x + 1, poll.y, poll, DIRECT_G);
		tryOpen(poll.x - 1, poll.y, poll, DIRECT_G);
		if (is4Dir) {
			return;
		}
		tryOpen(poll.x + 1, poll.y + 1, poll, OBLIQUE_G);
		tryOpen(poll.x + 1, poll.y - 1, poll, OBLIQUE_G);
		tryOpen(poll.x - 1, poll.y + 1, poll, OBLIQUE_G);
		tryOpen(poll.x - 1, poll.y - 1, poll, OBLIQUE_G);
	}

	protected void tryOpen(int x, int y, AStarNode from, int g) {
		if (x >= maxX || x < 0) {
			return;
		}
		if (y >= maxY || y < 0) {
			return;
		}
		AStarNode node = map[x][y];
		if (isOpen(node) || isClose(node)) {
			return;
		}
		if (!node.canIn(from, o) && (goToEnd || node != end)) { // 终点不检测
			if (node.canClose(from, o)) {
				close(node);
			}
			return;
		}
		node.tempNext = from;
		node.tempG = from.tempG + g;
		node.tempF = calcH(node) + node.tempG;
		open(node);
	}

	protected int calcH(AStarNode node) {
		return (Math.abs(end.x - node.x) + Math.abs(end.y - node.y)) * 10;
	}

	protected void open(AStarNode node) {
		node.id = id;
		openList.add(node);
	}

	protected void close(AStarNode node) {
		node.id = -id;
	}

	protected boolean isOpen(AStarNode node) {
		return node.id == id;
	}

	protected boolean isClose(AStarNode node) {
		return node.id == -id;
	}

    static class OpenList {

        AStarNode head;
		AStarNode cur;

        public void add(AStarNode node) {
            node.next = null;
			node.prev = null;
			if (cur == null) {
				cur = head = node;
				return;
			}
            AStarNode cur = this.cur;
			if (cur.compareTo(node) >= 0) {
                while (true) {
                    if (cur.compareTo(node) <= 0) {
                        node.next = cur.next;
						node.prev = cur;
						if (cur.next != null) {
							cur.next.prev = node;
						}
						cur.next = this.cur = node;
						return;
                    }
                    if (cur.prev == null) {
						node.next = cur;
						cur.prev = this.cur = node;
						head = this.cur;
						return;
					}
					cur = cur.prev;
                }
            } else {
				while (true) {
					if (cur.compareTo(node) >= 0) {
						node.prev = cur.prev;
						node.next = cur;
						if (cur.prev != null) {
                            cur.prev.next = node;
                        }
						cur.prev = this.cur = node;
						return;
					}

					if (cur.next == null) {
						node.prev = cur;
						cur.next = this.cur = node;
						return;
					}
					cur = cur.next;
				}
			}
        }

        public void clear() {
			cur = head = null;
		}

        public AStarNode poll() {
			AStarNode temp = head;
			head = head.next;
			if (head != null)
				head.prev = null;
			if (temp == cur) {
				cur = head;
			}
			return temp;
		}

        public boolean isEmpty() {
			return head == null;
		}

        @Override
		public String toString() {
			StringBuilder sb = new StringBuilder();
			AStarNode node = head;
			while (node != null) {
				sb.append(node).append(",");
				node = node.next;
			}
			if (sb.length() > 0) {
                sb.delete(sb.length() - 1, sb.length());
            }
			return sb.toString();
		}

    }

    

    public static class AStarNode extends Node implements Comparable<AStarNode> {

        protected int type; // 坐标点类型, 0可寻路
        protected int tempG;
        protected int tempF;
        protected AStarNode tempNext; // 下一个

        AStarNode next;
		AStarNode prev;

        public AStarNode(int type) {
			this.type = type;
		}

        public boolean canClose(AStarNode from, Object o) {
			return true;
		}

		protected boolean canIn(Node from, Object o) {
			return type == 0;
		}

        @Override
        public int compareTo(AStarNode o) {
            return tempF - o.tempF;
        }

    }

    public static class Node {
        int x;
        int y;
        int id;

        public Node(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public Node() {
		}

        public int x() {
            return x;
        }

        public int y() {
            return y;
        }

        @Override
        public String toString() {
			return "[" + x + "," + y + "]";
		}

        @Override
        public int hashCode() {
            return Objects.hash(x, y);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
		    if (o == null || getClass() != o.getClass()) return false;
            Node that = (Node) o;
            if (!Objects.equals(x, that.x)) return false;
		    if (!Objects.equals(y, that.y)) return false;
            return true;
        }
    }

    public static void main(String[] args) throws InterruptedException {
        int LEN = 100;
		AStarNode[][] map = createMapData(100);

		AStar a = new AStar(map);
		List<Node> p = a.findPath(new Node(1, 15), new Node(91, 11));
		System.out.println(p);
		a.printMap();

		// long s1 = System.currentTimeMillis();
        // Random random = new Random();
		// for (int i = 0; i < 100000; i++) {
		// 	int a1 = random.nextInt(99) + 2;
		// 	int a2 = random.nextInt(99) + 2;
		// 	a.findPath(new Node(1, 1), new Node(a1, a2));
		// }
		// long s2 = System.currentTimeMillis();
		// System.out.println(s2 - s1);
	}

	public static AStarNode[][] createMapData(int LEN) {
		// 100*100的地图
		AStarNode[][] map = new AStarNode[LEN][];
		for (int i = 0; i < LEN; i++) {
			AStarNode[] nodes = new AStarNode[LEN];
			for (int j = 0; j < LEN; j++) {
				nodes[j] = new AStarNode(0);
			}
			map[i] = nodes;
		}
		// 模拟阻挡点/不可寻路
		for (int i = 0; i < 30 && 6 + i < LEN; i++) {
			map[8][6 + i].type = 1;
		}
		
//		for (int i = 0; i < 10 && 1 + i < LEN; i++) {
//			map[8+i][6].type = 1;
//		}
		
//		for (int i = 0; i < 15; i++) {
//			map[13][i].type = 1;
//		}
		return map;
	}
}

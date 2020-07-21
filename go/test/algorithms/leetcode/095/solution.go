package solution095

import "fmt"

// TreeNode Definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n TreeNode) String() string {
	return fmt.Sprintf("TreeNode[%d, left=%v, right=%v] ", n.Val, n.Left, n.Right)
}

// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
// leetcode每日一题竟然会重复
// 于是用go写一遍加深印象
// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
// 示例：
// 输入：3
// 输出：
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// 解释：
// 以上的输出对应以下 5 种不同结构的二叉搜索树：

//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3
// 提示：
// 0 <= n <= 8
//对于BST来说，确定了根节点，左右子树的范围就决定了
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return generateSubTrees(1, n)
}

func generateSubTrees(start, end int) []*TreeNode {
	var nodes []*TreeNode
	if start > end {
		nodes = append(nodes, nil)
		return nodes
	}
	for i := start; i <= end; i++ {
		left := generateSubTrees(start, i-1) //左子树范围
		right := generateSubTrees(i+1, end)  //右子树范围
		for _, l := range left {
			for _, r := range right {
				nodes = append(nodes, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return nodes
}

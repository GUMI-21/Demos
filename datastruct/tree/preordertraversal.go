package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 递归,根->左->右 前序遍历,先输出根再输出左在输出右,终止条件是节点为空
func preorderTraversal(root *Node) []int {
	var (
		vals     []int
		preorder func(node *Node)
	)
	preorder = func(node *Node) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return vals
}

// 迭代,显示的维护一个栈
func preorderTraversalV1(root *Node) (vals []int) {
	var stack []*Node
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return vals
}

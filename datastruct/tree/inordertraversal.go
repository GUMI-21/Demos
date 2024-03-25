package tree

// 递归,设方法为将root节点保存到数组中,左->中->右
func inorderTraversal(root *Node) []int {
	var vals []int
	var inorderFunc func(*Node)
	inorderFunc = func(node *Node) {
		if node == nil {
			return
		}
		inorderFunc(node.Left)
		vals = append(vals, node.Val)
		inorderFunc(node.Right)
	}
	inorderFunc(root)
	return vals
}

// 迭代,维护一个栈
func inorderTraversalV1(root *Node) (res []int) {
	var stack []*Node
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

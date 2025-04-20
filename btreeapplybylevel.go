package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	levelsize := 1

	for len(queue) > 0 {
		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			f(node.Data)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelsize = len(queue)
	}
}

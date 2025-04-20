package piscine

	

	type TreeNode struct {

		Data   string

		Left   *TreeNode

		Right  *TreeNode

		Parent *TreeNode 

	}

	

	func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

		if root == nil {

			return nil

		}

	

		if node.Data < root.Data {

			root.Left = BTreeDeleteNode(root.Left, node)

			if root.Left != nil {

				root.Left.Parent = root 

			}

		} else if node.Data > root.Data {

			root.Right = BTreeDeleteNode(root.Right, node)

			if root.Right != nil {

				root.Right.Parent = root 

			}

		} else {

			if root.Left == nil {

				temp := root.Right

				if temp != nil {

					temp.Parent = root.Parent 

				}

				return temp

			} else if root.Right == nil {

				temp := root.Left

				if temp != nil {

					temp.Parent = root.Parent 

				}

				return temp

			}

	

			temp := BTreeMin(root.Right)

	

			root.Data = temp.Data

	

			root.Right = BTreeDeleteNode(root.Right, temp)

			if root.Right != nil {

				root.Right.Parent = root 

			}

		}

		return root

	}

	

	func BTreeMin(node *TreeNode) *TreeNode {

		current := node

		for current.Left != nil {

			current = current.Left

		}

		return current

	}

	

	func BTreeInsertData(root *TreeNode, data string) *TreeNode {

		if root == nil {

			return &TreeNode{Data: data}

		}

		if data < root.Data {

			root.Left = BTreeInsertData(root.Left, data)

			root.Left.Parent = root 

		} else if data > root.Data {

			root.Right = BTreeInsertData(root.Right, data)

			root.Right.Parent = root 

		}

		return root

	}

	

	func BTreeSearchItem(root *TreeNode, data string) *TreeNode {

		if root == nil || root.Data == data {

			return root

		}

		if data < root.Data {

			return BTreeSearchItem(root.Left, data)

		}

		return BTreeSearchItem(root.Right, data)

	}

	

	func BTreeApplyInorder(root *TreeNode, f func(string)) {

		if root != nil {

			BTreeApplyInorder(root.Left, f)

			f(root.Data)

			BTreeApplyInorder(root.Right, f)

		}

	}

	

	func BTreeIsBinary(root *TreeNode) bool {

		return isBST(root, nil, nil)

	}

	

	func isBST(node *TreeNode, min *TreeNode, max *TreeNode) bool {

		if node == nil {

			return true

		}

		if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {

			return false

		}

		return isBST(node.Left, min, node) && isBST(node.Right, node, max)

	}

	
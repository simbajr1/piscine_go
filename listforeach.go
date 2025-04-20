package piscine

// ListForEach applies the function f to each node in the list l
func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

// Add2_node adds 2 to an integer node's data or appends "2" to a string node's data
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Subtract3_node subtracts 3 from an integer node's data or appends "-3" to a string node's data
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

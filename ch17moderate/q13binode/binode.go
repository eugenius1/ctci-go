package q13binode

type BiNode struct {
	node1, node2 *BiNode
	data         int
}

type DataIterator chan int
type NodeIterator chan (*BiNode)

// logical distinction for what the data structure represents
type (
	Tree BiNode
	List BiNode
)

func (i DataIterator) GetAll() []int {
	result := make([]int, 0)
	for data := range i {
		result = append(result, data)
	}
	return result
}

func TreeIterator(treeRoot *BiNode) DataIterator {
	c := make(DataIterator)
	go iterateTreeData(treeRoot, c)
	return c
}

func ListIterator(listHead *BiNode) DataIterator {
	c := make(DataIterator)
	go iterateListData((*List)(listHead), c)
	return c
}

func TreeToList(treeRoot *BiNode) (listHead *BiNode) {
	if treeRoot == nil {
		return nil
	}

	if leftTree := treeRoot.node1; leftTree != nil {
		beforeRoot, parentOfBeforeRoot := (*Tree)(leftTree).last()

		if parentOfBeforeRoot != nil {
			// replace beforeRoot in the tree with beforeRoot.node1, as beforeRoot.node2 == nil
			parentOfBeforeRoot.node2 = beforeRoot.node1
		} else {
			// beforeRoot was root of leftTree so pull up the (lower) left side
			leftTree = beforeRoot.node1
		}
		linkListNodes(beforeRoot, treeRoot)

		if leftTree != nil {
			// store node before beforeRoot
			beforeBeforeRoot, _ := (*Tree)(leftTree).last()
			// recurse on the left tree
			listHead = TreeToList(leftTree)
			// link left list to beforeRoot
			linkListNodes(beforeBeforeRoot, beforeRoot)
		} else {
			// there was only 1 node on the left
			listHead = beforeRoot
		}
	}

	if listHead == nil {
		// nothing on the left tree
		listHead = treeRoot
	}

	if rightTree := treeRoot.node2; rightTree != nil {
		afterRoot, parentOfAfterRoot := (*Tree)(rightTree).first()

		if parentOfAfterRoot != nil {
			// replace afterRoot in the tree with afterRoot.node2, as afterRoot.node1 == nil
			parentOfAfterRoot.node1 = afterRoot.node2
		} else {
			// afterRoot was root of rightTree so pull up the (higher) right side
			rightTree = afterRoot.node2
		}
		linkListNodes(treeRoot, afterRoot)

		if rightTree != nil {
			// recurse on the right tree
			afterRoot.node2 = TreeToList(rightTree)
		}
	}

	return listHead
}

func linkListNodes(first, second *BiNode) {
	first.node2 = second
	second.node1 = first
}

func iterateTreeData(tree *BiNode, c DataIterator) {
	nodes := make(NodeIterator)
	go iterateTreeNodes(tree, nodes)
	for node := range nodes {
		c <- node.data
	}
	close(c)
}

// iterate non-nil tree nodes in order
func iterateTreeNodes(tree *BiNode, c NodeIterator) {
	iterateTreeNodesRecursively(tree, c)
	close(c)
}

func iterateTreeNodesRecursively(tree *BiNode, c NodeIterator) {
	if tree != nil {
		iterateTreeNodesRecursively(tree.node1, c)
		c <- tree
		iterateTreeNodesRecursively(tree.node2, c)
	}
}

func (t *Tree) first() (node, parent *BiNode) {
	node = (*BiNode)(t)
	for node != nil && node.node1 != nil {
		parent = node
		node = node.node1
	}
	return node, parent
}

func (t *Tree) last() (node, parent *BiNode) {
	node = (*BiNode)(t)
	for node != nil && node.node2 != nil {
		parent = node
		node = node.node2
	}
	return node, parent
}

// iterate list nodes in order assuming head is given
func iterateListNodes(head *List, c NodeIterator) {
	node := (*BiNode)(head)
	for node != nil {
		c <- node
		node = node.node2
	}
	close(c)
}

// iterate list data in order assuming head is given
func iterateListData(head *List, c DataIterator) {
	nodes := make(NodeIterator)
	go iterateListNodes(head, nodes)
	for node := range nodes {
		c <- node.data
	}
	close(c)
}

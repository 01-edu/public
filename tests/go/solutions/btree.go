package solutions

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(bt *TreeNode, elem string) *TreeNode {
	if bt == nil {
		return &TreeNode{Data: elem}
	}

	if elem < bt.Data {
		bt.Left = BTreeInsertData(bt.Left, elem)
		bt.Left.Parent = bt
	} else if elem >= bt.Data {
		bt.Right = BTreeInsertData(bt.Right, elem)
		bt.Right.Parent = bt
	}
	return bt
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)
	}
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyPostorder(root.Left, f)
		BTreeApplyPostorder(root.Right, f)
		f(root.Data)
	}
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}

	return BTreeSearchItem(root.Right, elem)
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := BTreeLevelCount(root.Left)
	right := BTreeLevelCount(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func BTreeIsBinary(root *TreeNode) bool {
	condLeft := true
	condRight := true

	if root == nil {
		return true
	}

	if root.Left != nil {
		condLeft = BTreeIsBinary(root.Left) && root.Data >= root.Left.Data
	}

	if root.Right != nil {
		condRight = BTreeIsBinary(root.Right) && root.Data <= root.Right.Data
	}

	return condLeft && condRight
}

func applyGivenOrder(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if level == 0 {
		arg := interface{}(root.Data)
		f(arg)
	} else if level > 0 {
		applyGivenOrder(root.Left, level-1, f)
		applyGivenOrder(root.Right, level-1, f)
	}
}

//Apply the function f level by level
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 0; i < h; i++ {
		applyGivenOrder(root, i, f)
	}
}

/*-----------------------------------
* node->  (A)                  (B)
*         /                   /  \
* temp-> (B)    -------->   (C)  (A)
*       /   \                    /
*     (C)  [z]                [z]
*------------------------------------*/
func BTreeRotateRight(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Left == nil {
		return node
	}

	temp := node.Left

	node.Left = temp.Right

	if temp.Right != nil {
		temp.Right.Parent = node
	}

	if node.Parent != nil {
		if node == node.Parent.Left {
			node.Parent.Left = temp
		} else {
			node.Parent.Right = temp
		}
	}

	temp.Right = node
	temp.Parent = node.Parent
	node.Parent = temp
	return temp
}

/*------------------------------------
*  node->(A)                  (B)
*         \                 /  \
*  temp-> (B) ------->    (A)  (C)
*        /   \              \
*       [z] (C)             [z]
*------------------------------------- */
func BTreeRotateLeft(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Right == nil {
		return node
	}

	temp := node.Right
	node.Right = temp.Left

	if temp.Left != nil {
		temp.Left.Parent = node
	}
	if node.Parent != nil {
		if node == node.Parent.Left {
			node.Parent.Left = temp
		} else {
			node.Parent.Right = temp
		}
	}

	temp.Left = node
	temp.Parent = node.Parent

	node.Parent = temp
	return temp
}

//Returns the maximum node in the subtree started by root
func BTreeMax(root *TreeNode) *TreeNode {

	if root == nil || root.Right == nil {
		return root
	}

	return BTreeMax(root.Right)
}

//Returns the minimum value in the subtree started by root
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	return BTreeMin(root.Left)
}

func BTreeTransplant(root, node, repla *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	replacement := node
	if node.Parent == nil {
		root = repla
	} else if node == node.Parent.Left {
		replacement.Parent.Left = repla
	} else {
		replacement.Parent.Right = repla
	}
	if repla != nil {
		repla.Parent = node.Parent
	}

	return root
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if node.Left == nil {
			root = BTreeTransplant(root, node, node.Right)
		} else if node.Right == nil {
			root = BTreeTransplant(root, node, node.Left)
		} else {
			y := BTreeMin(node.Right)
			if y != nil && y.Parent != node {
				root = BTreeTransplant(root, y, y.Right)

				y.Right = node.Right
				y.Right.Parent = y
			}
			root = BTreeTransplant(root, node, y)
			y.Left = node.Left
			y.Left.Parent = y
		}
	}
	return root
}

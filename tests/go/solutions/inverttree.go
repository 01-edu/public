package solutions

type TNode struct {
    Val   int
    Left  *TNode
    Right *TNode
}

func Invert(root *TNode) {
    if root != nil {
        temp := root.Left
        root.Left = root.Right
        root.Right = temp

        Invert(root.Left)
        Invert(root.Right)
    }
}

func InvertTree(root *TNode) *TNode {
    Invert(root)
    return root
}

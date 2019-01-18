/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return treeSymmetric(root.Left, root.Right)
}

func treeSymmetric(left *TreeNode, right *TreeNode) bool {
    
    if left == nil && right == nil {
        return true
    }
    
    if left == nil || right == nil {
        return false
    }
    
    if left.Val != right.Val {
        return false
    }
    
    return treeSymmetric(left.Left, right.Right) && treeSymmetric(left.Right, right.Left)
}

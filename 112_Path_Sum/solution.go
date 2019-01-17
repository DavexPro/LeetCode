/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    
    rest := sum - root.Val
    if root.Left == nil && root.Right == nil && rest == 0 {
        return true
    }
    
    return hasPathSum(root.Left, rest) || hasPathSum(root.Right, rest)
}

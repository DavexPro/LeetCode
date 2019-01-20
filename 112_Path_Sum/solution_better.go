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
    
    if root.Left == nil && root.Right == nil {
        return root.Val == sum
    }
    
    rest := sum - root.Val
    return hasPathSum(root.Left, rest) || hasPathSum(root.Right, rest)
}

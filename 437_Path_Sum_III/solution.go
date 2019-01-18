/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    
    if root == nil {
        return 0
    }
    
    return leafPathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func leafPathSum(root *TreeNode, sum int) int {
    
    if root == nil {
        return 0
    }
    
    path := leafPathSum(root.Left, sum - root.Val) + leafPathSum(root.Right, sum - root.Val)
    
    if root.Val == sum {
        path += 1
    }
    
    return path
}

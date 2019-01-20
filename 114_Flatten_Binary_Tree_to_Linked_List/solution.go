/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return
    }
    if root.Left == nil {
        flatten(root.Right)
        return
    }
    if root.Right == nil {
        root.Right, root.Left = root.Left, root.Right
        flatten(root.Right)
        return
    }

    flatten(root.Left)
    flatten(root.Right)
    
    curr := root.Left
    for curr.Right != nil {
        curr = curr.Right
    }
    curr.Right = root.Right
    root.Right = root.Left
    
    root.Left = nil
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var ret []int
    if root == nil {
        return ret
    }
    
    helper(root, &ret)
    return ret
}

func helper(root *TreeNode, ret *[]int) {
    if root == nil {
        return 
    }
    
    helper(root.Left, ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}

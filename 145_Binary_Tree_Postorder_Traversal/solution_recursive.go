/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var ret []int
    helper(root, &ret)
    return ret
}

func helper(root *TreeNode, ret *[]int) {
    if root == nil {
        return
    }    
    
    helper(root.Left, ret)
    helper(root.Right, ret)
    *ret = append(*ret, root.Val)
}

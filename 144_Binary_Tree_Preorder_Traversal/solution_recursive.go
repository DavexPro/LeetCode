/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var ret []int
    helper(root, &ret)
    return ret
}

func helper(root *TreeNode, ret *[]int){
    if root == nil {
        return 
    }
    
    *ret = append(*ret, root.Val)
    helper(root.Left, ret)
    helper(root.Right, ret)
}

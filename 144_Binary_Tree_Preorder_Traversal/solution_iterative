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
    var stack []*TreeNode
    
    curr := root
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            ret = append(ret, curr.Val)
            curr = curr.Left
        }
        
        // pop from stack
        curr = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        curr = curr.Right
    }
    
    return ret
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var ret []int
    var todoList []*TreeNode
    
    if root == nil {
        return ret
    }
    
    todoList = append(todoList, root)
    for len(todoList) > 0 {
        var right int
        var waitList []*TreeNode
        for _, node := range todoList {
            right = node.Val
            if node.Left != nil {
                waitList = append(waitList, node.Left)
            }
            if node.Right != nil {
                waitList = append(waitList, node.Right)
            }
        }
        
        ret = append(ret, right)
        todoList = waitList
    }
    
    return ret
}

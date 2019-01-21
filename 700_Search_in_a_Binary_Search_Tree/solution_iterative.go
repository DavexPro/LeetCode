/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    
    if root == nil {
        return nil
    }
    
    var todoList []*TreeNode
    todoList = append(todoList, root)

    for len(todoList) > 0 {
        var waitList []*TreeNode
        
        for _, node := range todoList {
            if node.Val == val {
                return node
            }
            if node.Left != nil {
                waitList = append(waitList, node.Left)
            }
            if node.Right != nil {
                waitList = append(waitList, node.Right)
            }
        }
        
        todoList = waitList
    }
    
    return nil
}

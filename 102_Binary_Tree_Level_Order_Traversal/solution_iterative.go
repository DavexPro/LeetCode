/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var ret [][]int
    var todoList []*TreeNode
    
    if root == nil {
        return ret
    }
    
    level := 0
    todoList = append(todoList, root)
    for len(todoList) > 0 {
        var waitList []*TreeNode
        
        for _, node := range todoList {
            if node == nil {
                continue
            }
            
            if len(ret) < level + 1 {
                ret = append(ret, []int{node.Val})
            }else{
                ret[level] = append(ret[level], node.Val)
            }
            
            waitList = append(waitList, node.Left)
            waitList = append(waitList, node.Right)
        }
        
        level ++
        todoList = waitList
    }
    
    return ret
}

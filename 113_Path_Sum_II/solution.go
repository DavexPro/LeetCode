/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    
    ret := make([][]int,0)
    path := make([]int,0)
    hasPathSum(root, sum, path, &ret)
    
    return ret
}

func hasPathSum(root *TreeNode, sum int, path []int, ret *[][]int) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil && root.Val == sum {
        path = append(path, root.Val)
        *ret = append(*ret, append([]int(nil), path...))
        return
    }
    
    hasPathSum(root.Left, sum - root.Val, append(path, root.Val), ret)
    hasPathSum(root.Right, sum - root.Val, append(path, root.Val), ret)
}

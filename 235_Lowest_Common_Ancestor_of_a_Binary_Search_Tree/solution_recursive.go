/**
* Definition for TreeNode.
* type TreeNode struct {
*     Val int
*     Left *ListNode
*     Right *ListNode
* }
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || p == nil || q == nil {
        return nil
    }

    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }else if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }
    
    return root
}

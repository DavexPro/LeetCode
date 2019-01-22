func sortedSquares(A []int) []int {
    if A == nil || len(A) == 0 {
        return A
    }
    
    for i, v := range A {
        A[i] = v*v
    }
    
    sort.Sort(sort.IntSlice(A))
    
    return A
}

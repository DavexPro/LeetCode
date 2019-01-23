func longestMountain(A []int) int {
    if A == nil {
        return 0
    }
    
    n := len(A)
    ret := 0
    base := 0
    
    for base < n {
        end := base
        
        // ensuer in array
        if end + 1 < n && A [end] < A[end + 1]{
            // find the peak
            for end + 1 < n && A[end] < A[end + 1] {
                end++
            }
            
            // still check whether in array
            // and check whether it is peak
            if end + 1 < n && A[end] > A[end + 1] {
                // find the bottom
                for end + 1 < n && A[end] > A[end + 1] {
                    end++
                }
                
                ret = max(ret, end - base + 1)
            }
        }
        base = max(end, base + 1)
    }
    return ret
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

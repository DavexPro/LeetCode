func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    
    // that must be error
    if m == 0 && n == 0 {
        return float64(0)
    }
    
    var halfLen int
    if (m + n) % 2 == 1 {
        halfLen = (m + n + 1) / 2
    }else{
        halfLen = (m + n + 2) / 2
    }
    
    var arr []int
    
    i, j := 0, 0
    for len(arr) < halfLen {
        if i < m && j < n {
            if nums1[i] > nums2[j] {
                arr = append(arr, nums2[j])
                j++
            }else{
                arr = append(arr, nums1[i])
                i++
            }
            continue
        }
        
        if i < m {
            arr = append(arr, nums1[i])
            i++
        }
        
        if j < n {
            arr = append(arr, nums2[j])
            j++
        }
    }
    
    if (m + n) % 2 == 1 {
        return float64(arr[len(arr) - 1])
    }else{
        sum := arr[len(arr) - 1] + arr[len(arr) - 2]
        return float64(sum) / 2.0
    }
    
    return float64(0)
}

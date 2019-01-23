func trap(height []int) int {
    if height == nil || len(height) == 0 {
        return 0
    }

    ret := 0
    left, right := 0, len(height) - 1
    leftMax, rightMax := 0, 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            }else{
                ret += leftMax - height[left]
            }
            left++
        }else{
            if height[right] >= rightMax {
                rightMax = height[right]
            }else{
                ret += rightMax - height[right]
            }
            right--
        }
    }
    
    return ret
}

func isHappy(n int) bool {
    if n == 1 || n == 10 {
        return true
    }
    
    visited := make(map[int]bool)
    
    ret := calcHappy(n)
    for true {
        if ret == 1 {
            return true
        }
        
        if visited[ret] == true{
            return false
        }
        
        visited[ret] = true
        ret = calcHappy(ret)
    }
    
    return false
}

func calcHappy(num int) int {
    ret := 0
    for num > 0 {
        digit := num % 10
        ret += digit * digit
        num /= 10
    }
    return ret
}

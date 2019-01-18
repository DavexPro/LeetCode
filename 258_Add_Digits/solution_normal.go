func addDigits(num int) int {
    if num < 10 {
        return num
    }
    
    sum := 0
    for num > 0 {
        digit := num % 10
        sum += digit
        num /= 10
    }
    
    if sum < 10 {
        return sum
    }
    
    return addDigits(sum)
}

func reverse(x int) int {
    
    num, sum := x, 0
    
    for num != 0 {
        digit := num % 10
        sum = sum*10 + digit
        num /= 10
    }
    
    if sum >int(math.Pow(2,31)-1) || sum < -int(math.Pow(2,31)) {
        sum = 0
    }
    
    return sum
}

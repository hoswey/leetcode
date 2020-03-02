func candy(ratings []int) int {
        
    candies := make([]int, len(ratings))
    candies[0] = 1
    
    for i := 1; i < len(ratings); i++ {
        
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1] + 1
        } else {
            candies[i] = 1
        }
    } 
    
    for i := len(ratings) - 2; i >= 0; i-- {        
        if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
            candies[i] = candies[i+1] + 1
        } 
    }

    sum := 0
    for _, c := range candies {
        sum += c
    }
    
    return sum
}
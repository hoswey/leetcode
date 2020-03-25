package main

import "fmt"

func twoSum(n int) []float64 {
    
	if n == 0 {
		return nil
	}

    times := make([]int, 5 * n + 1)
    recurse(n, n, 0, times)

    total := 1
    for i := 0; i < n; i++ {
    	total *= 6
    }

    probility := make([]float64, len(times))
    for i := 0; i < len(probility); i++ {
    	probility[i] = float64(times[i]) / float64(total)
    }

    return probility

}

func recurse(n int, cur int, sum int, times []int) {

	if cur == 0 {
		times[sum-n]++
		return 
	}

	for i := 1; i <= 6; i++ {
		recurse(n, cur - 1, sum + i, times)
	}
}

func main() {

	fmt.Printf("%v\n", twoSum(11))
}
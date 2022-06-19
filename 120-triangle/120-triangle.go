import "fmt"
import "math"

func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    
    memo := make([][]int, n)
    for i := range memo{
        memo[i] = make([]int, n)
        for j := range memo[i]{
            memo[i][j] = math.MaxInt64
        }
    }
    
    memo[0][0] = triangle[0][0]
    
    for i:=1; i<n; i++{
        for j := range triangle[i]{
            memo[i][j] = min(memo[i-1][j], memo[i-1][max(j-1,0)])+ triangle[i][j]
        }
    }
    
    result := math.MaxInt64
    for i := range memo[n-1] {
        if memo[n-1][i] < result {
            result = memo[n-1][i]
        }
    }
    return result
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
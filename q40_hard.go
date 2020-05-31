/**
题目描述
现有一组砝码，重量互不相等，分别为m1,m2,m3…mn；
每种砝码对应的数量为x1,x2,x3...xn。现在要用这些砝码去称物体的重量(放在同一侧)，问能称出多少种不同的重量。
注：
称重重量包括0
*/
package main
 
import "fmt"
 
func solve(n int) {
    m := make([]int, n)
    c := make([]int, n)
    sum := 0
    for i := range m {
        fmt.Scan(&m[i])
    }
    for i := range c {
        fmt.Scan(&c[i])
        sum += c[i]*m[i]
    }
    dp := make([]bool, sum+1)
    dp[0] = true
    for i := 0; i < n; i++ {
        for j := 0; j < c[i]; j++ {
            for k := sum; k >= m[i]; k-- {
                if dp[k-m[i]] {
                    dp[k] = true
                }
            }
        }
    }
    cnt := 0
    for i := range dp {
        if dp[i] {
            cnt++
        }
    }
    fmt.Println(cnt)
}
 
func main() {
    var n int
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        solve(n)
    }
}
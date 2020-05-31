/**
题目描述
Levenshtein 距离，又称编辑距离，指的是两个字符串之间，由一个转换成另一个所需的最少编辑操作次数。
许可的编辑操作包括将一个字符替换成另一个字符，插入一个字符，删除一个字符。
*/
package main
  
import "fmt"
  
func min(a ...int) int {
    ret := a[0]
    for i := range a {
        if a[i] < ret {
            ret = a[i]
        }
    }
    return ret
}
  
func solve(a, b string) {
    dp := make([][]int, len(a)+1)
    for i := range dp {
        dp[i] = make([]int, len(b)+1)
        dp[i][0] = i
    }
    for i := range dp[0] {
        dp[0][i] = i
    }
    for i := range a {
        for j := range b {
            if a[i] == b[j] {
                dp[i+1][j+1] = dp[i][j]
            } else {
                dp[i+1][j+1] = min(dp[i][j+1]+1, dp[i+1][j]+1, dp[i][j]+1)
            }
        }
    }
    fmt.Println(dp[len(a)][len(b)])
}
  
func main() {
    var a, b string
    for {
        n, err := fmt.Scan(&a, &b)
        if err != nil || n!= 2 {
            break
        }
        solve(a, b)
    }
}

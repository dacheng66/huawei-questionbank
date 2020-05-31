/**
题目描述
问题描述：在计算机中，通配符一种特殊语法，广泛应用于文件搜索、数据库、正则表达式等领域。现要求各位实现字符串通配符的算法。
要求：
实现如下2个通配符：
*：匹配0个或以上的字符（字符由英文字母和数字0-9组成，不区分大小写。下同）
？：匹配1个字符


输入：
通配符表达式；
一组字符串。


输出：
返回匹配的结果，正确输出true，错误输出false
*/
package main
  
import (
    "fmt"
)
  
func main() {
    var (
        wild string
        str string
    )
    for {
        _, err := fmt.Scan(&wild)
        if err != nil {
            break
        }
        fmt.Scan(&str)
          
        l1 := len(wild)
        l2 := len(str)
        dp := make([][]bool, l1+1)
        for i := range dp {
            dp[i] = make([]bool, l2+1)
        }
        dp[0][0] = true
        for i := 1; i <= l2; i++ {
            // 空通配符匹配字符串
            dp[0][i] = false
        }
        for i := 1; i <= l1; i++ {
            dp[i][0] = dp[i-1][0] && (wild[i] == '*')
        }
          
        for i := 1; i <= l1; i++ {
            for j := 1; j <= l2; j++ {
                w := wild[i-1]
                s := str[j-1]
                if w == '*' {
                    // * 号匹配 0 个 或 多个 字符
                    dp[i][j] = dp[i-1][j] || dp[i][j-1]
                } else {
                    dp[i][j] = (w == s || w == '?') && dp[i-1][j-1]
                }
            }
        }
        fmt.Println(dp[l1][l2])
    }
}

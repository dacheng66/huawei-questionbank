/**
题目描述
查找两个字符串a,b中的最长公共子串。若有多个，输出在较短串中最先出现的那个。
*/
package main
 
import (
    "fmt"
)
 
func main() {
    var a, b string
    for {
        res, _ := fmt.Scan(&a, &b)
        if res == 0 {
            break
        }
 
        if len(a) > len(b) {
            a, b = b, a
        }
        la := len(a) + 1
        lb := len(b) + 1
 
        arr := make([][]int, la)
        for i := 0; i < la; i++ {
            arr[i] = make([]int, lb)
        }
 
        var ma int
        var x int
        for i := 1; i < la; i++ {
            for j := 1; j < lb; j++ {
                if a[i-1] == b[j-1] {
                    arr[i][j] = arr[i-1][j-1] + 1
                    if ma < arr[i][j] {
                        ma = arr[i][j]
                        x = i
                    }
                } else {
                    arr[i][j] = 0
                }
            }
        }
 
        if ma == 0 {
            fmt.Println("")
            continue
        }
        var ans string
        ans = a[x-ma : x]
        fmt.Println(ans)
    }
}